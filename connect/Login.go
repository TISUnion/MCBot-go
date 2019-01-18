package connect

import (
	"bytes"

	. "github.com/TISUnion/MCBot-go/datatype"
)

func (c *Connect) LoginStart() {
	handshakeData := c.HandShake(2)
	c.Send(handshakeData)
	loginBuf := bytes.NewBuffer([]byte{0x00})
	StringInstance.Encode(c.player.Username, loginBuf)
	c.Send(loginBuf.Bytes())
}
