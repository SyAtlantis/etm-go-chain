package models

type System struct {
	Version      string
	LastHeight   int64
	LastBlock    *Block
	MyDelegates  []string
	DelegateList []string
}
