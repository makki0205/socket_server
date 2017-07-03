package main

import (
	"os"
	"github.com/makki0205/socket_server/socket"
	"fmt"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Println("指定された引数の数が間違っています。")
		os.Exit(1)
	}
	if os.Args[1] == "-l"{
		socket.ServerInit()
	}else{
		socket.ClientInit(os.Args[1])
	}
}