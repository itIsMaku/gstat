package tcpudp

import (
	"gstat/internal/protocol"
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		name      string
		proto     protocol.Protocol
		target    string
		reachable bool
	}{
		{"tcp-80-reachable", protocol.TCP, "rcore.cz:80", true},
		{"tcp-801-nonreachable", protocol.TCP, "rcore.cz:801", false},
		{"udp-80-alwaysreachable", protocol.UDP, "rcore.cz:80", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Check(tt.proto, tt.target)
			if result.Reachable != tt.reachable {
				t.Fatalf("Test failed, expected reachable=%v, got reachable=%v", tt.reachable, result.Reachable)
			}
		})
	}
}
