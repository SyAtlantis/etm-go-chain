package models

import (
	"bytes"
)

type TrUnvote struct {
}

func init() {
	tr := TrVote{}
	RegisterTrs(UNVOTE, &tr)
}

func (unvote *TrUnvote) create(tr *Transaction, data TrData) error {
	return nil
}

func (unvote *TrUnvote) getBytes(tr *Transaction) ([]byte, error) {
	bb := bytes.NewBuffer([]byte{})
	//for i := 0; i < len(tr.Asset.Vote.Votes); i++ {
	//	bb.WriteString(tr.Asset.Vote.Votes[i])
	//}

	return bb.Bytes(), nil
}

func (unvote *TrUnvote) verify(tr *Transaction) (bool error) {
	panic("implement me")
}

func (unvote *TrUnvote) process(tr *Transaction) error {
	panic("implement me")
}

func (unvote *TrUnvote) apply(tr *Transaction) error {

	return nil
}

func (unvote *TrUnvote) undo(tr *Transaction) error {
	panic("implement me")
}

func (unvote *TrUnvote) applyUnconfirmed(tr *Transaction) error {
	panic("implement me")
}

func (unvote *TrUnvote) undoUnconfirmed(tr *Transaction) error {
	panic("implement me")
}
