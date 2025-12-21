package protocol

import "fmt"

type Protocol string

const (
	HTTP = "http"
	TCP  = "tcp"
	UDP  = "udp"
)

type Result struct {
	Target    string
	Protocol  Protocol
	Reachable bool
	Message   string
}

func (result Result) String() string {
	return fmt.Sprintf(`Target: %s
Protocol: %s
Reachable: %v
Message: %s`, result.Target, result.Protocol, result.Reachable, result.Message)
}
