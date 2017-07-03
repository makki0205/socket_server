package socket

import (
	"net"
)

type Socket struct {
	conn net.Conn
}

func NewSocketManager(conn net.Conn) *Socket{
	return &Socket{conn: conn}

}
func (self *Socket)GetConn() net.Conn{
	return self.conn
}
func (this *Socket)Send(msg string){
	this.conn.Write([]byte(msg))
}

func (self *Socket)Read() string{
	buf := make([]byte, 5000)
	n, err := self.conn.Read(buf)
	if err != nil {
		return ""
	}
	return string(buf[:n])
}

func (self *Socket)Close(){
	self.conn.Close()
}
