package p2p

import (
	"fmt"
	"net"
)

// TCPPeer represents a remote node over a TCP established connection.
type TCPPeer struct {
	conn net.Conn
	// outbound 表示是否是主动发起的连接
	// if we accept and retrieve connection, outbound is false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// Close implements the Peer interface, closes the connection to the remote node.
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddr string
	HandshakeFunc HandshakeFunc
	Decoder    Decoder
	// OnPeer is a callback function that is called when a new peer is connected
	// if the function returns an error, the connection is closed.
	OnPeer   func(Peer) error
}

type TCPTransport struct {
	// 属性先定义成不可导出的，以后有需要改
	TCPTransportOpts
	listener net.Listener
	// rpcch is a channel that is used to send and receive RPC messages
	rpcch    chan RPC
}

// NewTCPTransport creates a new TCP transport
func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch: 			  make(chan RPC),
	}
}

// Consume implements the Transport interface, which returns a read-only channel
// for reading the imcoming messages received from another peer in the network.
func (t *TCPTransport) Consume () <-chan RPC {
	return t.rpcch
}

// "If func is more important, then put it more above"
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err)
		}

		fmt.Println("new incoming connection: ", conn)
		go t.handleConnection(conn)
	}
}

// A method - a function with a special receiver argument.
func (t *TCPTransport) handleConnection(conn net.Conn) {
	var err error

	defer func ()  {
		fmt.Println("Dropping peer connection: ", err)
		conn.Close()
	}()

	// here is outbound connection, because we dial it
	peer := NewTCPPeer(conn, true)

	if err = t.HandshakeFunc(peer); err != nil {
		fmt.Println("Error shaking hands with peer!")
		return
	}

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			fmt.Println("Error calling OnPeer!")
			return
		}
	}

	// Read Loop
	rpc := RPC{}
	for {
		if err = t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Println("Error decoding message!")
			break
		}

		rpc.From = conn.RemoteAddr()
		t.rpcch <- rpc
	}
}