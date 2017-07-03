package socket

import (
	"fmt"
	"net"
)

func ClientInit(ip string){
	conn, err := net.Dial("tcp", ip + ":1234")
	if err != nil {
		fmt.Printf("Dial error: %s\n", err)
		return
	}
	defer conn.Close()
	sock := NewSocketManager(conn)
	go ReadHandle(sock)
	SendHandle(sock)
}