package handletype

import (
	. "github.com/TISUnion/MCBot-go/datatype"
)

type Keeplive struct {
}

func (*Keeplive) Handle(pk *Packet, pl *Player) []byte {
	res := append([]byte{0x00, 0x0E}, *pk.Data...)
	return res
}
