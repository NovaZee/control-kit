package defs

var IniPath = APPPath + "conf/authConfig.ini"

const (
	APPPath = "/opt/app/"
)

const (
	SectionKniKey    = "kni"
	SectionFStackKey = "f-stack"
	SectionappKey    = "app"

	SectionAuthKey            = "auth"
	SectionAuthStatsReportKey = "auth%d_statsreport"
	SectionAuthFlowReportKey  = "auth%d_flowreport"
	SectionAuthReportLogKey   = "auth%d_reportlog"
	SectionAuthFlowLogKey     = "auth%d_flowlog"

	SectionPortKey = "port%d"
)

const KniSection = APPNodesConfigPrefix + "/%s/app/kni"
const FStackSection = APPNodesConfigPrefix + "/%s/app/f-stack"
const appSection = APPNodesConfigPrefix + "/%s/app/app"
const AuthSection = APPNodesConfigPrefix + "/%s/biz%d/auth%d"
const AuthStatsReportSection = APPNodesConfigPrefix + "/%s/biz%d/auth%d_statsreport"
const AuthFlowReportSection = APPNodesConfigPrefix + "/%s/biz%d/auth%d_flowreport"
const AuthReportLogSection = APPNodesConfigPrefix + "/%s/biz%d/auth%d_reportlog"
const AuthFlowLogSection = APPNodesConfigPrefix + "/%s/biz%d/auth%d_flowlog"
const PortSection = APPNodesConfigPrefix + "/%s/ports/port%d"

// ETCD etc结构
///config
//└── node-1-config
//    ├── app
//    │   ├── [kni]
//    │   ├── [f-stack]
//    │   └── [app]
//    ├── biz
//    │   ├── biz1
//    │   │   ├── [authx]
//    │   │   ├── [authx_flowlog]
//    │   │   ├── [authx_flowreport]
//    │   │   ├── [authx_reportlog]
//    │   │   └── [authx_statsreport]
//    │   ├── biz2
//    │   │   ├── [authx]
//    │   │   ├── [authx_flowlog]
//    │   │   ├── [authx_flowreport]
//    │   │   ├── [authx_reportlog]
//    │   │   └── [authx_statsreport]
//    └── port
//        ├── [port0]
//        ├── [port1]
//        ├── [port2]
//        └── [port3]
