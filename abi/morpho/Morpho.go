// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package morpho

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

// IMorphoMarket is an auto generated low-level Go binding around an user-defined struct.
type IMorphoMarket struct {
	TotalSupplyAssets *big.Int
	TotalSupplyShares *big.Int
	TotalBorrowAssets *big.Int
	TotalBorrowShares *big.Int
	LastUpdate        *big.Int
	Fee               *big.Int
}

// IMorphoMarketParams is an auto generated low-level Go binding around an user-defined struct.
type IMorphoMarketParams struct {
	LoanToken       common.Address
	CollateralToken common.Address
	Oracle          common.Address
	Irm             common.Address
	Lltv            *big.Int
}

// IMorphoPosition is an auto generated low-level Go binding around an user-defined struct.
type IMorphoPosition struct {
	SupplyShares *big.Int
	BorrowShares *big.Int
	Collateral   *big.Int
}

// MorphoMetaData contains all meta data concerning the Morpho contract.
var MorphoMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIMorpho.MarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"CreateMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"idToMarketParams\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structIMorpho.MarketParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"market\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"totalSupplyAssets\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalSupplyShares\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalBorrowAssets\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalBorrowShares\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastUpdate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fee\",\"type\":\"uint128\"}],\"internalType\":\"structIMorpho.Market\",\"name\":\"m\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"position\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"supplyShares\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"borrowShares\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"collateral\",\"type\":\"uint128\"}],\"internalType\":\"structIMorpho.Position\",\"name\":\"p\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MorphoABI is the input ABI used to generate the binding from.
// Deprecated: Use MorphoMetaData.ABI instead.
var MorphoABI = MorphoMetaData.ABI

// Morpho is an auto generated Go binding around an Ethereum contract.
type Morpho struct {
	MorphoCaller     // Read-only binding to the contract
	MorphoTransactor // Write-only binding to the contract
	MorphoFilterer   // Log filterer for contract events
}

// MorphoCaller is an auto generated read-only Go binding around an Ethereum contract.
type MorphoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MorphoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MorphoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MorphoSession struct {
	Contract     *Morpho           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MorphoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MorphoCallerSession struct {
	Contract *MorphoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MorphoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MorphoTransactorSession struct {
	Contract     *MorphoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MorphoRaw is an auto generated low-level Go binding around an Ethereum contract.
type MorphoRaw struct {
	Contract *Morpho // Generic contract binding to access the raw methods on
}

// MorphoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MorphoCallerRaw struct {
	Contract *MorphoCaller // Generic read-only contract binding to access the raw methods on
}

// MorphoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MorphoTransactorRaw struct {
	Contract *MorphoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMorpho creates a new instance of Morpho, bound to a specific deployed contract.
func NewMorpho(address common.Address, backend bind.ContractBackend) (*Morpho, error) {
	contract, err := bindMorpho(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Morpho{MorphoCaller: MorphoCaller{contract: contract}, MorphoTransactor: MorphoTransactor{contract: contract}, MorphoFilterer: MorphoFilterer{contract: contract}}, nil
}

// NewMorphoCaller creates a new read-only instance of Morpho, bound to a specific deployed contract.
func NewMorphoCaller(address common.Address, caller bind.ContractCaller) (*MorphoCaller, error) {
	contract, err := bindMorpho(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoCaller{contract: contract}, nil
}

// NewMorphoTransactor creates a new write-only instance of Morpho, bound to a specific deployed contract.
func NewMorphoTransactor(address common.Address, transactor bind.ContractTransactor) (*MorphoTransactor, error) {
	contract, err := bindMorpho(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoTransactor{contract: contract}, nil
}

// NewMorphoFilterer creates a new log filterer instance of Morpho, bound to a specific deployed contract.
func NewMorphoFilterer(address common.Address, filterer bind.ContractFilterer) (*MorphoFilterer, error) {
	contract, err := bindMorpho(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MorphoFilterer{contract: contract}, nil
}

// bindMorpho binds a generic wrapper to an already deployed contract.
func bindMorpho(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MorphoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Morpho *MorphoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Morpho.Contract.MorphoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Morpho *MorphoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Morpho.Contract.MorphoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Morpho *MorphoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Morpho.Contract.MorphoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Morpho *MorphoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Morpho.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Morpho *MorphoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Morpho.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Morpho *MorphoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Morpho.Contract.contract.Transact(opts, method, params...)
}

// IdToMarketParams is a free data retrieval call binding the contract method 0x2c3c9157.
//
// Solidity: function idToMarketParams(bytes32 id) view returns((address,address,address,address,uint256))
func (_Morpho *MorphoCaller) IdToMarketParams(opts *bind.CallOpts, id [32]byte) (IMorphoMarketParams, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "idToMarketParams", id)

	if err != nil {
		return *new(IMorphoMarketParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IMorphoMarketParams)).(*IMorphoMarketParams)

	return out0, err

}

// IdToMarketParams is a free data retrieval call binding the contract method 0x2c3c9157.
//
// Solidity: function idToMarketParams(bytes32 id) view returns((address,address,address,address,uint256))
func (_Morpho *MorphoSession) IdToMarketParams(id [32]byte) (IMorphoMarketParams, error) {
	return _Morpho.Contract.IdToMarketParams(&_Morpho.CallOpts, id)
}

// IdToMarketParams is a free data retrieval call binding the contract method 0x2c3c9157.
//
// Solidity: function idToMarketParams(bytes32 id) view returns((address,address,address,address,uint256))
func (_Morpho *MorphoCallerSession) IdToMarketParams(id [32]byte) (IMorphoMarketParams, error) {
	return _Morpho.Contract.IdToMarketParams(&_Morpho.CallOpts, id)
}

// Market is a free data retrieval call binding the contract method 0x5c60e39a.
//
// Solidity: function market(bytes32 id) view returns((uint128,uint128,uint128,uint128,uint128,uint128) m)
func (_Morpho *MorphoCaller) Market(opts *bind.CallOpts, id [32]byte) (IMorphoMarket, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "market", id)

	if err != nil {
		return *new(IMorphoMarket), err
	}

	out0 := *abi.ConvertType(out[0], new(IMorphoMarket)).(*IMorphoMarket)

	return out0, err

}

// Market is a free data retrieval call binding the contract method 0x5c60e39a.
//
// Solidity: function market(bytes32 id) view returns((uint128,uint128,uint128,uint128,uint128,uint128) m)
func (_Morpho *MorphoSession) Market(id [32]byte) (IMorphoMarket, error) {
	return _Morpho.Contract.Market(&_Morpho.CallOpts, id)
}

// Market is a free data retrieval call binding the contract method 0x5c60e39a.
//
// Solidity: function market(bytes32 id) view returns((uint128,uint128,uint128,uint128,uint128,uint128) m)
func (_Morpho *MorphoCallerSession) Market(id [32]byte) (IMorphoMarket, error) {
	return _Morpho.Contract.Market(&_Morpho.CallOpts, id)
}

// Position is a free data retrieval call binding the contract method 0x93c52062.
//
// Solidity: function position(bytes32 id, address user) view returns((uint256,uint128,uint128) p)
func (_Morpho *MorphoCaller) Position(opts *bind.CallOpts, id [32]byte, user common.Address) (IMorphoPosition, error) {
	var out []interface{}
	err := _Morpho.contract.Call(opts, &out, "position", id, user)

	if err != nil {
		return *new(IMorphoPosition), err
	}

	out0 := *abi.ConvertType(out[0], new(IMorphoPosition)).(*IMorphoPosition)

	return out0, err

}

// Position is a free data retrieval call binding the contract method 0x93c52062.
//
// Solidity: function position(bytes32 id, address user) view returns((uint256,uint128,uint128) p)
func (_Morpho *MorphoSession) Position(id [32]byte, user common.Address) (IMorphoPosition, error) {
	return _Morpho.Contract.Position(&_Morpho.CallOpts, id, user)
}

// Position is a free data retrieval call binding the contract method 0x93c52062.
//
// Solidity: function position(bytes32 id, address user) view returns((uint256,uint128,uint128) p)
func (_Morpho *MorphoCallerSession) Position(id [32]byte, user common.Address) (IMorphoPosition, error) {
	return _Morpho.Contract.Position(&_Morpho.CallOpts, id, user)
}

// MorphoBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Morpho contract.
type MorphoBorrowIterator struct {
	Event *MorphoBorrow // Event containing the contract specifics and raw log

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
func (it *MorphoBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoBorrow)
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
		it.Event = new(MorphoBorrow)
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
func (it *MorphoBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoBorrow represents a Borrow event raised by the Morpho contract.
type MorphoBorrow struct {
	Id       [32]byte
	Caller   common.Address
	OnBehalf common.Address
	Receiver common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x570954540bed6b1304a87dfe815a5eda4a648f7097a16240dcd85c9b5fd42a43.
//
// Solidity: event Borrow(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) FilterBorrow(opts *bind.FilterOpts, id [][32]byte, onBehalf []common.Address, receiver []common.Address) (*MorphoBorrowIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "Borrow", idRule, onBehalfRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &MorphoBorrowIterator{contract: _Morpho.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x570954540bed6b1304a87dfe815a5eda4a648f7097a16240dcd85c9b5fd42a43.
//
// Solidity: event Borrow(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *MorphoBorrow, id [][32]byte, onBehalf []common.Address, receiver []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "Borrow", idRule, onBehalfRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoBorrow)
				if err := _Morpho.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x570954540bed6b1304a87dfe815a5eda4a648f7097a16240dcd85c9b5fd42a43.
//
// Solidity: event Borrow(bytes32 indexed id, address caller, address indexed onBehalf, address indexed receiver, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) ParseBorrow(log types.Log) (*MorphoBorrow, error) {
	event := new(MorphoBorrow)
	if err := _Morpho.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoCreateMarketIterator is returned from FilterCreateMarket and is used to iterate over the raw logs and unpacked data for CreateMarket events raised by the Morpho contract.
type MorphoCreateMarketIterator struct {
	Event *MorphoCreateMarket // Event containing the contract specifics and raw log

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
func (it *MorphoCreateMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoCreateMarket)
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
		it.Event = new(MorphoCreateMarket)
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
func (it *MorphoCreateMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoCreateMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoCreateMarket represents a CreateMarket event raised by the Morpho contract.
type MorphoCreateMarket struct {
	Id           [32]byte
	MarketParams IMorphoMarketParams
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateMarket is a free log retrieval operation binding the contract event 0xac4b2400f169220b0c0afdde7a0b32e775ba727ea1cb30b35f935cdaab8683ac.
//
// Solidity: event CreateMarket(bytes32 indexed id, (address,address,address,address,uint256) marketParams)
func (_Morpho *MorphoFilterer) FilterCreateMarket(opts *bind.FilterOpts, id [][32]byte) (*MorphoCreateMarketIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "CreateMarket", idRule)
	if err != nil {
		return nil, err
	}
	return &MorphoCreateMarketIterator{contract: _Morpho.contract, event: "CreateMarket", logs: logs, sub: sub}, nil
}

// WatchCreateMarket is a free log subscription operation binding the contract event 0xac4b2400f169220b0c0afdde7a0b32e775ba727ea1cb30b35f935cdaab8683ac.
//
// Solidity: event CreateMarket(bytes32 indexed id, (address,address,address,address,uint256) marketParams)
func (_Morpho *MorphoFilterer) WatchCreateMarket(opts *bind.WatchOpts, sink chan<- *MorphoCreateMarket, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "CreateMarket", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoCreateMarket)
				if err := _Morpho.contract.UnpackLog(event, "CreateMarket", log); err != nil {
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

// ParseCreateMarket is a log parse operation binding the contract event 0xac4b2400f169220b0c0afdde7a0b32e775ba727ea1cb30b35f935cdaab8683ac.
//
// Solidity: event CreateMarket(bytes32 indexed id, (address,address,address,address,uint256) marketParams)
func (_Morpho *MorphoFilterer) ParseCreateMarket(log types.Log) (*MorphoCreateMarket, error) {
	event := new(MorphoCreateMarket)
	if err := _Morpho.contract.UnpackLog(event, "CreateMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MorphoRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the Morpho contract.
type MorphoRepayIterator struct {
	Event *MorphoRepay // Event containing the contract specifics and raw log

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
func (it *MorphoRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MorphoRepay)
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
		it.Event = new(MorphoRepay)
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
func (it *MorphoRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MorphoRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MorphoRepay represents a Repay event raised by the Morpho contract.
type MorphoRepay struct {
	Id       [32]byte
	Caller   common.Address
	OnBehalf common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x52acb05cebbd3cd39715469f22afbf5a17496295ef3bc9bb5944056c63ccaa09.
//
// Solidity: event Repay(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) FilterRepay(opts *bind.FilterOpts, id [][32]byte, caller []common.Address, onBehalf []common.Address) (*MorphoRepayIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}

	logs, sub, err := _Morpho.contract.FilterLogs(opts, "Repay", idRule, callerRule, onBehalfRule)
	if err != nil {
		return nil, err
	}
	return &MorphoRepayIterator{contract: _Morpho.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x52acb05cebbd3cd39715469f22afbf5a17496295ef3bc9bb5944056c63ccaa09.
//
// Solidity: event Repay(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *MorphoRepay, id [][32]byte, caller []common.Address, onBehalf []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var onBehalfRule []interface{}
	for _, onBehalfItem := range onBehalf {
		onBehalfRule = append(onBehalfRule, onBehalfItem)
	}

	logs, sub, err := _Morpho.contract.WatchLogs(opts, "Repay", idRule, callerRule, onBehalfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MorphoRepay)
				if err := _Morpho.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x52acb05cebbd3cd39715469f22afbf5a17496295ef3bc9bb5944056c63ccaa09.
//
// Solidity: event Repay(bytes32 indexed id, address indexed caller, address indexed onBehalf, uint256 assets, uint256 shares)
func (_Morpho *MorphoFilterer) ParseRepay(log types.Log) (*MorphoRepay, error) {
	event := new(MorphoRepay)
	if err := _Morpho.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
