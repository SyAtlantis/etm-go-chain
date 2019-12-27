package models

type TrUndelegate struct {
}

func init() {
	tr := TrUndelegate{}
	RegisterTrs(UNDELEGATE, &tr)
}

func (undelegate *TrUndelegate) create(tr *Transaction, data TrData) error {
	tr.Recipient = ""
	tr.Amount = 0
	return nil
}

func (undelegate *TrUndelegate) getBytes(tr *Transaction) ([]byte, error) {
	return nil, nil
}
