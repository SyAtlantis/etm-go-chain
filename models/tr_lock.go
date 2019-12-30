package models

import "encoding/json"

type TrLock struct {
}

func init() {
	tr := TrLock{}
	RegisterTrs(LOCK, &tr)
}

func (lock *TrLock) create(tr *Transaction, data TrData) error {
	args, err := json.Marshal(data.Args)
	tr.Args = string(args)
	return err
}

func (lock *TrLock) getBytes(tr *Transaction) ([]byte, error) {
	return nil, nil
}

func (lock *TrLock) verify(tr *Transaction) (bool error) {
	panic("implement me")
}

func (lock *TrLock) process(tr *Transaction) error {
	panic("implement me")
}

func (lock *TrLock) apply(tr *Transaction) error {
	panic("implement me")
}

func (lock *TrLock) undo(tr *Transaction) error {
	panic("implement me")
}

func (lock *TrLock) applyUnconfirmed(tr *Transaction) error {
	panic("implement me")
}

func (lock *TrLock) undoUnconfirmed(tr *Transaction) error {
	panic("implement me")
}
