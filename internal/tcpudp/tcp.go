package tcpudp

import (
	"fmt"
	"gstat/internal/protocol"
	"net"
	"time"
)

func Check(prot protocol.Protocol, host, port string) protocol.Result {
	target := fmt.Sprintf("%s:%s", host, port)
	res := protocol.Result{
		Target:   target,
		Protocol: protocol.HTTP,
	}

	timeout := 5 * time.Second
	connection, err := net.DialTimeout(string(prot), target, timeout)
	if err != nil {
		res.Reachable = false
		res.Message = err.Error()
		return res
	}

	defer connection.Close()
	res.Reachable = true
	return res
}
