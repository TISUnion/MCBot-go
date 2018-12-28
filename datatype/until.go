package datatype

import (
	"bytes"
	"encoding/binary"
)

/*
 *该包主要是MC传输协议用的字节序列编码解码函数
 *具体可参考https://wiki.vg/Protocol#Data_types
 */

//帮助函数

//数字转大端字节流
func Num2Bytes(num interface{}) ([]byte, bool) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, num)
	if err != nil {
		return nil, false
	}
	return buf.Bytes(), true
}

//数字转字符串
func Num2String(num interface{}) (string, bool) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, num)
	if err != nil {
		return "", false
	}
	return buf.String(), true
}

//在发送的数据前面加上长度
func AddLength(data []byte) []byte {
	var res []byte
	res = VarIntInstance.Encode_(len(data))
	return append(res, data...)
}
