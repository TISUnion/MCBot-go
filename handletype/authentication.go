package handletype

import (
	"bytes"

	. "github.com/TISUnion/MCBot-go/datatype"
)

type Authentication struct {
}

func (*Authentication) Handle(pk *Packet, pl *Player) []byte {
	buf := bytes.NewBuffer(*pk.Data)
	//获取serverid
	serverid := StringInstance.Decode(buf)
	//获取rsa参数
	publicKey := StringInstance.Decode(buf)
	verifyToken := StringInstance.Decode(buf)
	//使用服务器的公钥加密参数
	secret_verify_token, ok := RsaEncrypt([]byte(publicKey), []byte(verifyToken))
	if !ok {
		return []byte{}
	}
	sharedSecret, ok := RsaEncrypt([]byte(publicKey), []byte(pl.SharedSecret))
	if !ok {
		return []byte{}
	}
	//封装数据包
	resbuf := bytes.NewBuffer([]byte{0x01})
	StringInstance.Encode(string(sharedSecret), resbuf)
	StringInstance.Encode(string(secret_verify_token), resbuf)
	//发出session请求
	sessionTextbuf := bytes.NewBufferString(serverid)
	sessionTextbuf.Write([]byte(pl.SharedSecret))
	sessionTextbuf.Write([]byte(publicKey))
	serverHash := AuthDigest(sessionTextbuf.String())
	if !pl.SetSession(serverHash) {
		return []byte{}
	}
	return resbuf.Bytes()
}
