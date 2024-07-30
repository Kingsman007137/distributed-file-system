package p2p

import "net"

// Message holds arbitrary data that is being sent over the each
// transport bewteen two nodes in the network.
type Message struct {
	Playload []byte
	From net.Addr
}