package conf

func (nc *NodeConfig) GetTargetConfig(node string) string {
	nc.mu.RLock()
	defer nc.mu.RUnlock()
	return nc.config[node]
}

func (nc *NodeConfig) EditTargetConfig(node string, config string) string {
	nc.mu.RLock()
	defer nc.mu.RUnlock()
	nc.config[node] = config
	// TODO: write to etcd，提交更新命令
	return nc.config[node]
}
