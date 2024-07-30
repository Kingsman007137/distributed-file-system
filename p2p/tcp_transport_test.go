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
	listenAddress := "localhost:8080"
	transport := NewTCPTransport(listenAddress)
	// 判断是否创建成功
	assert.Equal(t, transport.listenAddress, listenAddress)

	// 测试 ListenAndAccept 方法
	assert.Nil(t, transport.ListenAndAccept())
	// fmt.Println(transport.ListenAndAccept())
}