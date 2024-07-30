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

type TCPTransport struct {
	// 属性先定义成不可导出的，以后有需要改
	listenAddress string
	listener      net.Listener

	mutex  sync.Mutex
	peers  map[net.Addr]Peer
}

// NewTCPTransport creates a new TCP transport
func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress,
	}
}

// "If func is more important, then put it more above"
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
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
		go t.handleConnection(conn)
	}
}

// A method - a function with a special receiver argument.
func (t *TCPTransport) handleConnection(conn net.Conn) {
	// here is inbound connection, because we dial it
	peer := NewTCPPeer(conn, true)
	fmt.Println("Handling connection from: ", peer)
}