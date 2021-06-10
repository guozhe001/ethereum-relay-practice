package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/guozhe001/ethereum-relay-practice/constant"
	"github.com/guozhe001/ethereum-relay-practice/model"
	"github.com/guozhe001/ethereum-relay-practice/util"
	"math/big"
)

type ETHRPCRequest struct {
	nonceManager *NonceManager // nonce 管理者实例
	client       *ETHRPCClient
}

func NewETCRPCRequest(nodeUrl string) *ETHRPCRequest {
	return &ETHRPCRequest{nonceManager: NewNonceManager(), client: NewETHRPCClient(nodeUrl)}
}

// GetTransactionByHash 根据hash获取交易
func (e *ETHRPCRequest) GetTransactionByHash(hash string) (model.Transaction, error) {
	transaction := model.Transaction{}
	// 在调用call的时候，要把transaction的指针传过去，不然接收不到
	err := e.client.client.Call(&transaction, constant.MethodGetTransactionByHash, hash)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

// GetTransactions 批量查询交易
func (e *ETHRPCRequest) GetTransactions(transactionHash []string) ([]*model.Transaction, error) {
	var result []*model.Transaction
	var request []rpc.BatchElem
	for _, hash := range transactionHash {
		transaction := model.Transaction{}
		request = append(request, rpc.BatchElem{Method: constant.MethodGetTransactionByHash, Args: []interface{}{hash}, Result: &transaction})
		// 这里在append时必须是指针，因为现在定义的是空的transaction而调用BatchCall时传入的时transaction的指针，如果此处不是设置的指针则返回的还是定义时的空transaction
		result = append(result, &transaction)
	}
	err := e.client.client.BatchCall(request)
	return result, err
}

// GetBalance 获取指定地址的余额
func (e *ETHRPCRequest) GetBalance(address string) (string, error) {
	var balance string
	err := e.client.client.Call(&balance, constant.MethodGetBalance, address, constant.TagLatest)
	if err != nil {
		return "", err
	}
	return util.WeiToEth(balance)
}

// GetBalances 批量获取余额,暂时就不实现了，跟批量查询交易的逻辑一样
func (e *ETHRPCRequest) GetBalances(addresses []string) ([]*string, error) {
	return nil, nil
}

// GetGasPrice 获取当前的gas价格
func (e *ETHRPCRequest) GetGasPrice() (string, error) {
	var gasPrice string
	err := e.client.client.Call(&gasPrice, constant.MethodGasPrice)
	return gasPrice, err
}

// GetERC20Balances 获取满足ERC20定义的合约的余额，使用eth_call方式查询
func (e *ETHRPCRequest) GetERC20Balances(params []model.ERC20BalanceRequest) ([]model.ERC20BalanceResponse, error) {
	methodId := "0x70a08231"
	var batchElems []rpc.BatchElem
	var result []model.ERC20BalanceResponse
	gasPrice, err := e.GetGasPrice()
	if err != nil {
		return nil, err
	}
	for _, req := range params {
		var balance string
		zero, _ := util.GetZero(24)
		arg := &model.EthCallRequest{
			To:       common.HexToAddress(req.ContractAddress),
			Data:     methodId + zero + req.UserAddress[2:],
			Gas:      hexutil.EncodeUint64(41420),
			GasPrice: gasPrice,
		}
		batchElems = append(batchElems, rpc.BatchElem{Method: constant.MethodEthCall, Result: &balance, Args: []interface{}{arg, constant.TagLatest}})
		result = append(result, model.ERC20BalanceResponse{ContractAddress: req.ContractAddress, UserAddress: req.UserAddress, Balance: &balance})
	}
	if err := e.client.client.BatchCall(batchElems); err != nil {
		return nil, err
	}
	// 遍历所有的结果检查是否有报错的
	for _, elem := range batchElems {
		if elem.Error != nil {
			return nil, elem.Error
		}
	}
	return result, nil
}

// GetBlockNumber 获取当前网络的最新的高度
func (e *ETHRPCRequest) GetBlockNumber() (big.Int, error) {
	blockNumber, err := e.GetBlockNumberHex()
	if err != nil {
		return big.Int{}, err
	}
	setString, _ := new(big.Int).SetString(blockNumber[2:], 16)
	return *setString, err
}

// GetBlockNumber 获取当前网络的最新的高度
func (e *ETHRPCRequest) GetBlockNumberHex() (string, error) {
	blockNumber := ""
	err := e.client.client.Call(&blockNumber, constant.MethodBlockNumber)
	return blockNumber, err
}

// GetBlockByNumber 根据区块高度获取区块信息
func (e *ETHRPCRequest) GetBlockByNumber(blockNumber big.Int, haveTransaction bool) (model.Block, error) {
	return e.GetBlockByHexNumber(util.BigIntToHex(blockNumber), haveTransaction)
}

// GetBlockByNumber 根据区块高度获取区块信息
func (e *ETHRPCRequest) GetBlockByHexNumber(blockHexNumber string, haveTransaction bool) (model.Block, error) {
	block := model.Block{}
	err := e.client.client.Call(&block, constant.MethodGetBlockByNumber, blockHexNumber, haveTransaction)
	return block, err
}

// GetBlockByHash 根据区块hash获取区块信息
func (e *ETHRPCRequest) GetBlockByHash(hash string, haveTransaction bool) (model.Block, error) {
	block := model.Block{}
	err := e.client.client.Call(&block, constant.MethodBlockByHash, hash, haveTransaction)
	return block, err
}

// EthCall 调用eth_call
func (e *ETHRPCRequest) EthCall(result interface{}, request model.EthCallRequest) error {
	return e.client.client.Call(&result, constant.MethodEthCall, request, constant.TagLatest)
}

// EthSendRawTransaction 调用eth_sendRawTransaction方法
// Creates new message call transaction or a contract creation for signed transactions.
func (e *ETHRPCRequest) EthSendRawTransaction(data string) (txHash string, err error) {
	if err = e.client.client.Call(&txHash, constant.MethodEthSendRawTransaction, data); err != nil {
		return txHash, err
	}
	return txHash, nil
}

// GetNonce 获取地址的 nonce 值
func (e *ETHRPCRequest) GetNonce(address string) (uint64, error) {
	methodName := "eth_getTransactionCount"
	nonce := ""
	// 因为我们要查询最新的，根据基于 etTransactionCount 情况下的区块号关系，选取 pending
	err := e.client.client.Call(&nonce, methodName, address, "pending")
	if err != nil {
		return 0, fmt.Errorf("发送交易失败! %s", err.Error())
	}
	n, _ := new(big.Int).SetString(nonce[2:], 16) // 采用大数类型将 16 进制的结果转为 10 进制
	return n.Uint64(), nil                        // 返回交易 hash
}

//
//// 发送 ERC20 代币交易，或称转账 ERC20 代币
//// 参数分别是：
//// 	   交易发起地址，代币合约地址，交易接收地址，代币数量，油费设置，代币的 decimal 值
//func (e *ETHRPCRequest) invokeSoloTop(fromStr, contact, valueStr, methodId string, gasLimit, gasPrice uint64, decimal int) (string, error) {
//
//  to := common.HexToAddress(contact) // 将合约 contact 字符串类型的转为 address 类型的
//  gasPrice_ := new(big.Int).SetUint64(gasPrice)
//
//  // 结构体中的 value 字段为 0
//  amount := new(big.Int).SetInt64(0)
//
//  // 获取 nonce
//  if nonce := e.nonceManager.GetNonce(fromStr); nonce == nil {
//    // nonce 不存在，开始访问节点获取
//    n, err := e.GetNonce(fromStr)
//    if err != nil {
//      return "", fmt.Errorf("获取 nonce 失败 %s", err.Error())
//    }
//    nonce = new(big.Int).SetUint64(n)
//  }
//
//  // 构建 data，真实的 value 转账数值由 data 携带
//  data := BuildERC20TransferData(methodId, valueStr, receiver, decimal)
//  dataBytes := []byte(data)
//
//  // 构建交易结构体
//  transaction := types.NewTransaction(
//    nonce.Uint64(),
//    to,
//    amount,
//    gasLimit,
//    gasPrice_,
//    dataBytes)
//
//  return e.SendTransaction(fromStr, transaction)
//}
//
//// 发送交易，根据入参 transaction 的不同变量设置，达到发送不同种类的交易
//func (e *ETHRPCRequest) SendTransaction(address string, transaction *types.Transaction) (string, error) {
//  // 签名交易数据
//  signTx, err := SignETHTransaction(address, transaction)
//  if err != nil {
//    return "", fmt.Errorf("签名失败! %s", err.Error())
//  }
//  // rlp 序列化
//  txRlpData, err := rlp.EncodeToBytes(signTx)
//  if nil != err {
//    return "", fmt.Errorf("rlp 序列化失败! %s", err.Error())
//  }
//  // 下面调用以太坊的 rpc 接口
//  txHash := ""
//  methodName := "eth_sendRawTransaction"
//  err = e.client.client.Call(&txHash, methodName, common.ToHex(txRlpData))
//  if err != nil {
//    return "", fmt.Errorf("发送交易失败! %s", err.Error())
//  }
//  return txHash, nil // 返回交易 hash
//}
