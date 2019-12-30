package models

import "encoding/json"

type TrUnlock struct {
}

func init() {
	tr := TrUnlock{}
	RegisterTrs(UNLOCK, &tr)
}

func (unlock *TrUnlock) create(tr *Transaction, data TrData) error {
	args, err := json.Marshal(data.Args)
	tr.Args = string(args)
	return err
}

func (unlock *TrUnlock) getBytes(tr *Transaction) ([]byte, error) {
	return nil, nil
}

func (unlock *TrUnlock) verify(tr *Transaction) (bool error) {
	panic("implement me")
}

func (unlock *TrUnlock) process(tr *Transaction) error {
	panic("implement me")
}

func (unlock *TrUnlock) apply(tr *Transaction) error {
	panic("implement me")
}

func (unlock *TrUnlock) undo(tr *Transaction) error {
	panic("implement me")
}

func (unlock *TrUnlock) applyUnconfirmed(tr *Transaction) error {
	panic("implement me")
}

func (unlock *TrUnlock) undoUnconfirmed(tr *Transaction) error {
	panic("implement me")
}
