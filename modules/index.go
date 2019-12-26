package modules

var (
	accounts     Accounts
	blocks       Blocks
	peers        Peers
	systems      Systems
	transactions Transactions
)

func InitModules() {
	accounts = NewAccounts()
	blocks = NewBlocks()
	peers = NewPeers()
	systems = NewSystems()
	transactions = NewTransactions()
}
