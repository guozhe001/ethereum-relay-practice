// Package model 定义结构体
package model

// Transaction 交易结构体
type Transaction struct {
	BlockHash        string `json:"blockHash"`        //blockHash: DATA, 32 Bytes - hash of the block where this transaction was in. null when its pending.
	BlockNumber      string `json:"blockNumber"`      // blockNumber: QUANTITY - block number where this transaction was in.null when its pending.
	From             string `json:"from"`             // from: DATA, 20 Bytes - address of the sender.
	Gas              string `json:"gas"`              // gas: QUANTITY - gas provided by the sender.
	GasPrice         string `json:"gasPrice"`         // gasPrice: QUANTITY - gas price provided by the sender in Wei.
	Hash             string `json:"hash"`             // hash: DATA, 32 Bytes - hash of the transaction.
	Input            string `json:"input"`            // input: DATA - the data send along with the transaction.
	Nonce            string `json:"nonce"`            // nonce: QUANTITY - the number of transactions made by the sender prior to this one.
	To               string `json:"to"`               // to: DATA, 20 Bytes - address of the receiver.null when its a contract creation transaction.
	TransactionIndex string `json:"transactionIndex"` // transactionIndex: QUANTITY - integer of the transactions index position in the block.null when its pending.
	Value            string `json:"value"`            // value: QUANTITY - value transferred in Wei.
	V                string `json:"v"`                // v: QUANTITY - ECDSA recovery id
	R                string `json:"r"`                // r: DATA, 32 Bytes - ECDSA signature r
	S                string `json:"s"`                // s: DATA, 32 Bytes - ECDSA signature s
}
