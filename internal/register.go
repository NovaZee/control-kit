package internal

import (
	_logger "actuator/logger"
	"context"
	"fmt"
	"github.com/NovaZee/control-kit/defs"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

func (a *Actuator) RegisterAndKeepAlive(ctx context.Context, flag string, vi []string) error {
	// 创建租约并设置 TTL
	leaseResp, err := a.etcdX.Client.Grant(ctx, 10)
	if err != nil {
		return err
	}

	// 设置服务注册的键值
	key := fmt.Sprintf("%s/%s/", defs.APPNodesOnlinePrefix, defs.NodeId)
	register := defs.Register{
		Instance: defs.NodeId,
		BindIp:   vi,
	}

	// 将键值绑定到租约
	_, err = a.etcdX.Client.Put(ctx, key, register.String(), clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}
	// 开始自动续租
	ch, err := a.etcdX.Client.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		return err
	}
	// 监听租约续租的响应
	go func() {
		for {
			<-ch
			_logger.Ins.Debug("Lease renewed for service instance", zap.String("instanceId", defs.NodeId))
		}
	}()
	_logger.Ins.Info("Register node success", zap.String("key", key), zap.String("value", register.String()))
	return nil
}
