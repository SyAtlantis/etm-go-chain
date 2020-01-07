package models

import (
	"encoding/json"
	"errors"
	"strconv"
)

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
	sender := tr.SAccount
	if sender.IsEmpty() {
		return errors.New("sender account is empty")
	}
	lockAmount, err := strconv.ParseInt(tr.Args, 10, 64)
	if err != nil {
		return errors.New("lock amount is not the type of int64")
	}
	if sender.Balance < lockAmount {
		return errors.New("not enough balance to lock")
	}
	l := &Lock{
		LockAmount:    lockAmount,
		TransactionId: tr.Id,
		Account:       &sender,
	}
	sender.Locks = append(sender.Locks, l)
	sender.Balance -= lockAmount
	if err = sender.Merge(); err != nil {
		return err
	}

	return nil
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
