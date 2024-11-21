package defs

import "encoding/json"

// NodeId  xp实例标识 不表示某个中转，存在一对多关系
var NodeId string

// ETCD etc结构
// ├── actuator                        # actuator 组件信息目录
// │   └── nodes                       # 节点信息目录
// │       └── config                  # 节点配置目录
// │           ├── node1-config        # 节点 "node1" 的配置信息
// │           └── ...                 # 更多节点配置
// └── app                          # app 组件信息目录
//     ├── nodes                       # 节点信息目录，包含配置、在线状态、监控和自定义信息
//     │   ├── config                  # 配置目录
//     │   │   ├── node1-config        # 节点 "node1" 的配置信息
//     │   │   ├── node2-config        # 节点 "node2" 的配置信息
//     │   │   └── ...                 # 更多节点配置
//     │   ├── online                  # z注册节点列表，存储当前在线的节点 ID 构成：node-ipToUint32
//     │   │   ├── node1-id            # 在线节点 "node1" 的标识
//     │   │   └── node2-id            # 在线节点 "node2" 的标识
//     │   ├── metrics                 # 监控信息目录，用于存储各个节点的监控数据
//     │   │   ├── node1               # 节点 "node1" 的监控数据
//     │   │   │   └── cpu             # CPU 使用率
//     │   │   ├── node2               # 节点 "node2" 的监控数据
//     │   │   │   └── memory          # 内存使用情况
//     │   │   ├── node3               # 节点 "node3" 的监控数据
//     │   │   │   └── disk            # 磁盘使用情况
//     │   │   └── ...                 # 更多节点监控数据
//     │   └── custom                  # 自定义信息目录
//     │       └── node                # 自定义的节点信息
//     └── ...                         # 待扩展

const (
	app      = "/app"
	Actuator = "/actuator"
	Nodes    = "/nodes"
	Online   = "/online"
	Metrics  = "/metrics"

	ActuatorNodesConfigPrefix = "/actuator/nodes/config"
	appNodesConfigPrefix      = "/app/nodes/config"
	appNodesOnlinePrefix      = "/app/nodes/online"
	appNodesMetricsPrefix     = "/app/nodes/metrics"
	appNodesCustomPrefix      = "/app/nodes/custom"
)

type Register struct {
	Instance string   `json:"instance"`
	BindIp   []string `json:"bind_ip"`
}

func (r Register) String() string {
	jsonStr, _ := json.Marshal(r)
	return string(jsonStr)
}
