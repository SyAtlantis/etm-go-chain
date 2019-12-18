package transactions

type transfer struct {
}

func NewTransfer() Transaction {
	return &transfer{}
}

func (t transfer) Create() {
	panic("implement me")
}

func (t transfer) GetBytes() []byte {
	panic("implement me")
}

func (t transfer) GetHash() [32]byte {
	panic("implement me")
}

func (t transfer) GetId() string {
	panic("implement me")
}

func (t transfer) GetSignature() string {
	panic("implement me")
}
