package p2p

import (
	"fmt"
	"net"
	"sync"
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

type TCPTransportOpts struct {
	ListenAddr string
	HandshakeFunc HandshakeFunc
	Decoder    Decoder
}

type TCPTransport struct {
	// 属性先定义成不可导出的，以后有需要改
	TCPTransportOpts
	listener      net.Listener

	mutex  sync.Mutex
	peers  map[net.Addr]Peer
}

// NewTCPTransport creates a new TCP transport
func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
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
	// here is outbound connection, because we dial it
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Println("Error shaking hands with peer: ", err)
		return
	}

	// Read Loop
	msg := &Message{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Println("Error decoding message: ", err)
			break
		}

		fmt.Println("Received message: ", string(msg.Playload))
	}
}