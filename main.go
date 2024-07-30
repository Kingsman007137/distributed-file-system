package main

import (
	"fmt"
	"log"

	"github.com/Kingsman007137/distributed-file-system/p2p"
)

func main() {
	tr := p2p.NewTCPTransport("localhost:3000")
	fmt.Println(tr)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	// block forever
	select {}
}