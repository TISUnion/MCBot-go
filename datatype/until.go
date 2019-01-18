package datatype

import (
	"bytes"
	"compress/zlib"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

/*
 *该包主要是处理MC传输协议用的字节序列的帮助函数
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

//剪切[]byte
func CutBytes(bt *[]byte, length int) *[]byte {
	res := (*bt)[:length]
	if len(*bt) > length {
		*bt = (*bt)[length:]
	} else {
		*bt = []byte{}
	}
	return &res
}

//随机生成对称加密ase密钥
func RandByte(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterbyte := []byte("abcdefghijklmnopqrstuvwxyz123467890")
	b := make([]byte, n)
	for i := range b {
		b[i] = letterbyte[rand.Intn(len(letterbyte))]
	}
	return string(b)
}

//发请求函数
func _request(url string, param string) (string, bool) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Sprintf("请求接口：%s出错，因为：%v", url, err)
		return "", false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	return string(body), true
}

//向mojong获取验证刷新token等操作
func TokenRequest(apiType string, param string) (string, bool) {
	url := fmt.Sprintf("https://authserver.mojang.com/%s", apiType)
	return _request(url, param)
}

//向mojong发送session
func SessionRequest(param string) (string, bool) {
	url := "https://sessionserver.mojang.com/session/minecraft/join"
	return _request(url, param)
}

//使用服务器发的公钥进行rsa加密
func RsaEncrypt(publicKey []byte, orignData []byte) ([]byte, bool) {
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(publicKey)
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	res, err := rsa.EncryptPKCS1v15(crand.Reader, pub, orignData)
	if err != nil {
		fmt.Println(err)
		return []byte{}, false
	}
	return res, true
}

/**
 *生成登陆hash
 *参考代码：https://gist.github.com/toqueteos/5372776
 */
func AuthDigest(s string) string {
	// little endian
	twosComplement := func(p []byte) []byte {
		carry := true
		for i := len(p) - 1; i >= 0; i-- {
			p[i] = byte(^p[i])
			if carry {
				carry = p[i] == 0xff
				p[i]++
			}
		}
		return p
	}
	h := sha1.New()
	io.WriteString(h, s)
	hash := h.Sum(nil)

	// Check for negative hashes
	negative := (hash[0] & 0x80) == 0x80
	if negative {
		hash = twosComplement(hash)
	}

	// Trim away zeroes
	res := strings.TrimLeft(fmt.Sprintf("%x", hash), "0")
	if negative {
		res = "-" + res
	}

	return res
}

//进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

//进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}
