package connect

import (
	. "MCBot-go/datatype"
)

//获取服务器信息
func (c *Connect) GetStatus() string {
	data := c.HandShake(1)
	c.Send(data)
	c.Send([]byte{0x00})
	ret := c.ReadAll()
	c.Close()
	// 数据包长度
	VarIntInstance.Decode(ret)
	// packetID
	VarIntInstance.Decode(ret)
	return StringInstance.Decode(ret)
}