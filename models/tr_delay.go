package models

import "encoding/json"

type TrDelay struct {
}

func init() {
	tr := TrDelay{}
	RegisterTrs(DELAY, &tr)
}

func (delay *TrDelay) create(tr *Transaction, data TrData) error {
	tr.Recipient = &Account{
		Address: data.RecipientId,
	}
	tr.Amount = data.Amount
	args, err := json.Marshal(data.Args)
	tr.Args = string(args)
	return err
}

func (delay *TrDelay) getBytes(tr *Transaction) ([]byte, error) {
	return nil, nil
}
