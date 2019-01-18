package connect

import (
	"bytes"
	"crypto/cipher"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	. "github.com/TISUnion/MCBot-go/datatype"
	"github.com/TISUnion/MCBot-go/handletype"

	"github.com/tidwall/gjson"
)

var ProtocolVersion int

type Connect struct {
	Conn         net.Conn
	Host         string
	Port         uint16
	IsOnline     bool
	player       *Player
	sharedSecret string
	isAES        bool
	aesDecrypter cipher.Stream
	aesEncrypter cipher.Stream
}

//向服务器发送数据
func (c *Connect) Send(bs []byte) {
	if len(bs) != 0 {
		stream := AddLength(bs)
		if c.isAES && c.IsOnline {
			stream = c.aesCbf8Encode(stream)
		}
		c.Conn.Write(stream)
	}
}

//处理获取的数据包
func (c *Connect) DealPacket(pk *Packet) {
	if !c.IsOnline && (pk.Id == 0x01) {
		return
	}
	c.Send(handletype.DistinguishPacket(pk, c.player))
}

//读取服务器发送过来的数据
func (c *Connect) ReadAll() *bytes.Buffer {
	data := make([]byte, 1000)
	n, _ := c.Conn.Read(data)
	data = data[:n]
	if n == 0 {
		return bytes.NewBuffer(data)
	}
	if c.isAES {
		data = c.aesCbf8Decode(data)
	}
	return bytes.NewBuffer(data)
}

func debug(data []byte, st string) {
	f, _ := os.OpenFile("test3", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	defer f.Close()
	str := "["
	for _, v := range data {
		str = fmt.Sprintf("%s\\x%x", str, v)
	}
	f.WriteString(st + str + "]")
}

//获取协议版本号
func (c *Connect) GetProtocolVersion() int {
	if ProtocolVersion == 0 {
		address := fmt.Sprintf("%s:%d", c.Host, c.Port)
		tempConn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		tempC := &Connect{Conn: tempConn, Host: c.Host, Port: c.Port}
		defer tempC.Close()
		ProtocolVersion = int(gjson.Get(tempC.GetStatus(), "version.protocol").Int())
	}
	return ProtocolVersion
}

//释放资源
func (c *Connect) Close() {
	if c.player != nil {
		c.player.Signout()
		c.player.Invalidate()
	}
	c.Conn.Close()
}

//开始连接并运行
func (c *Connect) Start() {
	c.LoginStart()
	c.HandleBuf()
}

//AES解密
func (c *Connect) aesCbf8Decode(secretData []byte) []byte {
	res := make([]byte, len(secretData))
	c.aesDecrypter.XORKeyStream(res, secretData)
	return res
}

//AES加密
func (c *Connect) aesCbf8Encode(originData []byte) []byte {
	res := make([]byte, len(originData))
	c.aesEncrypter.XORKeyStream(res, originData)
	return res
}
