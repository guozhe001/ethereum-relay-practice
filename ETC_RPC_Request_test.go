package main

import (
  "encoding/json"
  "fmt"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/common/hexutil"
  "github.com/guozhe001/ethereum-relay-practice/constant"
  "github.com/guozhe001/ethereum-relay-practice/model"
  "github.com/guozhe001/ethereum-relay-practice/util"
  "github.com/guozhe001/ethereum-relay-practice/wallet"
  "github.com/stretchr/testify/assert"
  "log"
  "math"
  "math/big"
  "strings"
  "testing"
)

func TestGetTransactionByHash(t *testing.T) {
  request := NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameEthereumMainNet))
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
  request := NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameEthereumMainNet))
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
  request := NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameEthereumMainNet))
  balance, err := request.GetBalance("0x7496F2A6ee0890D1eC6c316245BF92308699Aee4")
  assert.NoError(t, err)
  log.Println(balance, "wei")
}

//func TestGetERC20BalanceLocal(t *testing.T) {
//  request := NewETCRPCRequest(constant.MyLocalNetNodeUrl)
//  balanceRequest := model.ERC20BalanceRequest{
//    ContractAddress: "0x8D77117F45044A57b0564C5E3AF1d45F7442E581",
//    UserAddress:     "0xc0F680767D4Ae17C7adaF8C6d0b4805Bc207805e",
//    ContractDecimal: 21,
//  }
//  invokeGetERC20Balance(t, request, []model.ERC20BalanceRequest{balanceRequest})
//}

func TestGetERC20BalanceRosten(t *testing.T) {
  balanceRequest := model.ERC20BalanceRequest{
    ContractAddress: "0xB8DfEe0D9aC703E75EE3D031148227B3BbB26524",
    UserAddress:     "0x2560be5793F9AA00963e163A1287807Feb897e2F",
    ContractDecimal: 21,
  }
  invokeGetERC20Balance(t, NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameEthereumRopstenNet)), []model.ERC20BalanceRequest{balanceRequest})
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
    // ????????????big.Int???????????????solidity??????uint256?????????go????????????????????????uint64
    balance, b := util.HexToBigInt(*r.Balance)
    assert.True(t, b)
    log.Printf("ERC20Token-contractAddress=%s, userAddress=%s, balance=%d \n", r.ContractAddress, r.UserAddress, balance)
  }
}

// ??????key
func getDecimalMapKey(contractAddress, userAddress string) string {
  return strings.Join([]string{contractAddress, userAddress}, ":")
}

func TestNumber(t *testing.T) {
  // ????????????????????????????????????????????????????????????????????????????????????0????????????????????????
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
  request := NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameEthereumRopstenNet))
  number, err := request.GetBlockNumber()
  assert.NoError(t, err)
  log.Printf("current block number is ???????????????%#v ????????????%s \n", number, number.String())
}

func TestGetBlockByNumber(t *testing.T) {
  invokeGetBlockByNumber(t, true)
}

func TestGetBlockByNumberFalse(t *testing.T) {
  invokeGetBlockByNumber(t, false)
}

func TestETHRPCRequest_GetBlockByHash(t *testing.T) {
  request := NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameEthereumRopstenNet))
  block := invokeGetBlockByNumber(t, false)
  blockByHash, err := request.GetBlockByHash(block.Hash, false)
  assert.NoError(t, err)
  assert.Equal(t, block.Number, blockByHash.Number)
}

func invokeGetBlockByNumber(t *testing.T, haveTransaction bool) model.Block {
  request := NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameEthereumRopstenNet))
  hexNumber, err := request.GetBlockNumberHex()
  assert.NoError(t, err)
  block, err := request.GetBlockByHexNumber(hexNumber, haveTransaction)
  assert.NoError(t, err)
  log.Printf("block=%#v", block)
  jsonString, err := json.Marshal(block)
  assert.NoError(t, err)
  log.Printf("block json=%s", jsonString)
  return block
}

func TestETHRPCRequest_InvokeERC20Transfer(t *testing.T) {
  transferMethodId, _ := util.GetMethodIdByFile("abi/GZToken_metadata.json", "transfer")
  request := NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameEthereumRopstenNet))
  price, _ := request.GetGasPrice()
  log.Print("price=", price)
  var success string
  data := model.EthCallRequest{
    To:       common.HexToAddress("0xB8DfEe0D9aC703E75EE3D031148227B3BbB26524"),
    Gas:      "0xde0b000",
    GasPrice: price,
    Data:     transferMethodId + common.HexToHash("0x15f94602F738f280d9A471B8fc34eDadeF6DD890").String()[2:] + common.HexToHash("0x2710").String()[2:],
  }
  err := request.EthCall(&success, data)
  if err != nil {
    log.Print(err)
  }
  assert.NoError(t, err)
  log.Print(success)
}

func TestGetNonce(t *testing.T) {
  request := NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameEthereumRopstenNet))
  nonce, err := request.GetNonce("0x2560be5793F9AA00963e163A1287807Feb897e2F")
  assert.NoError(t, err)
  log.Print("nonce=", nonce)
}

func TestBscGetTransactionByHash(t *testing.T) {
  request := NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameBSCMainNet))
  transaction, err := request.GetTransactionByHash("0x44558267d0794d978b72fc7e6a64ef68c1c8b2b4d7bbc2ac62c1907645d3929d")
  assert.NoError(t, err)
  assert.Equal(t, "0xdb4337fed84a47cf109b69884310d32c53eb3ff14d9268e54d3e10058062de5b", transaction.BlockHash)
  log.Printf("%#v", transaction)
}

// ??????ERC20token???balanceOf??????
func TestETHRPCRequest_InvokeERC20Balance(t *testing.T) {
  decimal := int64(math.Pow(10, 18))
  transferMethodId, _ := util.GetMethodIdByFile("abi/BSC_busd.json", "balanceOf")
  request := NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameBSCMainNet))
  price, _ := request.GetGasPrice()
  log.Print("price=", price)
  var balance string
  data := model.EthCallRequest{
    To:       common.HexToAddress(constant.ContractERC20BUSD),
    Gas:      "0xde0b000",
    GasPrice: price,
    Data:     transferMethodId + common.HexToHash("0xc4c1bb77a87405d7e58a15a9d83cd8454da4faa3").String()[2:],
  }

  err := request.EthCall(&balance, data)
  if err != nil {
    log.Print("err=", err)
  }
  assert.NoError(t, err)
  bigIntBalance, _ := util.HexToBigInt(balance)
  log.Print("balance ????????????=", bigIntBalance)

  b := bigIntBalance.Div(bigIntBalance, big.NewInt(decimal))
  log.Print("balance ????????????=", b)
}

func TestSendETHTransaction(t *testing.T) {
  err := wallet.UnlockETHWallet("0xcd97b8b5ef2f219bed97c9b5b8035454f3f9ed18", "this_is_for_test")
  assert.NoError(t, err)
  request := NewETCRPCRequest(constant.GetChainInfo(constant.ChainNameEthereumRopstenNet))
  txHash, err := request.SendETHTransaction("0xcd97b8b5ef2f219bed97c9b5b8035454f3f9ed18", "0xfc18d8882697552737921dF830102579f000008D",
    "0.0001", 21000, 5)
  assert.NoError(t, err)
  log.Println("txHash=", txHash)
}
