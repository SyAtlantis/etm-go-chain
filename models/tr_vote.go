package models

import "bytes"

type TrVote struct {
	Votes []string
}

func init() {
	tr := TrVote{}
	RegisterTrs(VOTE, &tr)
}

func (vote *TrVote) create(tr *Transaction, data TrData) error {
	//tr.Asset.Vote.Votes = data.Votes
	return nil
}

func (vote *TrVote) getBytes(tr *Transaction) ([]byte, error) {
	bb := bytes.NewBuffer([]byte{})
	//for i := 0; i < len(tr.Asset.Vote.Votes); i++ {
	//	bb.WriteString(tr.Asset.Vote.Votes[i])
	//}

	return bb.Bytes(), nil
}

func (vote *TrVote) verify(tr *Transaction) (bool error) {
	panic("implement me")
}

func (vote *TrVote) process(tr *Transaction) error {
	panic("implement me")
}

func (vote *TrVote) apply(tr *Transaction) error {
	panic("implement me")
}

func (vote *TrVote) undo(tr *Transaction) error {
	panic("implement me")
}

func (vote *TrVote) applyUnconfirmed(tr *Transaction) error {
	panic("implement me")
}

func (vote *TrVote) undoUnconfirmed(tr *Transaction) error {
	panic("implement me")
}
