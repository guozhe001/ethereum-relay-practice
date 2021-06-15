package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/guozhe001/ethereum-relay-practice/contract/solo"
	"github.com/guozhe001/ethereum-relay-practice/contract/weth"
	"log"
	"math"
	"math/big"
	"testing"
)

const (
	ropstenUrl string = "https://ropsten.infura.io/v3/94bc20a138044cd7974fcd20b91d68ba"
	pk                = "replace me"
)

func TestSend(t *testing.T) {
	client, err := ethclient.Dial(ropstenUrl)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)             // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xfc18d8882697552737921dF830102579f000008D")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func TestTransfer(t *testing.T) {
	client, err := ethclient.Dial(ropstenUrl)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Printf("nonce=%d\n", nonce)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(3))
	if err != nil {
		log.Fatal(err)
	}
	//auth.Nonce = big.NewInt(nonce)
	auth.Nonce = big.NewInt(int64(4))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	//auth.GasPrice = gasPrice
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(10))

	address := common.HexToAddress("0xb8dfee0d9ac703e75ee3d031148227b3bbb26524")
	instance, err := gztoken.NewMain(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.Transfer(auth, common.HexToAddress("0xfc18d8882697552737921dF830102579f000008D"), big.NewInt(int64(math.Pow(10, 21))))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}

func TestInvokeWETHTransfer(t *testing.T) {
	client, err := ethclient.Dial(ropstenUrl)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Printf("nonce=%d\n", nonce)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(3))
	if err != nil {
		log.Fatal(err)
	}
	//auth.Nonce = big.NewInt(int64(nonce))
	auth.Nonce = big.NewInt(int64(7))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(10))
	//auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(10))

	address := common.HexToAddress("0xc778417e063141139fce010982780140aa0cd5ab")
	instance, err := weth.NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}
	wei := big.NewInt(int64(math.Pow(10, 18)))
	tx, err := instance.Transfer(auth, common.HexToAddress("0xfc18d8882697552737921dF830102579f000008D"), wei.Div(wei, big.NewInt(10000)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}

const MyBscMainNetUrl string = "https://apis.ankr.com/40d05760bba540ab95cb87bc5b449ab8/602045ef2a200f867afc105676fe0511/binance/full/main"

func TestInvokeSoloTopBorrow(t *testing.T) {
	client, err := ethclient.Dial(MyBscMainNetUrl)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Printf("nonce=%d\n", nonce)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(56))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	//auth.Nonce = big.NewInt(int64(7))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))

	//auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(10))

	address := common.HexToAddress("0x7033a512639119c759a51b250bfa461ae100894b")
	instance, err := solo.NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}
	wei := big.NewInt(int64(math.Pow(10, 18)))
	tx, err := instance.Borrow(auth, big.NewInt(6), wei.Mul(wei, big.NewInt(100000)))
	if err != nil {
		log.Fatal("invoke Borrow failed: ", err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}

func TestInvokeSoloTopDeposit(t *testing.T) {
	client, err := ethclient.Dial(MyBscMainNetUrl)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Printf("nonce=%d\n", nonce)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(56))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	//auth.Nonce = big.NewInt(int64(7))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))
	//auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(10))

	address := common.HexToAddress("0x7033a512639119c759a51b250bfa461ae100894b")
	instance, err := solo.NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}
	wei := big.NewInt(int64(math.Pow(10, 16)))
	tx, err := instance.Deposit(auth, big.NewInt(6), wei.Mul(wei, big.NewInt(1)))
	if err != nil {
		log.Fatal("invoke Deposit failed: ", err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}
