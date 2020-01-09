package models

import "errors"

type TrTransfer struct {
}

func init() {
	tr := TrTransfer{}
	RegisterTrs(TRANSFER, &tr)
}

func (transfer *TrTransfer) create(tr *Transaction, data TrData) error {
	tr.Recipient = data.RecipientId
	tr.Amount = data.Amount
	return nil
}

func (transfer *TrTransfer) getBytes(tr *Transaction) ([]byte, error) {
	return nil, nil
}

func (transfer *TrTransfer) verify(tr *Transaction) (bool error) {
	panic("implement me")
}

func (transfer *TrTransfer) process(tr *Transaction) error {
	panic("implement me")
}

func (transfer *TrTransfer) apply(tr *Transaction) error {
	recipient := tr.RAccount
	if recipient.IsEmpty() {
		return errors.New("recipient account is empty")
	}

	amount := tr.Amount + tr.Fee
	recipient.Balance += amount

	return nil
}

func (transfer *TrTransfer) undo(tr *Transaction) error {
	panic("implement me")
}

func (transfer *TrTransfer) applyUnconfirmed(tr *Transaction) error {
	panic("implement me")
}

func (transfer *TrTransfer) undoUnconfirmed(tr *Transaction) error {
	panic("implement me")
}
