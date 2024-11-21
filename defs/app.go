package defs

var IniPath = appPath + "conf/authConfig.ini"

const (
	appPath = "/opt/app/"
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

const KniSection = appNodesConfigPrefix + "/%s/app/kni"
const FStackSection = appNodesConfigPrefix + "/%s/app/f-stack"
const appSection = appNodesConfigPrefix + "/%s/app/app"
const AuthSection = appNodesConfigPrefix + "/%s/biz%d/auth%d"
const AuthStatsReportSection = appNodesConfigPrefix + "/%s/biz%d/auth%d_statsreport"
const AuthFlowReportSection = appNodesConfigPrefix + "/%s/biz%d/auth%d_flowreport"
const AuthReportLogSection = appNodesConfigPrefix + "/%s/biz%d/auth%d_reportlog"
const AuthFlowLogSection = appNodesConfigPrefix + "/%s/biz%d/auth%d_flowlog"
const PortSection = appNodesConfigPrefix + "/%s/ports/port%d"

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
