package model

// ERC20BalanceRequest ERC20余额查询请求
type ERC20BalanceRequest struct {
	ContractAddress string // 合约地址
	UserAddress     string // 用户地址
	ContractDecimal int64  // 合约的Decimal
}

// ERC20BalanceResponse ERC20余额查询响应
type ERC20BalanceResponse struct {
	ContractAddress string  // 合约地址
	UserAddress     string  // 用户地址
	Balance         *string // 用户在合约地址的余额
}
