package utils

import (
	"encoding/binary"
	"fmt"
	"net"
)

func LookUpLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// 过滤掉 loopback 地址和非 IPV4 地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}

	return "", fmt.Errorf("no valid IP address found")
}

func IpToUint32(ipStr string) uint32 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func GetMACBasedIdentifier() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range interfaces {
		if iface.Flags&net.FlagLoopback == 0 && iface.HardwareAddr != nil {
			mac := iface.HardwareAddr.String()
			//hash := sha1.New()
			//hash.Write([]byte(mac))
			//return hex.EncodeToString(hash.Sum(nil)), nil
			return mac, nil
		}
	}
	return "", fmt.Errorf("no valid MAC address found")
}
