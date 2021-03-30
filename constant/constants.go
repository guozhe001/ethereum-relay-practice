package constant

// MyMainNetNodeUrl 我们要请求的node的url
const (
	MyMainNetNodeUrl    string = "https://mainnet.infura.io/v3/94bc20a138044cd7974fcd20b91d68ba"
	MyLocalNetNodeUrl          = "http://localhost:8545"
	MyRopstenNetNodeUrl        = "https://ropsten.infura.io/v3/94bc20a138044cd7974fcd20b91d68ba"
)

// 一些api方法
const (
	MethodGetTransactionByHash string = "eth_getTransactionByHash"
	MethodGetBalance                  = "eth_getBalance"
	MethodEthCall                     = "eth_call"     // Executes a new message call immediately without creating a transaction on the block chain.
	MethodGasPrice                    = "eth_gasPrice" // Returns the current price per gas in wei.
)

// QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter
const (
	TagLatest   string = "latest"
	TagEarliest        = "earliest"
	TagPending         = "pending"
)

// 64个0
const ZeroOf64 = "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"