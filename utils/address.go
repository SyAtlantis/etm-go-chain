package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/base58check"
	"golang.org/x/crypto/ripemd160"
	"strings"
)

type Address struct {
}

func (addr *Address) IsAddress(address string) bool {
	if !strings.HasPrefix(address, "A") {
		return false
	}

	_, err := base58check.Decode(address[1:])
	if err != nil {
		return false
	}

	return true
}

func (addr *Address) GenerateAddress(publicKey []byte) string {
	sha256Inst := sha256.New()
	sha256Inst.Reset()
	sha256Inst.Write(publicKey)
	sha256Bytes := sha256Inst.Sum(nil)

	ripemd160Inst := ripemd160.New()
	ripemd160Inst.Reset()
	ripemd160Inst.Write(sha256Bytes)
	ripemd160Bytes := ripemd160Inst.Sum(nil)

	address, err := base58check.Encode("", fmt.Sprintf("%x", ripemd160Bytes))
	if err != nil {
		return ""
	}

	address = "A" + address
	return address

}

func (addr *Address) GenerateAddressByStr(publicKey string) (string, error) {
	pub, err := hex.DecodeString(publicKey)
	if err != nil {
		return "", err
	}
	address := addr.GenerateAddress(pub)
	return address, nil
}
