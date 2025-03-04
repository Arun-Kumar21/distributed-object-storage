package p2p

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
)

type Node struct {
	Address string
	Peers	map[string]bool
	PeerFile string
	Mutex	sync.Mutex
}


func (n *Node) StartListening() {
	listener, err := net.Listen("tcp4", n.Address)
	if err != nil {
		fmt.Println("Error in starting node:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Node is listening on ", n.Address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}		
		
		go n.handleConnection(conn)
	}
}

func (n *Node) handleConnection (conn net.Conn) {
	defer conn.Close()

	var peerAddr string;
	fmt.Fscanf(conn, "%s\n", peerAddr)
	n.AddPeer(peerAddr)

	for peer := range n.Peers {
		if peer != peerAddr {
			fmt.Println(conn, "%s\n", peer)
		}
	}
}

// Load Peer of a node
func (n *Node) LoadPeer () {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	data, err := os.ReadFile(n.PeerFile)
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return 
	}

	err = json.Unmarshal(data, &n.Peers)
	if err != nil {
		fmt.Println("Error reading peer file:", err)
		n.Peers = make(map[string]bool)
	}
}

// Save Peer of node to peer.json
func (n *Node) SavePeer () {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	data, err := json.MarshalIndent(n.Peers, "", "  ")
	if err != nil {
		fmt.Println("Error creating json:", err)
		return
	}

	err = os.WriteFile(n.PeerFile, data, 0644)
	if err != nil {
		fmt.Println("Error while writing file:", err)
		return 
	}
}

// Add a peer to node
func (n *Node) AddPeer (peer string) {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	if _, exists := n.Peers[peer]; !exists {
		n.Peers[peer] = true
		fmt.Println("New peer added:", peer)
		n.SavePeer()
	}
}


func (n *Node) ConnectToPeer (peerAddr string) {
	conn, err := net.Dial("tcp4", peerAddr)
	if err != nil {
		fmt.Println("Failed to connect peer:", err)
		return
	}

	defer conn.Close()
	fmt.Println(conn, "%s\n", n.Address)

	var newPeer string;
	for {
		_, err = fmt.Fscanf(conn, "%s\n", &newPeer)
		if err != nil {
			break
		}

		n.AddPeer(newPeer)
	}
}