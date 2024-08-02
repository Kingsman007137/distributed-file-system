package main

import (
	"fmt"
	"log"

	"github.com/Kingsman007137/distributed-file-system/p2p"
)

func OnPeer(p p2p.Peer) error {
	p.Close()
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    "localhost:3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer: 	   OnPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Println("Received message: ", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	// block forever
	select {}
}
