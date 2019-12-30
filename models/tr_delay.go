package models

import "encoding/json"

type TrDelay struct {
}

func init() {
	tr := TrDelay{}
	RegisterTrs(DELAY, &tr)
}

func (delay *TrDelay) create(tr *Transaction, data TrData) error {
	tr.Recipient = data.RecipientId
	tr.Amount = data.Amount
	args, err := json.Marshal(data.Args)
	tr.Args = string(args)
	return err
}

func (delay *TrDelay) getBytes(tr *Transaction) ([]byte, error) {
	return nil, nil
}

func (delay *TrDelay) verify(tr *Transaction) (bool error) {
	panic("implement me")
}

func (delay *TrDelay) process(tr *Transaction) error {
	panic("implement me")
}

func (delay *TrDelay) apply(tr *Transaction) error {
	panic("implement me")
}

func (delay *TrDelay) undo(tr *Transaction) error {
	panic("implement me")
}

func (delay *TrDelay) applyUnconfirmed(tr *Transaction) error {
	panic("implement me")
}

func (delay *TrDelay) undoUnconfirmed(tr *Transaction) error {
	panic("implement me")
}
