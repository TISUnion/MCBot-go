package connect

import (
	"crypto/aes"

	. "github.com/TISUnion/MCBot-go/datatype"
)

//粘包处理
func (c *Connect) pushBuf(BufPool *[]byte, length int) {
	for (len(*BufPool) == 0) || (length > len(*BufPool)) {
		tempBuf := c.ReadAll().Bytes()
		*BufPool = append(*BufPool, tempBuf...)
	}
}

//启动aes加密通信
func (c *Connect) startAES() {
	//生成AES密钥
	sharedSecret := RandByte(16)
	c.sharedSecret = sharedSecret
	c.player.SharedSecret = sharedSecret
	//创建解密器
	aesBlock, err := aes.NewCipher([]byte(sharedSecret))
	if err != nil {
		return
	}
	iv := []byte(sharedSecret)
	aesDecrypter := NewCFB8Decrypter(aesBlock, iv)
	c.aesDecrypter = aesDecrypter
	//创建加密器
	if err != nil {
		return
	}
	aesEncrypter := NewCFB8Encrypter(aesBlock, iv)
	c.aesEncrypter = aesEncrypter
}

//循环获取服务端发来的数据
func (c *Connect) HandleBuf() {
	var BufPool []byte = []byte{}
	var length int
	for {
		c.pushBuf(&BufPool, length)
		length = VarIntInstance.DecodeByte(&BufPool)
		packet := &Packet{Length: length}
		c.pushBuf(&BufPool, length)
		packet.Data = CutBytes(&BufPool, length)
		packet.GetId()
		if packet.Id == 0x01 && c.IsOnline {
			//如果正版登陆成功，则启动aes加密通信（懒得用接口重构)
			c.startAES()
		}
		c.DealPacket(packet)
		if packet.Id == 0x01 && c.IsOnline {
			//设置aes加密启动标志位
			c.isAES = true
		}
	}

}
