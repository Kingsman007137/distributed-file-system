package p2p

// 先定义空接口，以后有需要再添加方法和属性
// Peer is an interface that represents a remote node in the network.
type Peer interface {
	Close() error
}

// Transport is anything that handles the communication
// between the nodes in the network. This can be of the
// form (TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}