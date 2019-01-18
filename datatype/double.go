package datatype

import (
	"bytes"
	"encoding/binary"
)

type Double struct {
}

func (*Double) DecodeByte(bs *[]byte) float64 {
	var res float64
	buf := bytes.NewReader((*bs)[:8])
	*bs = (*bs)[8:]
	binary.Read(buf, binary.BigEndian, &res)
	return res
}

func (*Double) Encode(dou float64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, dou)
	return buf.Bytes()
}

var DoubleInstance *Double

func init() {
	DoubleInstance = &Double{}
}
