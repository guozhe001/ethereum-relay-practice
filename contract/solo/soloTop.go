// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solo

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"funcName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"MonitorEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"funcName\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"MonitorEventCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"cancellingOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIBankController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"flashloan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mulSig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingBorrow\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mulSig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"proposeNewAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"setPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setUnpaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"tokenIn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"tokenOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawTokens\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contract *ContractCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contract *ContractSession) Admin() (common.Address, error) {
	return _Contract.Contract.Admin(&_Contract.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contract *ContractCallerSession) Admin() (common.Address, error) {
	return _Contract.Contract.Admin(&_Contract.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Contract *ContractCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Contract *ContractSession) Controller() (common.Address, error) {
	return _Contract.Contract.Controller(&_Contract.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Contract *ContractCallerSession) Controller() (common.Address, error) {
	return _Contract.Contract.Controller(&_Contract.CallOpts)
}

// MulSig is a free data retrieval call binding the contract method 0x62757215.
//
// Solidity: function mulSig() view returns(address)
func (_Contract *ContractCaller) MulSig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "mulSig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MulSig is a free data retrieval call binding the contract method 0x62757215.
//
// Solidity: function mulSig() view returns(address)
func (_Contract *ContractSession) MulSig() (common.Address, error) {
	return _Contract.Contract.MulSig(&_Contract.CallOpts)
}

// MulSig is a free data retrieval call binding the contract method 0x62757215.
//
// Solidity: function mulSig() view returns(address)
func (_Contract *ContractCallerSession) MulSig() (common.Address, error) {
	return _Contract.Contract.MulSig(&_Contract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractSession) Paused() (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractCallerSession) Paused() (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Contract *ContractCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Contract *ContractSession) Pauser() (common.Address, error) {
	return _Contract.Contract.Pauser(&_Contract.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Contract *ContractCallerSession) Pauser() (common.Address, error) {
	return _Contract.Contract.Pauser(&_Contract.CallOpts)
}

// ProposedAdmin is a free data retrieval call binding the contract method 0x32f751ec.
//
// Solidity: function proposedAdmin() view returns(address)
func (_Contract *ContractCaller) ProposedAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "proposedAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedAdmin is a free data retrieval call binding the contract method 0x32f751ec.
//
// Solidity: function proposedAdmin() view returns(address)
func (_Contract *ContractSession) ProposedAdmin() (common.Address, error) {
	return _Contract.Contract.ProposedAdmin(&_Contract.CallOpts)
}

// ProposedAdmin is a free data retrieval call binding the contract method 0x32f751ec.
//
// Solidity: function proposedAdmin() view returns(address)
func (_Contract *ContractCallerSession) ProposedAdmin() (common.Address, error) {
	return _Contract.Contract.ProposedAdmin(&_Contract.CallOpts)
}

// MonitorEventCallback is a paid mutator transaction binding the contract method 0x7bc597aa.
//
// Solidity: function MonitorEventCallback(bytes32 funcName, bytes payload) returns()
func (_Contract *ContractTransactor) MonitorEventCallback(opts *bind.TransactOpts, funcName [32]byte, payload []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "MonitorEventCallback", funcName, payload)
}

// MonitorEventCallback is a paid mutator transaction binding the contract method 0x7bc597aa.
//
// Solidity: function MonitorEventCallback(bytes32 funcName, bytes payload) returns()
func (_Contract *ContractSession) MonitorEventCallback(funcName [32]byte, payload []byte) (*types.Transaction, error) {
	return _Contract.Contract.MonitorEventCallback(&_Contract.TransactOpts, funcName, payload)
}

// MonitorEventCallback is a paid mutator transaction binding the contract method 0x7bc597aa.
//
// Solidity: function MonitorEventCallback(bytes32 funcName, bytes payload) returns()
func (_Contract *ContractTransactorSession) MonitorEventCallback(funcName [32]byte, payload []byte) (*types.Transaction, error) {
	return _Contract.Contract.MonitorEventCallback(&_Contract.TransactOpts, funcName, payload)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 underlying, uint256 borrowAmount) returns()
func (_Contract *ContractTransactor) Borrow(opts *bind.TransactOpts, underlying *big.Int, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "borrow", underlying, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 underlying, uint256 borrowAmount) returns()
func (_Contract *ContractSession) Borrow(underlying *big.Int, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Borrow(&_Contract.TransactOpts, underlying, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 underlying, uint256 borrowAmount) returns()
func (_Contract *ContractTransactorSession) Borrow(underlying *big.Int, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Borrow(&_Contract.TransactOpts, underlying, borrowAmount)
}

// CancellingOut is a paid mutator transaction binding the contract method 0x15f14eaa.
//
// Solidity: function cancellingOut(address token) returns()
func (_Contract *ContractTransactor) CancellingOut(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancellingOut", token)
}

// CancellingOut is a paid mutator transaction binding the contract method 0x15f14eaa.
//
// Solidity: function cancellingOut(address token) returns()
func (_Contract *ContractSession) CancellingOut(token common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CancellingOut(&_Contract.TransactOpts, token)
}

// CancellingOut is a paid mutator transaction binding the contract method 0x15f14eaa.
//
// Solidity: function cancellingOut(address token) returns()
func (_Contract *ContractTransactorSession) CancellingOut(token common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CancellingOut(&_Contract.TransactOpts, token)
}

// ClaimAdministration is a paid mutator transaction binding the contract method 0xad2cb239.
//
// Solidity: function claimAdministration() returns()
func (_Contract *ContractTransactor) ClaimAdministration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimAdministration")
}

// ClaimAdministration is a paid mutator transaction binding the contract method 0xad2cb239.
//
// Solidity: function claimAdministration() returns()
func (_Contract *ContractSession) ClaimAdministration() (*types.Transaction, error) {
	return _Contract.Contract.ClaimAdministration(&_Contract.TransactOpts)
}

// ClaimAdministration is a paid mutator transaction binding the contract method 0xad2cb239.
//
// Solidity: function claimAdministration() returns()
func (_Contract *ContractTransactorSession) ClaimAdministration() (*types.Transaction, error) {
	return _Contract.Contract.ClaimAdministration(&_Contract.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) payable returns()
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts, pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deposit", pid, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) payable returns()
func (_Contract *ContractSession) Deposit(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, pid, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) payable returns()
func (_Contract *ContractTransactorSession) Deposit(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, pid, amount)
}

// Flashloan is a paid mutator transaction binding the contract method 0x63ad2c41.
//
// Solidity: function flashloan(address receiver, address token, uint256 amount, bytes params) returns()
func (_Contract *ContractTransactor) Flashloan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "flashloan", receiver, token, amount, params)
}

// Flashloan is a paid mutator transaction binding the contract method 0x63ad2c41.
//
// Solidity: function flashloan(address receiver, address token, uint256 amount, bytes params) returns()
func (_Contract *ContractSession) Flashloan(receiver common.Address, token common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Contract.Contract.Flashloan(&_Contract.TransactOpts, receiver, token, amount, params)
}

// Flashloan is a paid mutator transaction binding the contract method 0x63ad2c41.
//
// Solidity: function flashloan(address receiver, address token, uint256 amount, bytes params) returns()
func (_Contract *ContractTransactorSession) Flashloan(receiver common.Address, token common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Contract.Contract.Flashloan(&_Contract.TransactOpts, receiver, token, amount, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _controller, address _mulSig) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _controller common.Address, _mulSig common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _controller, _mulSig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _controller, address _mulSig) returns()
func (_Contract *ContractSession) Initialize(_controller common.Address, _mulSig common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _controller, _mulSig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _controller, address _mulSig) returns()
func (_Contract *ContractTransactorSession) Initialize(_controller common.Address, _mulSig common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _controller, _mulSig)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xe61604cf.
//
// Solidity: function liquidateBorrow(address borrower, address underlyingBorrow, address underlyingCollateral, uint256 repayAmount) payable returns()
func (_Contract *ContractTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, underlyingBorrow common.Address, underlyingCollateral common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "liquidateBorrow", borrower, underlyingBorrow, underlyingCollateral, repayAmount)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xe61604cf.
//
// Solidity: function liquidateBorrow(address borrower, address underlyingBorrow, address underlyingCollateral, uint256 repayAmount) payable returns()
func (_Contract *ContractSession) LiquidateBorrow(borrower common.Address, underlyingBorrow common.Address, underlyingCollateral common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LiquidateBorrow(&_Contract.TransactOpts, borrower, underlyingBorrow, underlyingCollateral, repayAmount)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xe61604cf.
//
// Solidity: function liquidateBorrow(address borrower, address underlyingBorrow, address underlyingCollateral, uint256 repayAmount) payable returns()
func (_Contract *ContractTransactorSession) LiquidateBorrow(borrower common.Address, underlyingBorrow common.Address, underlyingCollateral common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LiquidateBorrow(&_Contract.TransactOpts, borrower, underlyingBorrow, underlyingCollateral, repayAmount)
}

// ProposeNewAdmin is a paid mutator transaction binding the contract method 0xa6376746.
//
// Solidity: function proposeNewAdmin(address admin_) returns()
func (_Contract *ContractTransactor) ProposeNewAdmin(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "proposeNewAdmin", admin_)
}

// ProposeNewAdmin is a paid mutator transaction binding the contract method 0xa6376746.
//
// Solidity: function proposeNewAdmin(address admin_) returns()
func (_Contract *ContractSession) ProposeNewAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ProposeNewAdmin(&_Contract.TransactOpts, admin_)
}

// ProposeNewAdmin is a paid mutator transaction binding the contract method 0xa6376746.
//
// Solidity: function proposeNewAdmin(address admin_) returns()
func (_Contract *ContractTransactorSession) ProposeNewAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ProposeNewAdmin(&_Contract.TransactOpts, admin_)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address token, uint256 repayAmount) payable returns(uint256)
func (_Contract *ContractTransactor) Repay(opts *bind.TransactOpts, token common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "repay", token, repayAmount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address token, uint256 repayAmount) payable returns(uint256)
func (_Contract *ContractSession) Repay(token common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Repay(&_Contract.TransactOpts, token, repayAmount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address token, uint256 repayAmount) payable returns(uint256)
func (_Contract *ContractTransactorSession) Repay(token common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Repay(&_Contract.TransactOpts, token, repayAmount)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Contract *ContractTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Contract *ContractSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetController(&_Contract.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Contract *ContractTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetController(&_Contract.TransactOpts, _controller)
}

// SetPaused is a paid mutator transaction binding the contract method 0x37a66d85.
//
// Solidity: function setPaused() returns()
func (_Contract *ContractTransactor) SetPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPaused")
}

// SetPaused is a paid mutator transaction binding the contract method 0x37a66d85.
//
// Solidity: function setPaused() returns()
func (_Contract *ContractSession) SetPaused() (*types.Transaction, error) {
	return _Contract.Contract.SetPaused(&_Contract.TransactOpts)
}

// SetPaused is a paid mutator transaction binding the contract method 0x37a66d85.
//
// Solidity: function setPaused() returns()
func (_Contract *ContractTransactorSession) SetPaused() (*types.Transaction, error) {
	return _Contract.Contract.SetPaused(&_Contract.TransactOpts)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Contract *ContractTransactor) SetPauser(opts *bind.TransactOpts, _pauser common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPauser", _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Contract *ContractSession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetPauser(&_Contract.TransactOpts, _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Contract *ContractTransactorSession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetPauser(&_Contract.TransactOpts, _pauser)
}

// SetUnpaused is a paid mutator transaction binding the contract method 0x3c89edce.
//
// Solidity: function setUnpaused() returns()
func (_Contract *ContractTransactor) SetUnpaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setUnpaused")
}

// SetUnpaused is a paid mutator transaction binding the contract method 0x3c89edce.
//
// Solidity: function setUnpaused() returns()
func (_Contract *ContractSession) SetUnpaused() (*types.Transaction, error) {
	return _Contract.Contract.SetUnpaused(&_Contract.TransactOpts)
}

// SetUnpaused is a paid mutator transaction binding the contract method 0x3c89edce.
//
// Solidity: function setUnpaused() returns()
func (_Contract *ContractTransactorSession) SetUnpaused() (*types.Transaction, error) {
	return _Contract.Contract.SetUnpaused(&_Contract.TransactOpts)
}

// TokenIn is a paid mutator transaction binding the contract method 0x7f51975f.
//
// Solidity: function tokenIn(address token, uint256 amountIn) payable returns()
func (_Contract *ContractTransactor) TokenIn(opts *bind.TransactOpts, token common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "tokenIn", token, amountIn)
}

// TokenIn is a paid mutator transaction binding the contract method 0x7f51975f.
//
// Solidity: function tokenIn(address token, uint256 amountIn) payable returns()
func (_Contract *ContractSession) TokenIn(token common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TokenIn(&_Contract.TransactOpts, token, amountIn)
}

// TokenIn is a paid mutator transaction binding the contract method 0x7f51975f.
//
// Solidity: function tokenIn(address token, uint256 amountIn) payable returns()
func (_Contract *ContractTransactorSession) TokenIn(token common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TokenIn(&_Contract.TransactOpts, token, amountIn)
}

// TokenOut is a paid mutator transaction binding the contract method 0x1beb7f5c.
//
// Solidity: function tokenOut(address token, uint256 amountOut) returns()
func (_Contract *ContractTransactor) TokenOut(opts *bind.TransactOpts, token common.Address, amountOut *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "tokenOut", token, amountOut)
}

// TokenOut is a paid mutator transaction binding the contract method 0x1beb7f5c.
//
// Solidity: function tokenOut(address token, uint256 amountOut) returns()
func (_Contract *ContractSession) TokenOut(token common.Address, amountOut *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TokenOut(&_Contract.TransactOpts, token, amountOut)
}

// TokenOut is a paid mutator transaction binding the contract method 0x1beb7f5c.
//
// Solidity: function tokenOut(address token, uint256 amountOut) returns()
func (_Contract *ContractTransactorSession) TokenOut(token common.Address, amountOut *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TokenOut(&_Contract.TransactOpts, token, amountOut)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address underlying, uint256 withdrawTokens) returns(uint256)
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, underlying common.Address, withdrawTokens *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", underlying, withdrawTokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address underlying, uint256 withdrawTokens) returns(uint256)
func (_Contract *ContractSession) Withdraw(underlying common.Address, withdrawTokens *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, underlying, withdrawTokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address underlying, uint256 withdrawTokens) returns(uint256)
func (_Contract *ContractTransactorSession) Withdraw(underlying common.Address, withdrawTokens *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, underlying, withdrawTokens)
}

// WithdrawUnderlying is a paid mutator transaction binding the contract method 0xfdb87252.
//
// Solidity: function withdrawUnderlying(address underlying, uint256 withdrawAmount) returns(uint256)
func (_Contract *ContractTransactor) WithdrawUnderlying(opts *bind.TransactOpts, underlying common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawUnderlying", underlying, withdrawAmount)
}

// WithdrawUnderlying is a paid mutator transaction binding the contract method 0xfdb87252.
//
// Solidity: function withdrawUnderlying(address underlying, uint256 withdrawAmount) returns(uint256)
func (_Contract *ContractSession) WithdrawUnderlying(underlying common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawUnderlying(&_Contract.TransactOpts, underlying, withdrawAmount)
}

// WithdrawUnderlying is a paid mutator transaction binding the contract method 0xfdb87252.
//
// Solidity: function withdrawUnderlying(address underlying, uint256 withdrawAmount) returns(uint256)
func (_Contract *ContractTransactorSession) WithdrawUnderlying(underlying common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawUnderlying(&_Contract.TransactOpts, underlying, withdrawAmount)
}

// ContractFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Contract contract.
type ContractFlashLoanIterator struct {
	Event *ContractFlashLoan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFlashLoan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFlashLoan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFlashLoan represents a FlashLoan event raised by the Contract contract.
type ContractFlashLoan struct {
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed receiver, address indexed token, uint256 amount, uint256 fee)
func (_Contract *ContractFilterer) FilterFlashLoan(opts *bind.FilterOpts, receiver []common.Address, token []common.Address) (*ContractFlashLoanIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FlashLoan", receiverRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractFlashLoanIterator{contract: _Contract.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed receiver, address indexed token, uint256 amount, uint256 fee)
func (_Contract *ContractFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *ContractFlashLoan, receiver []common.Address, token []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FlashLoan", receiverRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFlashLoan)
				if err := _Contract.contract.UnpackLog(event, "FlashLoan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashLoan is a log parse operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed receiver, address indexed token, uint256 amount, uint256 fee)
func (_Contract *ContractFilterer) ParseFlashLoan(log types.Log) (*ContractFlashLoan, error) {
	event := new(ContractFlashLoan)
	if err := _Contract.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMonitorEventIterator is returned from FilterMonitorEvent and is used to iterate over the raw logs and unpacked data for MonitorEvent events raised by the Contract contract.
type ContractMonitorEventIterator struct {
	Event *ContractMonitorEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMonitorEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMonitorEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMonitorEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMonitorEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMonitorEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMonitorEvent represents a MonitorEvent event raised by the Contract contract.
type ContractMonitorEvent struct {
	FuncName [32]byte
	Payload  []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMonitorEvent is a free log retrieval operation binding the contract event 0x8df0b9408c77393ca0d6ccf0e770ef1432b46111947c518125b420a00207381e.
//
// Solidity: event MonitorEvent(bytes32 indexed funcName, bytes payload)
func (_Contract *ContractFilterer) FilterMonitorEvent(opts *bind.FilterOpts, funcName [][32]byte) (*ContractMonitorEventIterator, error) {

	var funcNameRule []interface{}
	for _, funcNameItem := range funcName {
		funcNameRule = append(funcNameRule, funcNameItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MonitorEvent", funcNameRule)
	if err != nil {
		return nil, err
	}
	return &ContractMonitorEventIterator{contract: _Contract.contract, event: "MonitorEvent", logs: logs, sub: sub}, nil
}

// WatchMonitorEvent is a free log subscription operation binding the contract event 0x8df0b9408c77393ca0d6ccf0e770ef1432b46111947c518125b420a00207381e.
//
// Solidity: event MonitorEvent(bytes32 indexed funcName, bytes payload)
func (_Contract *ContractFilterer) WatchMonitorEvent(opts *bind.WatchOpts, sink chan<- *ContractMonitorEvent, funcName [][32]byte) (event.Subscription, error) {

	var funcNameRule []interface{}
	for _, funcNameItem := range funcName {
		funcNameRule = append(funcNameRule, funcNameItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MonitorEvent", funcNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMonitorEvent)
				if err := _Contract.contract.UnpackLog(event, "MonitorEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMonitorEvent is a log parse operation binding the contract event 0x8df0b9408c77393ca0d6ccf0e770ef1432b46111947c518125b420a00207381e.
//
// Solidity: event MonitorEvent(bytes32 indexed funcName, bytes payload)
func (_Contract *ContractFilterer) ParseMonitorEvent(log types.Log) (*ContractMonitorEvent, error) {
	event := new(ContractMonitorEvent)
	if err := _Contract.contract.UnpackLog(event, "MonitorEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
