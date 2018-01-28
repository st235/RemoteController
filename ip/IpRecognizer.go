package ip

import (
	"net"
	"os"
	"fmt"
)

func GetMachineIp(port string) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return fmt.Sprintf("http://%s%s/", ipnet.IP.String(), port)
			}
		}
	}

	return fmt.Sprintf("http://localhost%s/", port)
}
