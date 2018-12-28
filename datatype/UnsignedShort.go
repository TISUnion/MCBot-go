package datatype

import "bytes"

//UnsignedShort是常规的16位2byte无符号整形
type UnsignedShort struct{}

func (*UnsignedShort) Encode(us uint16, bs *bytes.Buffer) bool {
	if v, ok := Num2Bytes(us); ok {
		bs.Write(v)
		return true
	}
	return false
}

func (*UnsignedShort) Decode(bs *bytes.Buffer) uint16 {
	var res uint16
	bt1, _ := bs.ReadByte()
	bt2, _ := bs.ReadByte()
	res |= (uint16(bt1) << 8) | (uint16(bt2))
	return res
}

var UnsignedShortInstance *UnsignedShort

func init() {
	UnsignedShortInstance = &UnsignedShort{}
}
