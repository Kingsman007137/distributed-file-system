package p2p

import (
	// "fmt"
	"testing"

	// run "go get github.com/stretchr/testify"
	// & "go mod tidy" before running the test
	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	// 创建一个 TCPTransport 实例
	opts := TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       GOBDecoder{},
	}
	tr := NewTCPTransport(opts)
	assert.Equal(t, tr.ListenAddr, ":3000")

	assert.Nil(t, tr.ListenAndAccept())
	// fmt.Println(transport.ListenAndAccept())
}