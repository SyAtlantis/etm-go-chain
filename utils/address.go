package utils

import (
	"fmt"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"github.com/anaskhan96/base58check"
	"strings"
)

type Address struct {
	
}

func (addr * Address) IsAddress(address string)  bool{
	if  !strings.HasPrefix(address,"A"){
		return false
	}
	
	_,err := base58check.Decode(address[1:])
	if err != nil{
		return false
	}
	
	return true
}

func (addr *Address) GenerateAddresss(publicKey []byte) string {
	//hash :=sha256.Sum256(publicKey)
	
	//ripemd := ripemd160.New()
	//ripemd.Reset()
	//ripemd.Write(hash)
	//ripemdBytes := ripemd.Sum(nil)
	
	sha256Inst := sha256.New()
	sha256Inst.Reset()
	sha256Inst.Write([]byte(publicKey))
	sha256Bytes := sha256Inst.Sum(nil)
	
	ripemd160Inst := ripemd160.New()
	ripemd160Inst.Reset()
	ripemd160Inst.Write([]byte(sha256Bytes))
	ripemd160Bytes := ripemd160Inst.Sum(nil)
	
	address, err := base58check.Encode("", fmt.Sprintf("%x", ripemd160Bytes))
	if err != nil {
		return ""
	}
	
	address = "A" + address
	return address

}