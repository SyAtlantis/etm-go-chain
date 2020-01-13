package mian

import (
	"crypto/sha256"
	"etm-go-chain/utils"
	"fmt"
)

func main() {
	hash := sha256.Sum256([]byte("real rally sketch sorry place parrot typical cart stone mystery age nominee"))
	fmt.Printf("%x\n", hash)
	
	ed := utils.Ed{}
	keypair := ed.MakeKeypair(hash[:])
	fmt.Println(keypair)
	fmt.Println(fmt.Sprintf("%x", keypair.PublicKey))
	fmt.Println(fmt.Sprintf("%x", keypair.PrivateKey))
	
	//sign := ed.Sign(hash[:],keypair)
	//fmt.Println(fmt.Sprintf("%x", sign))
	//
	//fmt.Println(ed.Verify(hash[:],sign,keypair.PublicKey))
	
	//test address
	address := utils.Address{}
	addr := address.GenerateAddress(keypair.PublicKey)
	fmt.Println(addr)
	fmt.Println(address.IsAddress(addr))
	
}
