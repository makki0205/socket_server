package socket


import (
	"fmt"
	"net"
	"github.com/makki0205/socket_server/file"
	"bufio"
	"os"
	"strconv"
	"strings"
)
var sockets []*Socket

func ServerInit() {
	listener, err := net.Listen("tcp", "0.0.0.0:2000")
	if err != nil {
		fmt.Printf("Listen error: %s\n", err)
		return
	}
	addrs, err := net.InterfaceAddrs()
	for _, addr := range addrs {
		if -1 != strings.Index(addr.String(), "192.") {
			fmt.Println(addr.String(), "is listen")
		}
	}
	fmt.Println()
	go AcceptHandle(listener)
	ServerSendHandle()
}
func AcceptHandle(listener net.Listener){
	for  {
		conn, _ := listener.Accept()
		fmt.Printf("\n%vが接続しました。\n", conn.RemoteAddr())
		socket := NewSocketManager(conn)
		go ServerHandle(socket)
		sockets = append(sockets, socket)
	}
}
func ServerSendHandle(){
	fRep := file.NewfileRepository("assets")
	s := bufio.NewScanner(os.Stdin)
	fRep.Print()
	for  {
		print("番号を入力してください >")
		s.Scan()
		index, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println("数値を入力してください")
			continue
		}
		if index > fRep.FileCount {
			fmt.Println("ファイルが見つかりません!!!!")
			continue
		}
		json_string :=fRep.GetFileString(index)
		SendAll(json_string)
	}
}
//func ServerSendHandle(){
//	var buf string
//	stdin := bufio.NewScanner(os.Stdin)
//	for stdin.Scan() {
//		if err := stdin.Err(); err != nil {
//			fmt.Fprintln(os.Stderr, err)
//		}
//		buf = stdin.Text()
//		if buf == "exit"{
//			break
//		}
//		SendAll(buf)
//	}
//}