package testutil

import "github.com/crgimenes/go-osc"

type Client interface {
	Send(packet osc.Packet) error
	SetLocalAddr(ip string, port int) error
}
