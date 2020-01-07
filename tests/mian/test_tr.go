package mian

//
//func generateTr(secret string, data base.UserData) base.Transaction {
//	hash := sha256.Sum256([]byte(secret))
//	sender := base.Account{}
//	ed := utils.Ed{}
//	keypair := ed.MakeKeypair(hash[:])
//	sender.PublicKey = fmt.Sprintf("%x", keypair.PublicKey)
//	//sender.PublicKey = string(keypair.PublicKey[:])
//	//sender.PublicKey = hex.Dump(keypair.PublicKey)
//
//	data.Keypair = keypair
//	data.Sender = sender
//
//	tr := base.Transaction{}
//	tr.Create(data)
//
//	return tr
//}
//
//func trTransfer() base.Transaction {
//	data := base.UserData{
//		Type:        0,
//		Amount:      123456789,
//		Fee:         10000000,
//		Timestamp:   18972768,
//		RecipientId: "A79wqbYgZC5Bb923wWDXKjD7KBDn5BD6gg",
//	}
//	tr := generateTr("race forget pause shoe trick first abuse insane hope budget river enough", data)
//	return tr
//}
//
//func trSecond() base.Transaction {
//	hash2 := sha256.Sum256([]byte("asd"))
//	ed := utils.Ed{}
//	secondKeypair := ed.MakeKeypair(hash2[:])
//	data := base.UserData{
//		Type:          1,
//		Fee:           500000000,
//		Timestamp:     18972768,
//		SecondKeypair: secondKeypair,
//	}
//	tr := generateTr("immense buffalo organ pond illegal erupt prepare arrow cliff fit abstract task", data)
//	return tr
//}
//
//func trDelegate() base.Transaction {
//	data := base.UserData{
//		Type:      2,
//		Fee:       10000000,
//		Timestamp: 18972768,
//		Username:  "aaa",
//	}
//	tr := generateTr("worry net spend unfold desert trust dove waste grain people swap twelve", data)
//	tr.Asset.Delegate.PublicKey = tr.SenderPublicKey
//	return tr
//}
//
//func trUndelegate() base.Transaction {
//	data := base.UserData{
//		Type:      120,
//		Fee:       10000000,
//		Timestamp: 18972768,
//	}
//	tr := generateTr("worry net spend unfold desert trust dove waste grain people swap twelve", data)
//	return tr
//}
//
//func trVote() base.Transaction {
//	data := base.UserData{
//		Type:      3,
//		Fee:       10000000,
//		Timestamp: 18972768,
//		Votes:     []string{"+c6b1f18afa85a21df50cf9580c63c0aca4643a4a4e4ec93c2e397c81e87879b9"},
//	}
//	tr := generateTr("worry net spend unfold desert trust dove waste grain people swap twelve", data)
//	return tr
//}
//
//func trLock() base.Transaction {
//	data := base.UserData{
//		Type:      101,
//		Fee:       10000000,
//		Timestamp: 18972768,
//		Args:      []string{"100000000"},
//	}
//	tr := generateTr("worry net spend unfold desert trust dove waste grain people swap twelve", data)
//	return tr
//}
//
//func trUnlock() base.Transaction {
//	data := base.UserData{
//		Type:      102,
//		Fee:       10000000,
//		Timestamp: 18972768,
//		Args:      []string{"6ceed269a8b94598db804c87ab39a8f3cf41b31b4f54a2b3267fd49f362a71ff"},
//	}
//	tr := generateTr("worry net spend unfold desert trust dove waste grain people swap twelve", data)
//	return tr
//}
//
//func trDelay() base.Transaction {
//	data := base.UserData{
//		Type:        110,
//		Amount:      123456789,
//		Fee:         10000000,
//		Timestamp:   18972768,
//		RecipientId: "A79wqbYgZC5Bb923wWDXKjD7KBDn5BD6gg",
//		Args:        []string{"123"},
//	}
//	tr := generateTr("race forget pause shoe trick first abuse insane hope budget river enough", data)
//	return tr
//}
//
//// uia
//
//func trUiaIssuer() base.Transaction {
//	data := base.UserData{
//		Type:      9,
//		Fee:       10000000000,
//		Timestamp: 18972768,
//		Name:      "QQQ",
//		Desc:      "QQQ desc",
//	}
//	tr := generateTr("real rally sketch sorry place parrot typical cart stone mystery age nominee", data)
//	return tr
//}
//
//func trUiaAsset() base.Transaction {
//	data := base.UserData{
//		Type:           10,
//		Fee:            50000000000,
//		Timestamp:      18972768,
//		Name:           "QQQ.WWW",
//		Desc:           "QQQ desc",
//		Maximun:        "1000000000000",
//		Precision:      4,
//		Strategy:       "",
//		AllawWriteOff:  0,
//		AllowWhiteList: 0,
//		AllowBlackList: 0,
//	}
//	tr := generateTr("real rally sketch sorry place parrot typical cart stone mystery age nominee", data)
//	return tr
//}
//
//func trUiaIssue() base.Transaction {
//	data := base.UserData{
//		Type:      13,
//		Fee:       10000000,
//		Timestamp: 18972768,
//		Currency:  "QQQ.WWW",
//		UiaAmount: "1000000000000",
//	}
//	tr := generateTr("real rally sketch sorry place parrot typical cart stone mystery age nominee", data)
//	return tr
//}
//
//func trUiaTransfer() base.Transaction {
//	data := base.UserData{
//		Type:        14,
//		Fee:         10000000,
//		Timestamp:   18972768,
//		RecipientId: "A79wqbYgZC5Bb923wWDXKjD7KBDn5BD6gg",
//		Currency:    "QQQ.WWW",
//		UiaAmount:   "100000000",
//	}
//	tr := generateTr("real rally sketch sorry place parrot typical cart stone mystery age nominee", data)
//	return tr
//}
//
//func trUiaFlags() base.Transaction {
//	data := base.UserData{
//		Type:      11,
//		Fee:       10000000,
//		Timestamp: 18972768,
//		Currency:  "QQQ.WWW",
//		FlagType:  1,
//		Flag:      2,
//	}
//	tr := generateTr("real rally sketch sorry place parrot typical cart stone mystery age nominee", data)
//	return tr
//}
//
//func trUiaAcl() base.Transaction {
//	data := base.UserData{
//		Type:      12,
//		Fee:       10000000,
//		Timestamp: 18972768,
//		Currency:  "QQQ.WWW",
//		Operator:  "aaaa",
//		Flag:      1,
//		List:      []string{"a", "b"},
//	}
//	tr := generateTr("real rally sketch sorry place parrot typical cart stone mystery age nominee", data)
//	return tr
//}

func main() {
	//tr := trTransfer()
	//tr := trSecond()
	//tr := trDelegate()
	//tr := trUndelegate()
	//tr := trVote()
	//tr := trLock()
	//tr := trUnlock()
	//tr := trDelay()
	
	//tr := trUiaIssuer()
	//tr := trUiaAsset()
	//tr := trUiaIssue()
	//tr := trUiaTransfer()
	//tr := trUiaFlags()
	//tr := trUiaAcl()
	//
	//trByte,_ := json.Marshal(&tr)
	//
	//fmt.Println("tr=", string(trByte))
}
