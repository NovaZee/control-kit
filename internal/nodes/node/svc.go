package node

import "github.com/NovaZee/control-kit/osutil"

func (m *NodesWithMetrics) GetNodes() map[string]*osutil.NodeMetrics {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make(map[string]*osutil.NodeMetrics)
	for _, metricsMap := range m.nodes {
		for key, metrics := range metricsMap {
			metricsMap[key] = metrics
			break
		}
	}
	return result
}

func (m *NodesWithMetrics) GetNode(nodeID string) map[string]*osutil.NodeMetrics {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if _, ok := m.nodes[nodeID]; ok {
		return m.nodes[nodeID]
	}
	return nil
}
