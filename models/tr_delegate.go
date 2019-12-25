package models

import (
	"bytes"
	"strings"
)

type TrDelegate struct {
	Username  string
	PublicKey string
}

func init() {
	tr := TrDelegate{}
	RegisterTrs(DELEGATE, &tr)
}

func (delegate *TrDelegate) create(tr *Transaction, data TrData) error {
	tr.Recipient = &Account{
		Address: "A4MFB3MaPd355ug19GYPMSakCAWKbLjDTb",
	}
	tr.Amount = 1000 * 100000000
	tr.Asset.Delegate = TrDelegate{
		Username:  strings.ToLower(data.Username),
		PublicKey: data.Sender.PublicKey,
	}
	return nil
}

func (delegate *TrDelegate) getBytes(tr *Transaction) ([]byte, error) {
	bb := bytes.NewBuffer([]byte{})
	bb.WriteString(tr.Asset.Delegate.Username)

	return bb.Bytes(), nil
}
