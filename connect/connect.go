package connect

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"time"

	. "github.com/TISUnion/MCBot-go/datatype"

	"github.com/tidwall/gjson"
)

var ProtocolVersion int

type Connect struct {
	Conn     net.Conn
	Host     string
	Port     uint16
	IsOnline bool
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

//获取协议版本号
func (c *Connect) GetProtocolVersion() int {
	if ProtocolVersion == 0 {
		address := fmt.Sprintf("%s:%d", c.Host, c.Port)
		tempConn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		tempC := &Connect{tempConn, c.Host, c.Port, c.IsOnline}
		defer tempC.Close()
		ProtocolVersion = int(gjson.Get(tempC.GetStatus(), "version.protocol").Int())
	}
	return ProtocolVersion
}

//释放资源
func (c *Connect) Close() {
	c.Conn.Close()
}
