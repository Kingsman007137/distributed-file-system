package p2p

import "net"

// RPC holds arbitrary data that is being sent over the each
// transport bewteen two nodes in the network.
type RPC struct {
	Playload []byte
	From net.Addr
}