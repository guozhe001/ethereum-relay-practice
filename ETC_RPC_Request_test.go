package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/guozhe001/ethereum-relay-practice/constant"
	"github.com/guozhe001/ethereum-relay-practice/model"
	"github.com/guozhe001/ethereum-relay-practice/util"
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"strings"
	"testing"
)

func TestGetTransactionByHash(t *testing.T) {
	request := NewETCRPCRequest(constant.MyMainNetNodeUrl)
	transaction, err := request.GetTransactionByHash("0x302bf290c08b398c9c8098adf814f0a9fed3be1e9853f4bd5786eabd9ad31cdb")
	assert.NoError(t, err)
	assert.Equal(t, "0x1470d74659ea7ebc2083afe8c06dcb8bed7cd787f04b809f79e190212f033db8", transaction.BlockHash)
	log.Printf("%#v", transaction)
}

func TestGetTransactions(t *testing.T) {
	m := make(map[string]string)
	m["0x302bf290c08b398c9c8098adf814f0a9fed3be1e9853f4bd5786eabd9ad31cdb"] = "0x1470d74659ea7ebc2083afe8c06dcb8bed7cd787f04b809f79e190212f033db8"
	m["0x67c549e4f4f4146ecdc786aed7dfa0e69578739a7da9a322756a68970cea2701"] = "0xe273a05464bef5e5079d7ec4596f69cdba000df4f1b53c6f9d370074c607e6d3"
	m["0x67c549e4f4f4146ecdc786aed7dfa0e69578739a7da9a322756a689notexits0"] = ""
	request := NewETCRPCRequest(constant.MyMainNetNodeUrl)
	hashs := []string{}
	for k, _ := range m {
		hashs = append(hashs, k)
	}
	transactions, err := request.GetTransactions(hashs)
	assert.Equal(t, len(m), len(transactions))
	assert.NoError(t, err)
	for _, transaction := range transactions {
		log.Printf("%#v", transaction)
		assert.Equal(t, m[transaction.Hash], transaction.BlockHash)
	}
}

func TestGetBalance(t *testing.T) {
	request := NewETCRPCRequest(constant.MyMainNetNodeUrl)
	balance, err := request.GetBalance("0x7496F2A6ee0890D1eC6c316245BF92308699Aee4")
	assert.NoError(t, err)
	log.Println(balance, "wei")
}

func TestGetERC20BalanceLocal(t *testing.T) {
	request := NewETCRPCRequest(constant.MyLocalNetNodeUrl)
	balanceRequest := model.ERC20BalanceRequest{
		ContractAddress: "0x8D77117F45044A57b0564C5E3AF1d45F7442E581",
		UserAddress:     "0xc0F680767D4Ae17C7adaF8C6d0b4805Bc207805e",
		ContractDecimal: 21,
	}
	invokeGetERC20Balance(t, request, []model.ERC20BalanceRequest{balanceRequest})
}

func TestGetERC20BalanceRosten(t *testing.T) {
	balanceRequest := model.ERC20BalanceRequest{
		ContractAddress: "0xB8DfEe0D9aC703E75EE3D031148227B3BbB26524",
		UserAddress:     "0x2560be5793F9AA00963e163A1287807Feb897e2F",
		ContractDecimal: 21,
	}
	invokeGetERC20Balance(t, NewETCRPCRequest(constant.MyRopstenNetNodeUrl), []model.ERC20BalanceRequest{balanceRequest})
}

func invokeGetERC20Balance(t *testing.T, request *ETHRPCRequest, requests []model.ERC20BalanceRequest) {
	decimalMap := make(map[string]int64)
	for _, req := range requests {
		decimalMap[getDecimalMapKey(req.ContractAddress, req.UserAddress)] = req.ContractDecimal
	}
	responses, err := request.GetERC20Balances(requests)
	if err != nil {
		panic(err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, responses)
	for _, r := range responses {
		log.Printf("%#v, %s", r, *r.Balance)
	}
	for _, r := range responses {
		// 必须使用big.Int接收，因为solidity中有uint256，但是go语言最大只能表示uint64
		balance, b := util.HexToBigInt(*r.Balance)
		assert.True(t, b)
		log.Printf("ERC20Token-contractAddress=%s, userAddress=%s, balance=%d \n", r.ContractAddress, r.UserAddress, balance)
	}
}

// 获取key
func getDecimalMapKey(contractAddress, userAddress string) string {
	return strings.Join([]string{contractAddress, userAddress}, ":")
}

func TestNumber(t *testing.T) {
	// 这样除是不行的，因为如果除数小于被除数的时候得到的永远是0，这里要特别注意
	i := 21000000000000000000000000000 / math.Pow(10, 21)
	log.Println(i)
}

func TestHexUtil(t *testing.T) {
	decode, err := hexutil.Decode("0x4a817c800")
	assert.NoError(t, err)
	fmt.Print(decode)
}

func TestResultLength(t *testing.T) {
	log.Printf("len(0x000000000000000000000000000000000000000043dacaf91c1a84ff08000000)=%d\n", len("0x000000000000000000000000000000000000000043dacaf91c1a84ff08000000"))
}

func TestGetBlockNumber(t *testing.T) {
	request := NewETCRPCRequest(constant.MyRopstenNetNodeUrl)
	number, err := request.GetBlockNumber()
	assert.NoError(t, err)
	log.Printf("current block number is 十六进制：%#v 十进制：%s \n", number, number.String())
}

func TestGetBlockByNumber(t *testing.T) {
	invokeGetBlockByNumber(t, true)
}

func TestGetBlockByNumberFalse(t *testing.T) {
	invokeGetBlockByNumber(t, false)
}

func invokeGetBlockByNumber(t *testing.T, haveTransaction bool) {
	request := NewETCRPCRequest(constant.MyRopstenNetNodeUrl)
	hexNumber, err := request.GetBlockNumberHex()
	assert.NoError(t, err)
	block, err := request.GetBlockByHexNumber(hexNumber, haveTransaction)
	assert.NoError(t, err)
	log.Printf("block=%#v", block)
	jsonString, err := json.Marshal(block)
	assert.NoError(t, err)
	log.Printf("block json=%s", jsonString)
}
