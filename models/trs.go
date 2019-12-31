package models

import "sort"

type Trs []*Transaction

type iTransactions interface {
	Sort()
}

func (trs Trs) Len() int {
	return len(trs)
}

func (trs Trs) Swap(i, j int) {
	trs[i], trs[j] = trs[j], trs[i]
}

func (trs Trs) Less(i, j int) bool {
	if trs[i].Type != trs[j].Type {
		if trs[i].Type == 1 {
			return false
		}
		if trs[j].Type == 1 {
			return true
		}
		return trs[i].Type < trs[j].Type
	}
	if trs[i].Amount != trs[j].Amount {
		return trs[i].Amount < trs[j].Amount
	}
	return trs[i].Id < trs[j].Id
}

func (trs Trs) Sort() {
	sort.Sort(trs)
}
