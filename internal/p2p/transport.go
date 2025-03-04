package p2p

import (
	"fmt"
	"net"
)

func SendMessage(peerAddr string, message string) {
	conn, err := net.Dial("tcp4", peerAddr)
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}


	defer conn.Close()

	if conn == nil {
		fmt.Println("Connection is nil")
		return
	}

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Send failed:", err)
		return 
	}
	fmt.Println("Message send to ", peerAddr)
}