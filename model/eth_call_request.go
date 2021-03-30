package model

import "github.com/ethereum/go-ethereum/common"

// EthCallRequest eth_all方法的请求参数
type EthCallRequest struct {
	From     common.Address `json:"from"`     // from: DATA, 20 Bytes - (optional) The address the transaction is sent from.
	To       common.Address `json:"to"`       // to: DATA, 20 Bytes - The address the transaction is directed to.
	Gas      string         `json:"gas"`      // gas: QUANTITY - (optional) Integer of the gas provided for the transaction execution. eth_call consumes zero gas, but this parameter may be needed by some executions.
	GasPrice string         `json:"gasPrice"` // gasPrice: QUANTITY - (optional) Integer of the gasPrice used for each paid gas
	Value    string         `json:"value"`    // value: QUANTITY - (optional) Integer of the value sent with this transaction
	Data     string         `json:"data"`     // data: DATA - (optional) Hash of the method signature and encoded parameters. For details see Ethereum Contract ABI in the Solidity documentation
}
