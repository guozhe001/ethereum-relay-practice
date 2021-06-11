package constant

import "math/big"

// MyMainNetNodeUrl 我们要请求的node的url
const (
  MyMainNetNodeUrl    string = "https://mainnet.infura.io/v3/94bc20a138044cd7974fcd20b91d68ba"
  MyLocalNetNodeUrl          = "http://localhost:8545"
  MyRopstenNetNodeUrl        = "https://ropsten.infura.io/v3/94bc20a138044cd7974fcd20b91d68ba"
  //MyBscMainNetUrl            = "https://bsc-dataseed1.binance.org"
  MyBscMainNetUrl = "https://apis.ankr.com/40d05760bba540ab95cb87bc5b449ab8/602045ef2a200f867afc105676fe0511/binance/full/main"
)

type ChainInfo struct {
  NodeUrl string   // 节点url
  ChainId *big.Int // 链ID
}

const (
  ChainNameEthereumMainNet    string = "EthereumMainNet"
  ChainNameEthereumRopstenNet        = "EthereumRopstenNet"
  ChainNameBSCMainNet                = "BSCMainNet"
  ChainNameLocalHost                 = "LocalHost"
)

var chains = make(map[string]ChainInfo)

func init() {
  chains[ChainNameEthereumMainNet] = ChainInfo{NodeUrl: MyMainNetNodeUrl, ChainId: big.NewInt(1)}
  chains[ChainNameEthereumRopstenNet] = ChainInfo{NodeUrl: MyRopstenNetNodeUrl, ChainId: big.NewInt(3)}
  chains[ChainNameLocalHost] = ChainInfo{NodeUrl: MyLocalNetNodeUrl, ChainId: big.NewInt(1212112212)}
  chains[ChainNameBSCMainNet] = ChainInfo{NodeUrl: MyBscMainNetUrl, ChainId: big.NewInt(56)}
}

func GetChainInfo(chainName string) ChainInfo {
  return chains[chainName]
}

// 一些api方法
const (
  MethodGetTransactionByHash  string = "eth_getTransactionByHash"
  MethodGetBalance                   = "eth_getBalance"
  MethodEthCall                      = "eth_call"               // Executes a new message call immediately without creating a transaction on the block chain.
  MethodGasPrice                     = "eth_gasPrice"           // Returns the current price per gas in wei.
  MethodBlockNumber                  = "eth_blockNumber"        // Returns the number of most recent block.
  MethodGetBlockByNumber             = "eth_getBlockByNumber"   // Returns information about a block by block number.
  MethodBlockByHash                  = "eth_getBlockByHash"     // Returns information about a block by hash.
  MethodEthSendRawTransaction        = "eth_sendRawTransaction" // Creates new message call transaction or a contract creation for signed transactions.
)

// QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter
const (
  TagLatest   string = "latest"
  TagEarliest        = "earliest"
  TagPending         = "pending"
)

// ZeroOf64 64个0
const ZeroOf64 = "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

// HexPrefix 十六进制字符串前缀
const HexPrefix = "0x"

const (
  ContractERC20BUSD = "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"
)

const (
  DefaultDecimal = 18
)
