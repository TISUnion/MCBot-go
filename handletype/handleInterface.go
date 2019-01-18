package handletype

import (
	. "github.com/TISUnion/MCBot-go/datatype"
)

type HandleMachine interface {
	Handle(*Packet, *Player) []byte
}
