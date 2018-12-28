package datatype

import (
	"bytes"
)

//VarInt是MC中的协议通信数据类型，最多5位
type VarInt struct{}

func (v *VarInt) Encode(b int, bs *bytes.Buffer) {
	bs.Write(v.Encode_(b))
}

func (*VarInt) Encode_(b int) []byte {
	var res []byte
	var temp byte
	for b != 0 {
		temp = byte(b & 0x7F)
		b >>= 7
		if b != 0 {
			temp |= 0x80
		}
		res = append(res, temp)
	}
	return res
}

func (*VarInt) Decode(bs *bytes.Buffer) int {
	var res int = 0
	var bytes_encountered byte = 0
	var read, value byte
	for ok := true; ok; ok = (read&0x80 != 0) {
		read, _ = bs.ReadByte()
		value = read & 0x7F
		res |= int(value) << (7 * bytes_encountered)
		bytes_encountered += 1
	}
	return res
}

var VarIntInstance *VarInt

func init() {
	VarIntInstance = &VarInt{}
}
