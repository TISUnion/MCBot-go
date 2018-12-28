package connect

import (
	. "MCBot-go/datatype"
	"bytes"
	"net"
)

var ProtocolVersion int

type Connect struct {
	Conn net.Conn
	Host string
	Port uint16
}

//向服务器发送数据
func (c *Connect) Send(bs []byte) {
	stream := AddLength(bs)
	c.Conn.Write(stream)
}

//读取服务器发送过来的数据
func (c *Connect) ReadAll() *bytes.Buffer {
	data := make([]byte, 5000)
	n, _ := c.Conn.Read(data)
	return bytes.NewBuffer(data[:n])
}

//释放资源
func (c *Connect) Close() {
	c.Conn.Close()
}
