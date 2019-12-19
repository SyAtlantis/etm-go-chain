package modules

var (
	accounts     Accounts
	blocks       Blocks
	peers        Peers
	system       System
	transactions Transactions
)

func InitModules() {
	accounts = NewAccounts()
	blocks = NewBlocks()
	peers = NewPeers()
	system = NewSystem()
	transactions = NewTransactions()
}
