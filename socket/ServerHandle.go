package socket

import (
	"fmt"
)

func ServerHandle(socket *Socket){
	for{
		message := socket.Read()
		if message == ""{
			fmt.Printf("%vが切断されました。", socket.GetConn().RemoteAddr())
			break
		}
		fmt.Println("[受信]: " + message)
	}
}
func SendAll(message string){
	for _, socket := range sockets {
		socket.Send(message)
	}
}