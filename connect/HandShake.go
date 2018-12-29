package connect

import (
	"bytes"

	. "github.com/TISUnion/MCBot-go/datatype"
)

//对服务器握手发送数据进行封装
func (c *Connect) HandShake(NextState int) []byte {
	var buf bytes.Buffer
	if NextState == 2 {
		buf.Write([]byte{0x00})
		ProtocolVersion := c.GetProtocolVersion()
		VarIntInstance.Encode(ProtocolVersion, &buf)
	} else {
		buf.Write([]byte{0x00, 0x00})
	}
	StringInstance.Encode(c.Host, &buf)
	UnsignedShortInstance.Encode(c.Port, &buf)
	VarIntInstance.Encode(NextState, &buf)
	return buf.Bytes()
}
