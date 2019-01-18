package handletype

import (
	. "github.com/TISUnion/MCBot-go/datatype"
)

func DistinguishPacket(pk *Packet, pl *Player) []byte {
	var PacketMap map[int]HandleMachine = map[int]HandleMachine{
		0x21: &Keeplive{},
		0x02: &PlayerMessage{},
		0x01: &Authentication{},
		// 0x32: &UpdatePosition{},e
	}
	if handleMachine, ok := PacketMap[pk.Id]; ok {
		return handleMachine.Handle(pk, pl)
	}
	return []byte{}
}
