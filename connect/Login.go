package connect

import (
	"bytes"
	"fmt"

	. "github.com/TISUnion/MCBot-go/datatype"
)

func (c *Connect) LoginStart() {
	handshakeData := c.HandShake(2)
	c.Send(handshakeData)
	loginBuf := bytes.NewBuffer([]byte{0x00})
	StringInstance.Encode("lightbrother0_0", loginBuf)
	c.Send(loginBuf.Bytes())
	ret := c.ReadAll()
	fmt.Println(ret.Bytes())
}

func (c *Connect) GetNameAndId() {
	ret := c.ReadAll()
	VarIntInstance.Decode(ret)
	VarIntInstance.Decode(ret)
	VarIntInstance.Decode(ret)
	id := StringInstance.Decode(ret)
	name := StringInstance.Decode(ret)
	fmt.Println(id, name)
	fmt.Println([]byte(id), []byte(name))
}
