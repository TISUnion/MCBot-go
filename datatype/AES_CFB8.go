/**
 * aes/cfb8加解密
 * 代码来源：https://github.com/kierdavis/cfb8
 *
 */

package datatype

import (
	"crypto/cipher"
)

type Cfb8 struct {
	b       cipher.Block
	sr      []byte
	srEnc   []byte
	srPos   int
	decrypt bool
}

func NewCFB8Encrypter(block cipher.Block, iv []byte) cipher.Stream {
	if len(iv) != block.BlockSize() {
		panic("cfb8.NewEncrypter: IV length must equal block size")
	}
	return newCFB8(block, iv, false)
}

func NewCFB8Decrypter(block cipher.Block, iv []byte) cipher.Stream {
	if len(iv) != block.BlockSize() {
		panic("cfb8.NewEncrypter: IV length must equal block size")
	}
	return newCFB8(block, iv, true)
}

func newCFB8(block cipher.Block, iv []byte, decrypt bool) cipher.Stream {
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return nil
	}

	x := &Cfb8{
		b:       block,
		sr:      make([]byte, blockSize*4),
		srEnc:   make([]byte, blockSize),
		srPos:   0,
		decrypt: decrypt,
	}
	copy(x.sr, iv)

	return x
}

func (x *Cfb8) XORKeyStream(dst, src []byte) {
	blockSize := x.b.BlockSize()

	for i := 0; i < len(src); i++ {
		x.b.Encrypt(x.srEnc, x.sr[x.srPos:x.srPos+blockSize])

		var c byte
		if x.decrypt {
			c = src[i]
			dst[i] = c ^ x.srEnc[0]
		} else {
			c = src[i] ^ x.srEnc[0]
			dst[i] = c
		}

		x.sr[x.srPos+blockSize] = c
		x.srPos++

		if x.srPos+blockSize == len(x.sr) {
			copy(x.sr, x.sr[x.srPos:])
			x.srPos = 0
		}
	}
}
