package connect

import (
	. "MCBot-go/datatype"
	"bytes"
)

//对服务器握手发送数据进行封装
func (c *Connect) HandShake(NextState int) []byte {
	var buf bytes.Buffer
	buf.Write([]byte{0x00, 0x00})
	if NextState == 1 {
		VarIntInstance.Encode(ProtocolVersion, &buf)
	}
	StringInstance.Encode(c.Host, &buf)
	UnsignedShortInstance.Encode(c.Port, &buf)
	VarIntInstance.Encode(NextState, &buf)
	return buf.Bytes()
}
