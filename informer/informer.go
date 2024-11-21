package informer

import (
	"context"
	"github.com/NovaZee/control-kit/core"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

// EventHandler 是事件处理器接口
type EventHandler interface {
	OnAdd(key string, res interface{})
	OnUpdate(key string, oldItem, newItem interface{})
	OnDelete(key string, res interface{})
}

// Informer 监听和同步资源变化
type Informer struct {
	ctx         context.Context
	cli         *core.EtcdX
	resources   map[string]*interface{}
	handler     EventHandler
	mu          sync.RWMutex
	key         string        // 监听的键
	cond        *sync.Cond    // 条件变量用于通知
	eventQueue  []interface{} // 存储事件的队列
	workerCount int           // 工作者数量
}

// NewInformer 创建一个新的 Informer 实例
func NewInformer(ctx context.Context, cli *core.EtcdX, handler EventHandler, workerCount int, key string) *Informer {
	i := &Informer{
		cli:         cli,
		ctx:         ctx,
		resources:   make(map[string]*interface{}),
		handler:     handler,
		workerCount: workerCount,
		eventQueue:  make([]interface{}, 500),
		key:         key,
		mu:          sync.RWMutex{},
	}
	i.cond = sync.NewCond(&i.mu)

	return i
}

// Start 启动Informer，它会定期检查数据源并触发事件
func (i *Informer) Start() {
	// 启动多个工作者 goroutine 处理队列
	for j := 0; j < i.workerCount; j++ {
		go i.worker(j)
	}

}

func (i *Informer) Watch() {
	rch := i.cli.Client.Watch(i.ctx, i.key, clientv3.WithPrefix())

	// 启动监听并响应数据变化
	go func() {
		for {
			select {
			case wresp := <-rch:
				for _, ev := range wresp.Events {
					switch ev.Type {
					case clientv3.EventTypePut:
						if ev.Kv.ModRevision == ev.Kv.CreateRevision {
							newEvent := EventTypeAdd{Item: ev.Kv}
							i.addEventToQueue(newEvent)
						} else {
							newEvent := EventTypeUpdate{OldItem: ev.Kv, NewItem: ev.Kv}
							i.addEventToQueue(newEvent)
						}
					case clientv3.EventTypeDelete:
						newEvent := EventTypeDelete{Item: ev.Kv}
						i.addEventToQueue(newEvent)
					}
				}
			}
		}
	}()
}

// worker 处理事件队列中的任务
func (i *Informer) worker(workerID int) {
	for {
		i.mu.Lock()

		// 等待条件变量被通知
		for len(i.eventQueue) == 0 {
			i.cond.Wait()
		}

		// 获取队列中的事件并处理
		event := i.eventQueue[0]
		i.eventQueue = i.eventQueue[1:]

		i.mu.Unlock()

		// 处理事件
		switch ev := event.(type) {
		case EventTypeAdd:
			i.handler.OnAdd(string(ev.Item.Key), ev.Item.Value)
		case EventTypeUpdate:
			i.handler.OnUpdate(string(ev.NewItem.Key), ev.OldItem.Value, ev.NewItem.Value)
		case EventTypeDelete:
			i.handler.OnDelete(string(ev.Item.Key), ev.Item.Value)
		}
	}
}

// addEventToQueue 将事件加入队列，并通知工作者处理
func (i *Informer) addEventToQueue(event interface{}) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.eventQueue = append(i.eventQueue, event)

	// 唤醒一个工作者
	i.cond.Signal()
}

type EventTypeAdd struct {
	Item *mvccpb.KeyValue
}

type EventTypeUpdate struct {
	OldItem, NewItem *mvccpb.KeyValue
}

type EventTypeDelete struct {
	Item *mvccpb.KeyValue
}
