package utils

import (
	"bytes"
	"golang.org/x/crypto/ed25519"
	"reflect"
)

type Ed struct {
}

type Keypair struct {
	PublicKey  []byte
	PrivateKey []byte
}

func (k *Keypair) IsEmpty() bool {
	return k == nil || reflect.DeepEqual(k, Keypair{})
}

// 构造keypair
func (ed *Ed) MakeKeypair(hash []byte) Keypair {
	pub, pri, err := ed25519.GenerateKey(bytes.NewReader(hash))
	if err != nil {
		return Keypair{}
	}

	return Keypair{
		PublicKey:  pub,
		PrivateKey: pri,
	}
}

// 创建签名
func (ed *Ed) Sign(hash []byte, keypair Keypair) []byte {
	return ed25519.Sign(keypair.PrivateKey, hash)
}

// 验证签名
func (ed *Ed) Verify(hash, signatureBuffer []byte, publicKeyBuffer []byte) bool {
	return ed25519.Verify(publicKeyBuffer, hash, signatureBuffer)
}
