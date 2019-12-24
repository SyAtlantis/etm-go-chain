package models

type Transfer struct {
	
}

func (Transfer) create(tr *Transaction, data TrData) {
	panic("implement me")
}

func (Transfer) getBytes(tr *Transaction) []byte {
	panic("implement me")
}
