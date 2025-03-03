package p2p

import (
	"fmt"
	"net"
)

type Node struct {
	Address string
}

func (n *Node) StartNode() {
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
		
		go handleConnection(conn)
	}
}

func handleConnection (conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Received Message:", string(buffer[:n]))
}