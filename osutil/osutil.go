package osutil

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"os/exec"
	"strings"
)

type SystemMetrics struct {
	CPUUsage       float64 `json:"cpu_usage"`
	MemoryUsage    float64 `json:"memory_usage"`
	DiskUsage      float64 `json:"disk_usage"`
	SuperVisor     string  `json:"supervisor"`
	TcpConnections int     `json:"tcp_connections"`
	FileDescriptor int     `json:"file_descriptor"`
}

type NodeMetrics struct {
	NodeID    string         `json:"node_id"`
	Status    string         `json:"status"`
	Metrics   *SystemMetrics `json:"metrics"`
	Timestamp int64          `json:"timestamp"`
}

// CollectMetrics 收集系统指标
func CollectMetrics(pname string) (*SystemMetrics, error) {
	// 获取 CPU 使用率
	cpuPercent, _ := cpu.Percent(0, false)

	// 获取内存使用情况
	memStats, _ := mem.VirtualMemory()

	// 获取磁盘使用情况
	diskStats, _ := disk.Usage("/")

	//tcp连接数
	netStats, _ := net.Connections("tcp")
	// 获取文件描述符
	//fileDescriptor, err := exec.Command("sh", "-c", "lsof -n | wc -l").Output()
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get file descriptor: %v", err)
	//}

	status, _ := getProgramStatus(pname)

	return &SystemMetrics{
		CPUUsage:       cpuPercent[0], // 取第一个元素
		MemoryUsage:    memStats.UsedPercent,
		DiskUsage:      diskStats.UsedPercent,
		TcpConnections: len(netStats),
		SuperVisor:     status,
	}, nil
}

func getProgramStatus(programName string) (string, error) {
	cmd := exec.Command("supervisorctl", "status", programName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(programName + " getProgramStatus error")
		return "", err
	}

	return parseSupervisordStatus(strings.TrimSpace(string(output))), nil
}

func parseSupervisordStatus(output string) string {
	// 将输出拆分为字段
	fields := strings.Fields(output)

	// 确保输出格式正确
	if len(fields) < 2 {
		return "UNKNOWN"
	}

	// 返回状态字段（例如 RUNNING、STOPPED 等）
	return fields[1]
}
