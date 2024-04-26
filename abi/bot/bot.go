// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bot

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BotMetaData contains all meta data concerning the Bot contract.
var BotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizedAssets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapData\",\"type\":\"bytes\"}],\"name\":\"morphoLiquidate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BotABI is the input ABI used to generate the binding from.
// Deprecated: Use BotMetaData.ABI instead.
var BotABI = BotMetaData.ABI

// Bot is an auto generated Go binding around an Ethereum contract.
type Bot struct {
	BotCaller     // Read-only binding to the contract
	BotTransactor // Write-only binding to the contract
	BotFilterer   // Log filterer for contract events
}

// BotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BotSession struct {
	Contract     *Bot              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BotCallerSession struct {
	Contract *BotCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BotTransactorSession struct {
	Contract     *BotTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BotRaw struct {
	Contract *Bot // Generic contract binding to access the raw methods on
}

// BotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BotCallerRaw struct {
	Contract *BotCaller // Generic read-only contract binding to access the raw methods on
}

// BotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BotTransactorRaw struct {
	Contract *BotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBot creates a new instance of Bot, bound to a specific deployed contract.
func NewBot(address common.Address, backend bind.ContractBackend) (*Bot, error) {
	contract, err := bindBot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bot{BotCaller: BotCaller{contract: contract}, BotTransactor: BotTransactor{contract: contract}, BotFilterer: BotFilterer{contract: contract}}, nil
}

// NewBotCaller creates a new read-only instance of Bot, bound to a specific deployed contract.
func NewBotCaller(address common.Address, caller bind.ContractCaller) (*BotCaller, error) {
	contract, err := bindBot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BotCaller{contract: contract}, nil
}

// NewBotTransactor creates a new write-only instance of Bot, bound to a specific deployed contract.
func NewBotTransactor(address common.Address, transactor bind.ContractTransactor) (*BotTransactor, error) {
	contract, err := bindBot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BotTransactor{contract: contract}, nil
}

// NewBotFilterer creates a new log filterer instance of Bot, bound to a specific deployed contract.
func NewBotFilterer(address common.Address, filterer bind.ContractFilterer) (*BotFilterer, error) {
	contract, err := bindBot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BotFilterer{contract: contract}, nil
}

// bindBot binds a generic wrapper to an already deployed contract.
func bindBot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bot *BotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bot.Contract.BotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bot *BotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bot.Contract.BotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bot *BotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bot.Contract.BotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bot *BotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bot *BotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bot *BotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bot.Contract.contract.Transact(opts, method, params...)
}

// ApproveERC20 is a paid mutator transaction binding the contract method 0xa8e5e4aa.
//
// Solidity: function approveERC20(address token, address to, uint256 amount) returns()
func (_Bot *BotTransactor) ApproveERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bot.contract.Transact(opts, "approveERC20", token, to, amount)
}

// ApproveERC20 is a paid mutator transaction binding the contract method 0xa8e5e4aa.
//
// Solidity: function approveERC20(address token, address to, uint256 amount) returns()
func (_Bot *BotSession) ApproveERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bot.Contract.ApproveERC20(&_Bot.TransactOpts, token, to, amount)
}

// ApproveERC20 is a paid mutator transaction binding the contract method 0xa8e5e4aa.
//
// Solidity: function approveERC20(address token, address to, uint256 amount) returns()
func (_Bot *BotTransactorSession) ApproveERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bot.Contract.ApproveERC20(&_Bot.TransactOpts, token, to, amount)
}

// MorphoLiquidate is a paid mutator transaction binding the contract method 0xc5a734b3.
//
// Solidity: function morphoLiquidate(bytes32 id, address borrower, uint256 seizedAssets, address pair, bytes swapData) payable returns()
func (_Bot *BotTransactor) MorphoLiquidate(opts *bind.TransactOpts, id [32]byte, borrower common.Address, seizedAssets *big.Int, pair common.Address, swapData []byte) (*types.Transaction, error) {
	return _Bot.contract.Transact(opts, "morphoLiquidate", id, borrower, seizedAssets, pair, swapData)
}

// MorphoLiquidate is a paid mutator transaction binding the contract method 0xc5a734b3.
//
// Solidity: function morphoLiquidate(bytes32 id, address borrower, uint256 seizedAssets, address pair, bytes swapData) payable returns()
func (_Bot *BotSession) MorphoLiquidate(id [32]byte, borrower common.Address, seizedAssets *big.Int, pair common.Address, swapData []byte) (*types.Transaction, error) {
	return _Bot.Contract.MorphoLiquidate(&_Bot.TransactOpts, id, borrower, seizedAssets, pair, swapData)
}

// MorphoLiquidate is a paid mutator transaction binding the contract method 0xc5a734b3.
//
// Solidity: function morphoLiquidate(bytes32 id, address borrower, uint256 seizedAssets, address pair, bytes swapData) payable returns()
func (_Bot *BotTransactorSession) MorphoLiquidate(id [32]byte, borrower common.Address, seizedAssets *big.Int, pair common.Address, swapData []byte) (*types.Transaction, error) {
	return _Bot.Contract.MorphoLiquidate(&_Bot.TransactOpts, id, borrower, seizedAssets, pair, swapData)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amount) returns()
func (_Bot *BotTransactor) WithdrawERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bot.contract.Transact(opts, "withdrawERC20", token, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amount) returns()
func (_Bot *BotSession) WithdrawERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bot.Contract.WithdrawERC20(&_Bot.TransactOpts, token, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amount) returns()
func (_Bot *BotTransactorSession) WithdrawERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bot.Contract.WithdrawERC20(&_Bot.TransactOpts, token, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Bot *BotTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bot.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Bot *BotSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _Bot.Contract.WithdrawETH(&_Bot.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Bot *BotTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _Bot.Contract.WithdrawETH(&_Bot.TransactOpts, amount)
}
