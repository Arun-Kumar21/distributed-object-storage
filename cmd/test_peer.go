package main

import (
	"time"

	"github.com/Arun-Kumar21/distributed-object-storage/internal/p2p"
)

func TestPeer() {
	node := &p2p.Node{Address: "localhost:8000"}
	go node.StartListening()

	time.Sleep(10 * time.Second)
}