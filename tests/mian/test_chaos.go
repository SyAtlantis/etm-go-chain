package mian

import (
	"etm-go-chain/utils"
	"fmt"
	"crypto/sha256"
)

func getDelegateIndex(id string, slot int64, limit int) int {
	hash := sha256.Sum256([]byte(id))
	h := fmt.Sprintf("%x", hash)
	//fmt.Println(h, len(h))
	
	//hex := map[byte][]byte{
	//	'0': []byte{0, 0, 0, 0},
	//	'1': []byte{0, 0, 0, 1},
	//	'2': []byte{0, 0, 1, 0},
	//	'3': []byte{0, 0, 1, 1},
	//	'4': []byte{0, 1, 0, 0},
	//	'5': []byte{0, 1, 0, 1},
	//	'6': []byte{0, 1, 1, 0},
	//	'7': []byte{0, 1, 1, 1},
	//	'8': []byte{1, 0, 0, 0},
	//	'9': []byte{1, 0, 0, 1},
	//	'a': []byte{1, 0, 1, 0},
	//	'b': []byte{1, 0, 1, 1},
	//	'c': []byte{1, 1, 0, 0},
	//	'd': []byte{1, 1, 0, 1},
	//	'e': []byte{1, 1, 1, 0},
	//	'f': []byte{1, 1, 1, 1},
	//}
	//var hash256 = [256]byte{}
	//hh := []byte(h)
	//fmt.Println(fmt.Sprintf("%b", h))
	//for i := 0; i < len(hh); i++ {
	//	//base, _ := strconv.ParseUint(h[i:i+1], 16, 8)
	//	//b := fmt.Sprintf("%04b", base)
	//	fmt.Println(hh[i])
	//	//copy(hash256[i*4:(i+1)*4], []byte(b))
	//	copy(hash256[i*4:(i+1)*4], hex[hh[i]])
	//}
	//fmt.Println(hash256,len(hash256))
	
	index := utils.Chaos(h, slot, limit)
	//index := 1
	return index
}

func main() {
	id := "9700849349fecf746aece91d903fc3b5ddf118dd806489cd8e998bdf3a3f5e73"
	var slot int64 = 3938238
	a:=[]byte{0,1}
	fmt.Println(a,int(a[0]))
	
	res := map[int]int{}
	
	for i := 0; i < 1; i++ {
		index := getDelegateIndex(id, slot+int64(i), 101)
		//fmt.Println(index)
		
		res[index] += 1
		
	}
	fmt.Println(res)
}
