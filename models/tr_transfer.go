package models

type TrTransfer struct {
}

func init() {
	tr := TrTransfer{}
	RegisterTrs(TRANSFER, &tr)
}

func (transfer *TrTransfer) create(tr *Transaction, data TrData) error {
	tr.Recipient = &Account{
		Address: data.RecipientId,
	}
	tr.Amount = data.Amount
	return nil
}

func (transfer *TrTransfer) getBytes(tr *Transaction) ([]byte, error) {
	return nil, nil
}
