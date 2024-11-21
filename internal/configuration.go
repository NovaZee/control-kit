package internal

import (
	_logger "actuator/logger"
	"context"
	"encoding/json"
	"fmt"
	"github.com/NovaZee/control-kit/defs"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

func (a *Actuator) SearchInI(ctx context.Context) ([]string, error) {
	cfg, err := ini.Load(defs.IniPath)
	if err != nil {
		return []string{}, err
	}
	kv := a.etcdX.Client

	netInterface := make([]string, 0)
	iniData := make(map[string]map[string]string)
	for _, section := range cfg.Sections() {
		sectionName := section.Name()
		iniData[sectionName] = section.KeysHash()
	}

	for i := 0; ; i++ {
		port := cfg.Section(fmt.Sprintf("%s", fmt.Sprintf(defs.SectionPortKey, i)))
		if len(port.Keys()) == 0 {
			break
		}
		netInterface = append(netInterface, port.Key("addr").String())
	}

	jsonData, err := json.Marshal(iniData)
	if err != nil {
		return []string{}, err
	}

	key := defs.APPNodesConfigPrefix + "/" + defs.NodeId + "/"
	_, err = kv.Put(ctx, key, string(jsonData))
	if err != nil {
		return []string{}, err
	}

	_logger.Ins.Info("Successfully loaded INI configuration", zap.String("key", key), zap.String("value", string(jsonData)))
	return netInterface, nil
}

//func (a *Actuator) SearchInI(ctx context.Context) ([]string, error) {
//	cfg, err := ini.Load(defs.IniPath)
//	if err != nil {
//		return []string{}, err
//	}
//	kv := a.etcdX.Client
//	sections := make(map[string]string)
//	netInterface := make([]string, 0)
//
//	kniKey := fmt.Sprintf(defs.KniSection, defs.NodeId)
//	fStackKey := fmt.Sprintf(defs.FStackSection, defs.NodeId)
//	xProxyKey := fmt.Sprintf(defs.XProxySection, defs.NodeId)
//
//	sections[kniKey] = emitJsonError(cfg.Section(defs.SectionKniKey).KeysHash())
//	sections[fStackKey] = emitJsonError(cfg.Section(defs.SectionFStackKey).KeysHash())
//	sections[xProxyKey] = emitJsonError(cfg.Section(defs.SectionXProxyKey).KeysHash())
//
//	for i := 0; ; i++ {
//		authSection := fmt.Sprintf("%s%d", defs.SectionAuthKey, i)
//		auth := cfg.Section(authSection)
//		if len(auth.Keys()) == 0 {
//			break
//		}
//
//		flowLog := cfg.Section(fmt.Sprintf("%s", fmt.Sprintf(defs.SectionAuthFlowLogKey, i)))
//		flowReport := cfg.Section(fmt.Sprintf("%s", fmt.Sprintf(defs.SectionAuthFlowReportKey, i)))
//		statsReport := cfg.Section(fmt.Sprintf("%s", fmt.Sprintf(defs.SectionAuthStatsReportKey, i)))
//		reportLog := cfg.Section(fmt.Sprintf("%s", fmt.Sprintf(defs.SectionAuthReportLogKey, i)))
//
//		sections[fmt.Sprintf(defs.AuthSection, defs.NodeId, i, i)] = emitJsonError(auth.KeysHash())
//		sections[fmt.Sprintf(defs.AuthFlowReportSection, defs.NodeId, i, i)] = emitJsonError(flowReport.KeysHash())
//		sections[fmt.Sprintf(defs.AuthFlowLogSection, defs.NodeId, i, i)] = emitJsonError(flowLog.KeysHash())
//		sections[fmt.Sprintf(defs.AuthReportLogSection, defs.NodeId, i, i)] = emitJsonError(reportLog.KeysHash())
//		sections[fmt.Sprintf(defs.AuthStatsReportSection, defs.NodeId, i, i)] = emitJsonError(statsReport.KeysHash())
//	}
//
//	for i := 0; ; i++ {
//		port := cfg.Section(fmt.Sprintf("%s", fmt.Sprintf(defs.SectionPortKey, i)))
//		if len(port.Keys()) == 0 {
//			break
//		}
//		sections[fmt.Sprintf(defs.PortSection, defs.NodeId, i)] = emitJsonError(port.KeysHash())
//		netInterface = append(netInterface, port.Key("addr").String())
//	}
//
//	for key, value := range sections {
//		_, err = kv.Put(ctx, key, value)
//		if err != nil {
//			xlog.Ins.Error("Error writing to etcd", zap.Error(err))
//		}
//	}
//
//	xlog.Ins.Info("Successfully loaded xproxy configuration", zap.Any("configurations", sections))
//	return netInterface, nil
//}

func emitJsonError(input map[string]string) string {
	jsonStr, _ := json.Marshal(input)
	return string(jsonStr)
}
