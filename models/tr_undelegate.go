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

func (undelegate *TrUndelegate) verify(tr *Transaction) (bool error) {
	panic("implement me")
}

func (undelegate *TrUndelegate) process(tr *Transaction) error {
	panic("implement me")
}

func (undelegate *TrUndelegate) apply(tr *Transaction) error {
	panic("implement me")
}

func (undelegate *TrUndelegate) undo(tr *Transaction) error {
	panic("implement me")
}

func (undelegate *TrUndelegate) applyUnconfirmed(tr *Transaction) error {
	panic("implement me")
}

func (undelegate *TrUndelegate) undoUnconfirmed(tr *Transaction) error {
	panic("implement me")
}
