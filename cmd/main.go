package main

import (
	"fmt"

	"github.com/Arun-Kumar21/distributed-object-storage/internal/p2p"
)

func main() {
	node1 := &p2p.Node{
		Address: "localhost:9001", 
		Peers: make(map[string]bool), 
		PeerFile: "peer1.json",
	}

	node2 := &p2p.Node{
		Address: "localhost:9002",
		Peers: make(map[string]bool),
		PeerFile: "peer2.json",
	}

	node1.LoadPeer()
	node2.LoadPeer()

	go node1.StartListening()
	go node2.StartListening()

	node2.ConnectToPeer("localhost:9001")

	fmt.Println("Node1 Peers:", node1.Peers)
	fmt.Println("Node2 Peers:", node2.Peers)
}
