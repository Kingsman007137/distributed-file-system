package p2p

import (
	"fmt"
	"net"
	"sync"
)

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
			go t.handleConnection(conn)
		}
	}
}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	fmt.Println("Handling connection from: ", conn.RemoteAddr())
}