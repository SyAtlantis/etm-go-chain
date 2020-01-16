package models

import (
	"bytes"
	"errors"
)

type TrDelegate struct {
	Username string
}

func init() {
	tr := TrDelegate{}
	RegisterSubTr(DELEGATE, &tr)
}

func (delegate *TrDelegate) create(tr *Transaction, data TrData) error {
	tr.Recipient = "A4MFB3MaPd355ug19GYPMSakCAWKbLjDTb"
	tr.Amount = 1000 * 100000000
	//tr.Asset.Delegate = TrDelegate{
	//	Username:  strings.ToLower(data.Username),
	//	PublicKey: data.Sender.PublicKey,
	//}
	return nil
}

func (delegate *TrDelegate) getBytes(tr *Transaction) ([]byte, error) {
	bb := bytes.NewBuffer([]byte{})
	//bb.WriteString(tr.Asset.Delegate.Username)

	return bb.Bytes(), nil
}

func (delegate *TrDelegate) verify(tr *Transaction) (bool error) {
	panic("implement me")
}

func (delegate *TrDelegate) process(tr *Transaction) error {
	panic("implement me")
}

func (delegate *TrDelegate) apply(tr *Transaction) error {
	sender := tr.SAccount
	if sender.IsEmpty() {
		return errors.New("sender account is empty")
	}
	sender.Delegate = &Delegate{
		Username:      tr.Args,
		TransactionId: tr.Id,
		Account:       sender,
	}

	return nil
}

func (delegate *TrDelegate) undo(tr *Transaction) error {
	panic("implement me")
}

func (delegate *TrDelegate) applyUnconfirmed(tr *Transaction) error {
	panic("implement me")
}

func (delegate *TrDelegate) undoUnconfirmed(tr *Transaction) error {
	panic("implement me")
}
