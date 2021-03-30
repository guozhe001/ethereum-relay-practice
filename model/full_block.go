package model

type Block struct {
	LogsBloom        string        `json:"logsBloom"`
	TotalDifficulty  string        `json:"totalDifficulty"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	ExtraData        string        `json:"extraData"`
	Transactions     []interface{} `json:"transactions"`
	Nonce            string        `json:"nonce"`
	Miner            string        `json:"miner"`
	Difficulty       string        `json:"difficulty"`
	GasLimit         string        `json:"gasLimit"`
	Number           string        `json:"number"`
	GasUsed          string        `json:"gasUsed"`
	Uncles           []string      `json:"uncles"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	Size             string        `json:"size"`
	TransactionsRoot string        `json:"transactionsRoot"`
	StateRoot        string        `json:"stateRoot"`
	MixHash          string        `json:"mixHash"`
	ParentHash       string        `json:"parentHash"`
	Hash             string        `json:"hash"`
	Timestamp        string        `json:"timestamp"`
}
