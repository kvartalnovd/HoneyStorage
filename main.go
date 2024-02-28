package main

import (
	"log"

	"github.com/kvartalnovd/foreverstore/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	// buf := new(bytes.Buffer)
	// for {
	// 	n, _ := conn.Read(buf)
	// 	// msg := buf[:n]
	// }

	select {}
}
