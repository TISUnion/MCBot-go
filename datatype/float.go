package datatype

import (
	"bytes"
	"encoding/binary"
)

type Float struct {
}

func (*Float) DecodeByte(bs *[]byte) float32 {
	var res float32
	buf := bytes.NewReader((*bs)[:4])
	*bs = (*bs)[4:]
	binary.Read(buf, binary.BigEndian, &res)
	return res
}

func (*Float) Encode(dou float32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, dou)
	return buf.Bytes()
}

var FloatInstance *Float

func init() {
	FloatInstance = &Float{}
}
