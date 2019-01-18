package handletype

import (
	. "github.com/TISUnion/MCBot-go/datatype"
)

type UpdatePosition struct {
}

func (*UpdatePosition) Handle(pk *Packet, pl *Player) []byte {
	x := DoubleInstance.DecodeByte(pk.Data)
	y := DoubleInstance.DecodeByte(pk.Data)
	z := DoubleInstance.DecodeByte(pk.Data)
	yaw := FloatInstance.DecodeByte(pk.Data)
	pitch := FloatInstance.DecodeByte(pk.Data)
	data := append([]byte{0x00, 0x11}, DoubleInstance.Encode(x)...)
	data = append(data, DoubleInstance.Encode(y)...)
	data = append(data, DoubleInstance.Encode(z)...)
	data = append(data, FloatInstance.Encode(yaw)...)
	data = append(data, FloatInstance.Encode(pitch)...)
	data = append(data, 0x01)
	return data
}
