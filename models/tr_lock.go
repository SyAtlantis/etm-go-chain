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
