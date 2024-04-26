// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package irm

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

// Market is an auto generated low-level Go binding around an user-defined struct.
type Market struct {
	TotalSupplyAssets *big.Int
	TotalSupplyShares *big.Int
	TotalBorrowAssets *big.Int
	TotalBorrowShares *big.Int
	LastUpdate        *big.Int
	Fee               *big.Int
}

// MarketParams is an auto generated low-level Go binding around an user-defined struct.
type MarketParams struct {
	LoanToken       common.Address
	CollateralToken common.Address
	Oracle          common.Address
	Irm             common.Address
	Lltv            *big.Int
}

// IrmMetaData contains all meta data concerning the Irm contract.
var IrmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"totalSupplyAssets\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalSupplyShares\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalBorrowAssets\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalBorrowShares\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastUpdate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fee\",\"type\":\"uint128\"}],\"internalType\":\"structMarket\",\"name\":\"market\",\"type\":\"tuple\"}],\"name\":\"borrowRateView\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IrmABI is the input ABI used to generate the binding from.
// Deprecated: Use IrmMetaData.ABI instead.
var IrmABI = IrmMetaData.ABI

// Irm is an auto generated Go binding around an Ethereum contract.
type Irm struct {
	IrmCaller     // Read-only binding to the contract
	IrmTransactor // Write-only binding to the contract
	IrmFilterer   // Log filterer for contract events
}

// IrmCaller is an auto generated read-only Go binding around an Ethereum contract.
type IrmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IrmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IrmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IrmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IrmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IrmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IrmSession struct {
	Contract     *Irm              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IrmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IrmCallerSession struct {
	Contract *IrmCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IrmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IrmTransactorSession struct {
	Contract     *IrmTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IrmRaw is an auto generated low-level Go binding around an Ethereum contract.
type IrmRaw struct {
	Contract *Irm // Generic contract binding to access the raw methods on
}

// IrmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IrmCallerRaw struct {
	Contract *IrmCaller // Generic read-only contract binding to access the raw methods on
}

// IrmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IrmTransactorRaw struct {
	Contract *IrmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIrm creates a new instance of Irm, bound to a specific deployed contract.
func NewIrm(address common.Address, backend bind.ContractBackend) (*Irm, error) {
	contract, err := bindIrm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Irm{IrmCaller: IrmCaller{contract: contract}, IrmTransactor: IrmTransactor{contract: contract}, IrmFilterer: IrmFilterer{contract: contract}}, nil
}

// NewIrmCaller creates a new read-only instance of Irm, bound to a specific deployed contract.
func NewIrmCaller(address common.Address, caller bind.ContractCaller) (*IrmCaller, error) {
	contract, err := bindIrm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IrmCaller{contract: contract}, nil
}

// NewIrmTransactor creates a new write-only instance of Irm, bound to a specific deployed contract.
func NewIrmTransactor(address common.Address, transactor bind.ContractTransactor) (*IrmTransactor, error) {
	contract, err := bindIrm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IrmTransactor{contract: contract}, nil
}

// NewIrmFilterer creates a new log filterer instance of Irm, bound to a specific deployed contract.
func NewIrmFilterer(address common.Address, filterer bind.ContractFilterer) (*IrmFilterer, error) {
	contract, err := bindIrm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IrmFilterer{contract: contract}, nil
}

// bindIrm binds a generic wrapper to an already deployed contract.
func bindIrm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IrmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Irm *IrmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Irm.Contract.IrmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Irm *IrmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Irm.Contract.IrmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Irm *IrmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Irm.Contract.IrmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Irm *IrmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Irm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Irm *IrmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Irm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Irm *IrmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Irm.Contract.contract.Transact(opts, method, params...)
}

// BorrowRateView is a free data retrieval call binding the contract method 0x8c00bf6b.
//
// Solidity: function borrowRateView((address,address,address,address,uint256) marketParams, (uint128,uint128,uint128,uint128,uint128,uint128) market) view returns(uint256)
func (_Irm *IrmCaller) BorrowRateView(opts *bind.CallOpts, marketParams MarketParams, market Market) (*big.Int, error) {
	var out []interface{}
	err := _Irm.contract.Call(opts, &out, "borrowRateView", marketParams, market)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRateView is a free data retrieval call binding the contract method 0x8c00bf6b.
//
// Solidity: function borrowRateView((address,address,address,address,uint256) marketParams, (uint128,uint128,uint128,uint128,uint128,uint128) market) view returns(uint256)
func (_Irm *IrmSession) BorrowRateView(marketParams MarketParams, market Market) (*big.Int, error) {
	return _Irm.Contract.BorrowRateView(&_Irm.CallOpts, marketParams, market)
}

// BorrowRateView is a free data retrieval call binding the contract method 0x8c00bf6b.
//
// Solidity: function borrowRateView((address,address,address,address,uint256) marketParams, (uint128,uint128,uint128,uint128,uint128,uint128) market) view returns(uint256)
func (_Irm *IrmCallerSession) BorrowRateView(marketParams MarketParams, market Market) (*big.Int, error) {
	return _Irm.Contract.BorrowRateView(&_Irm.CallOpts, marketParams, market)
}
