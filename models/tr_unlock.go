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
