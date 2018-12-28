package datatype

import "bytes"

type Bool struct{}

func (*Bool) Encode(b bool, bs *bytes.Buffer) {
	if b {
		bs.WriteByte(1)
	} else {
		bs.WriteByte(0)
	}
}

func (*Bool) Decode(bs *bytes.Buffer) bool {
	var res bool
	if t, _ := bs.ReadByte(); t == 0 {
		res = false
	} else {
		res = true
	}
	return res
}
