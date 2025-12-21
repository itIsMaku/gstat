package protocol

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
