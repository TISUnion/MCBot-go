package datatype

import "bytes"

type String struct{}

/*
 * Go中字符串默认Utf-8
 * 获取字节序列只需 []byte(Str)
 * 写入字符串需要先写入字符串的长度
 */
func (*String) Encode(Str string, bs *bytes.Buffer) {
	length := VarIntInstance.Encode_(len(Str))
	bs.Write(length)
	bs.Write([]byte(Str))
}

/*
 * 分成2组，第一组为字符串长度
 * 第二组为字符串字节序列
 */
func (*String) Decode(bs *bytes.Buffer) string {
	length := VarIntInstance.Decode(bs)
	var res []byte
	var temp byte
	for i := 0; i < length; i++ {
		temp, _ = bs.ReadByte()
		res = append(res, temp)
	}
	return string(res)
}

var StringInstance *String

func init() {
	StringInstance = &String{}
}
