package tcpudp

import (
	"gstat/internal/protocol"
	"net"
	"time"
)

func Check(prot protocol.Protocol, target string) protocol.Result {
	res := protocol.Result{
		Target:   target,
		Protocol: prot,
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
