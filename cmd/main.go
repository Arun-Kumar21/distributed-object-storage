package main

import (
	"fmt"
	"time"

	"github.com/Arun-Kumar21/distributed-object-storage/internal/p2p"
)

func main() {
	node := &p2p.Node{Address: "localhost:9000"}
	go node.StartNode()

	time.Sleep(2 * time.Second)

	fmt.Println("Sending message...")
	p2p.SendMessage("localhost:9000", "Hello peer")
}
