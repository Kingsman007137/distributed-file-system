package p2p

// Define custom handshake operation functions
type HandshakeFunc func(Peer) error

// Do nothing for the handshake, no need handshake
func NOPHandshakeFunc(Peer) error {
	return nil
}