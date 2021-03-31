package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/guozhe001/ethereum-relay-practice/constant"
	"github.com/guozhe001/ethereum-relay-practice/model"
	"github.com/guozhe001/ethereum-relay-practice/util"
	"log"
	"math/big"
)

type ETHRPCRequest struct {
	client *ETHRPCClient
}

func NewETCRPCRequest(nodeUrl string) *ETHRPCRequest {
	return &ETHRPCRequest{client: NewETHRPCClient(nodeUrl)}
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
	result := []*model.Transaction{}
	request := []rpc.BatchElem{}
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
	decode, _ := hexutil.Decode(gasPrice)
	log.Printf("now gas price is %s\n", decode)
	return gasPrice, err
}

// GetERC20Balances 获取满足ERC20定义的合约的余额，使用eth_call方式查询
func (e *ETHRPCRequest) GetERC20Balances(params []model.ERC20BalanceRequest) ([]model.ERC20BalanceResponse, error) {
	methodId := "0x70a08231"
	batchElems := []rpc.BatchElem{}
	result := []model.ERC20BalanceResponse{}
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

// InvokeERC20Transfer 调用ERC20的transfer方法
func (e *ETHRPCRequest) InvokeERC20Transfer(to string, value big.Int) error {
	var success bool
	methodId, err := util.GetMethodId("./abi/GZToken_metadata.json", "transfer")
	if err != nil {
		return err
	}
	err = e.client.client.Call(&success, constant.MethodEthCall, methodId)
	if err != nil {
		return err
	}
	return nil
}
