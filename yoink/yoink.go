// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yoink

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

// Score is an auto generated low-level Go binding around an user-defined struct.
type Score struct {
	Yoinks        *big.Int
	Time          *big.Int
	LastYoinkedAt *big.Int
}

// YoinkMetaData contains all meta data concerning the Yoink contract.
var YoinkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountBalanceOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArrayLengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeRemaining\",\"type\":\"uint256\"}],\"name\":\"SlowDown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC1155ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeHeld\",\"type\":\"uint256\"}],\"name\":\"Yoinked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COOLDOWN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLAG_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TROPHY_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"yoinks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastYoinkedAt\",\"type\":\"uint256\"}],\"internalType\":\"structScore\",\"name\":\"holderScore\",\"type\":\"tuple\"}],\"name\":\"generateFlagJSON\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"yoinks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastYoinkedAt\",\"type\":\"uint256\"}],\"internalType\":\"structScore\",\"name\":\"holderScore\",\"type\":\"tuple\"}],\"name\":\"generateFlagSVG\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"yoinks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastYoinkedAt\",\"type\":\"uint256\"}],\"internalType\":\"structScore\",\"name\":\"holderScore\",\"type\":\"tuple\"}],\"name\":\"generateScoreJSON\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"yoinks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastYoinkedAt\",\"type\":\"uint256\"}],\"internalType\":\"structScore\",\"name\":\"holderScore\",\"type\":\"tuple\"}],\"name\":\"generateScoreSVG\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"yoinks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastYoinkedAt\",\"type\":\"uint256\"}],\"internalType\":\"structScore\",\"name\":\"holderScore\",\"type\":\"tuple\"}],\"name\":\"generateTrophyJSON\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"yoinks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastYoinkedAt\",\"type\":\"uint256\"}],\"internalType\":\"structScore\",\"name\":\"holderScore\",\"type\":\"tuple\"}],\"name\":\"generateTrophySVG\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastYoinkedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastYoinkedBy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mostYoinks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"yoinker\",\"type\":\"address\"}],\"name\":\"score\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"yoinks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastYoinkedAt\",\"type\":\"uint256\"}],\"internalType\":\"structScore\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"topYoinker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalYoinks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yoink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YoinkABI is the input ABI used to generate the binding from.
// Deprecated: Use YoinkMetaData.ABI instead.
var YoinkABI = YoinkMetaData.ABI

// Yoink is an auto generated Go binding around an Ethereum contract.
type Yoink struct {
	YoinkCaller     // Read-only binding to the contract
	YoinkTransactor // Write-only binding to the contract
	YoinkFilterer   // Log filterer for contract events
}

// YoinkCaller is an auto generated read-only Go binding around an Ethereum contract.
type YoinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YoinkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YoinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YoinkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YoinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YoinkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YoinkSession struct {
	Contract     *Yoink            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YoinkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YoinkCallerSession struct {
	Contract *YoinkCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// YoinkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YoinkTransactorSession struct {
	Contract     *YoinkTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YoinkRaw is an auto generated low-level Go binding around an Ethereum contract.
type YoinkRaw struct {
	Contract *Yoink // Generic contract binding to access the raw methods on
}

// YoinkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YoinkCallerRaw struct {
	Contract *YoinkCaller // Generic read-only contract binding to access the raw methods on
}

// YoinkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YoinkTransactorRaw struct {
	Contract *YoinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYoink creates a new instance of Yoink, bound to a specific deployed contract.
func NewYoink(address common.Address, backend bind.ContractBackend) (*Yoink, error) {
	contract, err := bindYoink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yoink{YoinkCaller: YoinkCaller{contract: contract}, YoinkTransactor: YoinkTransactor{contract: contract}, YoinkFilterer: YoinkFilterer{contract: contract}}, nil
}

// NewYoinkCaller creates a new read-only instance of Yoink, bound to a specific deployed contract.
func NewYoinkCaller(address common.Address, caller bind.ContractCaller) (*YoinkCaller, error) {
	contract, err := bindYoink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YoinkCaller{contract: contract}, nil
}

// NewYoinkTransactor creates a new write-only instance of Yoink, bound to a specific deployed contract.
func NewYoinkTransactor(address common.Address, transactor bind.ContractTransactor) (*YoinkTransactor, error) {
	contract, err := bindYoink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YoinkTransactor{contract: contract}, nil
}

// NewYoinkFilterer creates a new log filterer instance of Yoink, bound to a specific deployed contract.
func NewYoinkFilterer(address common.Address, filterer bind.ContractFilterer) (*YoinkFilterer, error) {
	contract, err := bindYoink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YoinkFilterer{contract: contract}, nil
}

// bindYoink binds a generic wrapper to an already deployed contract.
func bindYoink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YoinkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yoink *YoinkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yoink.Contract.YoinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yoink *YoinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yoink.Contract.YoinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yoink *YoinkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yoink.Contract.YoinkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yoink *YoinkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yoink.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yoink *YoinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yoink.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yoink *YoinkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yoink.Contract.contract.Transact(opts, method, params...)
}

// COOLDOWN is a free data retrieval call binding the contract method 0xa2724a4d.
//
// Solidity: function COOLDOWN() view returns(uint256)
func (_Yoink *YoinkCaller) COOLDOWN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "COOLDOWN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COOLDOWN is a free data retrieval call binding the contract method 0xa2724a4d.
//
// Solidity: function COOLDOWN() view returns(uint256)
func (_Yoink *YoinkSession) COOLDOWN() (*big.Int, error) {
	return _Yoink.Contract.COOLDOWN(&_Yoink.CallOpts)
}

// COOLDOWN is a free data retrieval call binding the contract method 0xa2724a4d.
//
// Solidity: function COOLDOWN() view returns(uint256)
func (_Yoink *YoinkCallerSession) COOLDOWN() (*big.Int, error) {
	return _Yoink.Contract.COOLDOWN(&_Yoink.CallOpts)
}

// FLAGID is a free data retrieval call binding the contract method 0x091019fc.
//
// Solidity: function FLAG_ID() view returns(uint256)
func (_Yoink *YoinkCaller) FLAGID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "FLAG_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGID is a free data retrieval call binding the contract method 0x091019fc.
//
// Solidity: function FLAG_ID() view returns(uint256)
func (_Yoink *YoinkSession) FLAGID() (*big.Int, error) {
	return _Yoink.Contract.FLAGID(&_Yoink.CallOpts)
}

// FLAGID is a free data retrieval call binding the contract method 0x091019fc.
//
// Solidity: function FLAG_ID() view returns(uint256)
func (_Yoink *YoinkCallerSession) FLAGID() (*big.Int, error) {
	return _Yoink.Contract.FLAGID(&_Yoink.CallOpts)
}

// TROPHYID is a free data retrieval call binding the contract method 0x4536de9d.
//
// Solidity: function TROPHY_ID() view returns(uint256)
func (_Yoink *YoinkCaller) TROPHYID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "TROPHY_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TROPHYID is a free data retrieval call binding the contract method 0x4536de9d.
//
// Solidity: function TROPHY_ID() view returns(uint256)
func (_Yoink *YoinkSession) TROPHYID() (*big.Int, error) {
	return _Yoink.Contract.TROPHYID(&_Yoink.CallOpts)
}

// TROPHYID is a free data retrieval call binding the contract method 0x4536de9d.
//
// Solidity: function TROPHY_ID() view returns(uint256)
func (_Yoink *YoinkCallerSession) TROPHYID() (*big.Int, error) {
	return _Yoink.Contract.TROPHYID(&_Yoink.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256 result)
func (_Yoink *YoinkCaller) BalanceOf(opts *bind.CallOpts, owner common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "balanceOf", owner, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256 result)
func (_Yoink *YoinkSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _Yoink.Contract.BalanceOf(&_Yoink.CallOpts, owner, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256 result)
func (_Yoink *YoinkCallerSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _Yoink.Contract.BalanceOf(&_Yoink.CallOpts, owner, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_Yoink *YoinkCaller) BalanceOfBatch(opts *bind.CallOpts, owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "balanceOfBatch", owners, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_Yoink *YoinkSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Yoink.Contract.BalanceOfBatch(&_Yoink.CallOpts, owners, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_Yoink *YoinkCallerSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Yoink.Contract.BalanceOfBatch(&_Yoink.CallOpts, owners, ids)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() pure returns(string)
func (_Yoink *YoinkCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() pure returns(string)
func (_Yoink *YoinkSession) ContractURI() (string, error) {
	return _Yoink.Contract.ContractURI(&_Yoink.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() pure returns(string)
func (_Yoink *YoinkCallerSession) ContractURI() (string, error) {
	return _Yoink.Contract.ContractURI(&_Yoink.CallOpts)
}

// GenerateFlagJSON is a free data retrieval call binding the contract method 0xfdeba773.
//
// Solidity: function generateFlagJSON(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCaller) GenerateFlagJSON(opts *bind.CallOpts, holder common.Address, holderScore Score) (string, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "generateFlagJSON", holder, holderScore)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GenerateFlagJSON is a free data retrieval call binding the contract method 0xfdeba773.
//
// Solidity: function generateFlagJSON(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkSession) GenerateFlagJSON(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateFlagJSON(&_Yoink.CallOpts, holder, holderScore)
}

// GenerateFlagJSON is a free data retrieval call binding the contract method 0xfdeba773.
//
// Solidity: function generateFlagJSON(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCallerSession) GenerateFlagJSON(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateFlagJSON(&_Yoink.CallOpts, holder, holderScore)
}

// GenerateFlagSVG is a free data retrieval call binding the contract method 0x382e8562.
//
// Solidity: function generateFlagSVG(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCaller) GenerateFlagSVG(opts *bind.CallOpts, holder common.Address, holderScore Score) (string, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "generateFlagSVG", holder, holderScore)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GenerateFlagSVG is a free data retrieval call binding the contract method 0x382e8562.
//
// Solidity: function generateFlagSVG(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkSession) GenerateFlagSVG(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateFlagSVG(&_Yoink.CallOpts, holder, holderScore)
}

// GenerateFlagSVG is a free data retrieval call binding the contract method 0x382e8562.
//
// Solidity: function generateFlagSVG(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCallerSession) GenerateFlagSVG(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateFlagSVG(&_Yoink.CallOpts, holder, holderScore)
}

// GenerateScoreJSON is a free data retrieval call binding the contract method 0xf33b79dd.
//
// Solidity: function generateScoreJSON(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCaller) GenerateScoreJSON(opts *bind.CallOpts, holder common.Address, holderScore Score) (string, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "generateScoreJSON", holder, holderScore)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GenerateScoreJSON is a free data retrieval call binding the contract method 0xf33b79dd.
//
// Solidity: function generateScoreJSON(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkSession) GenerateScoreJSON(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateScoreJSON(&_Yoink.CallOpts, holder, holderScore)
}

// GenerateScoreJSON is a free data retrieval call binding the contract method 0xf33b79dd.
//
// Solidity: function generateScoreJSON(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCallerSession) GenerateScoreJSON(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateScoreJSON(&_Yoink.CallOpts, holder, holderScore)
}

// GenerateScoreSVG is a free data retrieval call binding the contract method 0x67ee08e9.
//
// Solidity: function generateScoreSVG(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCaller) GenerateScoreSVG(opts *bind.CallOpts, holder common.Address, holderScore Score) (string, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "generateScoreSVG", holder, holderScore)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GenerateScoreSVG is a free data retrieval call binding the contract method 0x67ee08e9.
//
// Solidity: function generateScoreSVG(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkSession) GenerateScoreSVG(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateScoreSVG(&_Yoink.CallOpts, holder, holderScore)
}

// GenerateScoreSVG is a free data retrieval call binding the contract method 0x67ee08e9.
//
// Solidity: function generateScoreSVG(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCallerSession) GenerateScoreSVG(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateScoreSVG(&_Yoink.CallOpts, holder, holderScore)
}

// GenerateTrophyJSON is a free data retrieval call binding the contract method 0x4b3ab9be.
//
// Solidity: function generateTrophyJSON(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCaller) GenerateTrophyJSON(opts *bind.CallOpts, holder common.Address, holderScore Score) (string, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "generateTrophyJSON", holder, holderScore)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GenerateTrophyJSON is a free data retrieval call binding the contract method 0x4b3ab9be.
//
// Solidity: function generateTrophyJSON(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkSession) GenerateTrophyJSON(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateTrophyJSON(&_Yoink.CallOpts, holder, holderScore)
}

// GenerateTrophyJSON is a free data retrieval call binding the contract method 0x4b3ab9be.
//
// Solidity: function generateTrophyJSON(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCallerSession) GenerateTrophyJSON(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateTrophyJSON(&_Yoink.CallOpts, holder, holderScore)
}

// GenerateTrophySVG is a free data retrieval call binding the contract method 0x631cfbf9.
//
// Solidity: function generateTrophySVG(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCaller) GenerateTrophySVG(opts *bind.CallOpts, holder common.Address, holderScore Score) (string, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "generateTrophySVG", holder, holderScore)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GenerateTrophySVG is a free data retrieval call binding the contract method 0x631cfbf9.
//
// Solidity: function generateTrophySVG(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkSession) GenerateTrophySVG(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateTrophySVG(&_Yoink.CallOpts, holder, holderScore)
}

// GenerateTrophySVG is a free data retrieval call binding the contract method 0x631cfbf9.
//
// Solidity: function generateTrophySVG(address holder, (uint256,uint256,uint256) holderScore) pure returns(string)
func (_Yoink *YoinkCallerSession) GenerateTrophySVG(holder common.Address, holderScore Score) (string, error) {
	return _Yoink.Contract.GenerateTrophySVG(&_Yoink.CallOpts, holder, holderScore)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool result)
func (_Yoink *YoinkCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool result)
func (_Yoink *YoinkSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Yoink.Contract.IsApprovedForAll(&_Yoink.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool result)
func (_Yoink *YoinkCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Yoink.Contract.IsApprovedForAll(&_Yoink.CallOpts, owner, operator)
}

// LastYoinkedAt is a free data retrieval call binding the contract method 0x6a99616f.
//
// Solidity: function lastYoinkedAt() view returns(uint256)
func (_Yoink *YoinkCaller) LastYoinkedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "lastYoinkedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastYoinkedAt is a free data retrieval call binding the contract method 0x6a99616f.
//
// Solidity: function lastYoinkedAt() view returns(uint256)
func (_Yoink *YoinkSession) LastYoinkedAt() (*big.Int, error) {
	return _Yoink.Contract.LastYoinkedAt(&_Yoink.CallOpts)
}

// LastYoinkedAt is a free data retrieval call binding the contract method 0x6a99616f.
//
// Solidity: function lastYoinkedAt() view returns(uint256)
func (_Yoink *YoinkCallerSession) LastYoinkedAt() (*big.Int, error) {
	return _Yoink.Contract.LastYoinkedAt(&_Yoink.CallOpts)
}

// LastYoinkedBy is a free data retrieval call binding the contract method 0xd4dbf9f4.
//
// Solidity: function lastYoinkedBy() view returns(address)
func (_Yoink *YoinkCaller) LastYoinkedBy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "lastYoinkedBy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastYoinkedBy is a free data retrieval call binding the contract method 0xd4dbf9f4.
//
// Solidity: function lastYoinkedBy() view returns(address)
func (_Yoink *YoinkSession) LastYoinkedBy() (common.Address, error) {
	return _Yoink.Contract.LastYoinkedBy(&_Yoink.CallOpts)
}

// LastYoinkedBy is a free data retrieval call binding the contract method 0xd4dbf9f4.
//
// Solidity: function lastYoinkedBy() view returns(address)
func (_Yoink *YoinkCallerSession) LastYoinkedBy() (common.Address, error) {
	return _Yoink.Contract.LastYoinkedBy(&_Yoink.CallOpts)
}

// MostYoinks is a free data retrieval call binding the contract method 0xd2d7774a.
//
// Solidity: function mostYoinks() view returns(uint256)
func (_Yoink *YoinkCaller) MostYoinks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "mostYoinks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MostYoinks is a free data retrieval call binding the contract method 0xd2d7774a.
//
// Solidity: function mostYoinks() view returns(uint256)
func (_Yoink *YoinkSession) MostYoinks() (*big.Int, error) {
	return _Yoink.Contract.MostYoinks(&_Yoink.CallOpts)
}

// MostYoinks is a free data retrieval call binding the contract method 0xd2d7774a.
//
// Solidity: function mostYoinks() view returns(uint256)
func (_Yoink *YoinkCallerSession) MostYoinks() (*big.Int, error) {
	return _Yoink.Contract.MostYoinks(&_Yoink.CallOpts)
}

// Score is a free data retrieval call binding the contract method 0x776f3843.
//
// Solidity: function score(address yoinker) view returns((uint256,uint256,uint256))
func (_Yoink *YoinkCaller) Score(opts *bind.CallOpts, yoinker common.Address) (Score, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "score", yoinker)

	if err != nil {
		return *new(Score), err
	}

	out0 := *abi.ConvertType(out[0], new(Score)).(*Score)

	return out0, err

}

// Score is a free data retrieval call binding the contract method 0x776f3843.
//
// Solidity: function score(address yoinker) view returns((uint256,uint256,uint256))
func (_Yoink *YoinkSession) Score(yoinker common.Address) (Score, error) {
	return _Yoink.Contract.Score(&_Yoink.CallOpts, yoinker)
}

// Score is a free data retrieval call binding the contract method 0x776f3843.
//
// Solidity: function score(address yoinker) view returns((uint256,uint256,uint256))
func (_Yoink *YoinkCallerSession) Score(yoinker common.Address) (Score, error) {
	return _Yoink.Contract.Score(&_Yoink.CallOpts, yoinker)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool result)
func (_Yoink *YoinkCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool result)
func (_Yoink *YoinkSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Yoink.Contract.SupportsInterface(&_Yoink.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool result)
func (_Yoink *YoinkCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Yoink.Contract.SupportsInterface(&_Yoink.CallOpts, interfaceId)
}

// TopYoinker is a free data retrieval call binding the contract method 0x6a974e6e.
//
// Solidity: function topYoinker() view returns(address)
func (_Yoink *YoinkCaller) TopYoinker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "topYoinker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TopYoinker is a free data retrieval call binding the contract method 0x6a974e6e.
//
// Solidity: function topYoinker() view returns(address)
func (_Yoink *YoinkSession) TopYoinker() (common.Address, error) {
	return _Yoink.Contract.TopYoinker(&_Yoink.CallOpts)
}

// TopYoinker is a free data retrieval call binding the contract method 0x6a974e6e.
//
// Solidity: function topYoinker() view returns(address)
func (_Yoink *YoinkCallerSession) TopYoinker() (common.Address, error) {
	return _Yoink.Contract.TopYoinker(&_Yoink.CallOpts)
}

// TotalYoinks is a free data retrieval call binding the contract method 0xa5d0dadd.
//
// Solidity: function totalYoinks() view returns(uint256)
func (_Yoink *YoinkCaller) TotalYoinks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "totalYoinks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalYoinks is a free data retrieval call binding the contract method 0xa5d0dadd.
//
// Solidity: function totalYoinks() view returns(uint256)
func (_Yoink *YoinkSession) TotalYoinks() (*big.Int, error) {
	return _Yoink.Contract.TotalYoinks(&_Yoink.CallOpts)
}

// TotalYoinks is a free data retrieval call binding the contract method 0xa5d0dadd.
//
// Solidity: function totalYoinks() view returns(uint256)
func (_Yoink *YoinkCallerSession) TotalYoinks() (*big.Int, error) {
	return _Yoink.Contract.TotalYoinks(&_Yoink.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Yoink *YoinkCaller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Yoink.contract.Call(opts, &out, "uri", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Yoink *YoinkSession) Uri(id *big.Int) (string, error) {
	return _Yoink.Contract.Uri(&_Yoink.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Yoink *YoinkCallerSession) Uri(id *big.Int) (string, error) {
	return _Yoink.Contract.Uri(&_Yoink.CallOpts, id)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Yoink *YoinkTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Yoink.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Yoink *YoinkSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Yoink.Contract.SafeBatchTransferFrom(&_Yoink.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Yoink *YoinkTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Yoink.Contract.SafeBatchTransferFrom(&_Yoink.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Yoink *YoinkTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Yoink.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Yoink *YoinkSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Yoink.Contract.SafeTransferFrom(&_Yoink.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Yoink *YoinkTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Yoink.Contract.SafeTransferFrom(&_Yoink.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool isApproved) returns()
func (_Yoink *YoinkTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, isApproved bool) (*types.Transaction, error) {
	return _Yoink.contract.Transact(opts, "setApprovalForAll", operator, isApproved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool isApproved) returns()
func (_Yoink *YoinkSession) SetApprovalForAll(operator common.Address, isApproved bool) (*types.Transaction, error) {
	return _Yoink.Contract.SetApprovalForAll(&_Yoink.TransactOpts, operator, isApproved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool isApproved) returns()
func (_Yoink *YoinkTransactorSession) SetApprovalForAll(operator common.Address, isApproved bool) (*types.Transaction, error) {
	return _Yoink.Contract.SetApprovalForAll(&_Yoink.TransactOpts, operator, isApproved)
}

// Yoink is a paid mutator transaction binding the contract method 0x9846cd9e.
//
// Solidity: function yoink() returns()
func (_Yoink *YoinkTransactor) Yoink(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yoink.contract.Transact(opts, "yoink")
}

// Yoink is a paid mutator transaction binding the contract method 0x9846cd9e.
//
// Solidity: function yoink() returns()
func (_Yoink *YoinkSession) Yoink() (*types.Transaction, error) {
	return _Yoink.Contract.Yoink(&_Yoink.TransactOpts)
}

// Yoink is a paid mutator transaction binding the contract method 0x9846cd9e.
//
// Solidity: function yoink() returns()
func (_Yoink *YoinkTransactorSession) Yoink() (*types.Transaction, error) {
	return _Yoink.Contract.Yoink(&_Yoink.TransactOpts)
}

// YoinkApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Yoink contract.
type YoinkApprovalForAllIterator struct {
	Event *YoinkApprovalForAll // Event containing the contract specifics and raw log

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
func (it *YoinkApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YoinkApprovalForAll)
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
		it.Event = new(YoinkApprovalForAll)
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
func (it *YoinkApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YoinkApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YoinkApprovalForAll represents a ApprovalForAll event raised by the Yoink contract.
type YoinkApprovalForAll struct {
	Owner      common.Address
	Operator   common.Address
	IsApproved bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool isApproved)
func (_Yoink *YoinkFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*YoinkApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Yoink.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &YoinkApprovalForAllIterator{contract: _Yoink.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool isApproved)
func (_Yoink *YoinkFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *YoinkApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Yoink.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YoinkApprovalForAll)
				if err := _Yoink.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool isApproved)
func (_Yoink *YoinkFilterer) ParseApprovalForAll(log types.Log) (*YoinkApprovalForAll, error) {
	event := new(YoinkApprovalForAll)
	if err := _Yoink.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YoinkTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Yoink contract.
type YoinkTransferBatchIterator struct {
	Event *YoinkTransferBatch // Event containing the contract specifics and raw log

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
func (it *YoinkTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YoinkTransferBatch)
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
		it.Event = new(YoinkTransferBatch)
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
func (it *YoinkTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YoinkTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YoinkTransferBatch represents a TransferBatch event raised by the Yoink contract.
type YoinkTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Yoink *YoinkFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*YoinkTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Yoink.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YoinkTransferBatchIterator{contract: _Yoink.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Yoink *YoinkFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *YoinkTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Yoink.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YoinkTransferBatch)
				if err := _Yoink.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Yoink *YoinkFilterer) ParseTransferBatch(log types.Log) (*YoinkTransferBatch, error) {
	event := new(YoinkTransferBatch)
	if err := _Yoink.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YoinkTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Yoink contract.
type YoinkTransferSingleIterator struct {
	Event *YoinkTransferSingle // Event containing the contract specifics and raw log

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
func (it *YoinkTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YoinkTransferSingle)
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
		it.Event = new(YoinkTransferSingle)
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
func (it *YoinkTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YoinkTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YoinkTransferSingle represents a TransferSingle event raised by the Yoink contract.
type YoinkTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Yoink *YoinkFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*YoinkTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Yoink.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YoinkTransferSingleIterator{contract: _Yoink.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Yoink *YoinkFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *YoinkTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Yoink.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YoinkTransferSingle)
				if err := _Yoink.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Yoink *YoinkFilterer) ParseTransferSingle(log types.Log) (*YoinkTransferSingle, error) {
	event := new(YoinkTransferSingle)
	if err := _Yoink.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YoinkURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Yoink contract.
type YoinkURIIterator struct {
	Event *YoinkURI // Event containing the contract specifics and raw log

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
func (it *YoinkURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YoinkURI)
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
		it.Event = new(YoinkURI)
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
func (it *YoinkURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YoinkURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YoinkURI represents a URI event raised by the Yoink contract.
type YoinkURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Yoink *YoinkFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*YoinkURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Yoink.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &YoinkURIIterator{contract: _Yoink.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Yoink *YoinkFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *YoinkURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Yoink.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YoinkURI)
				if err := _Yoink.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Yoink *YoinkFilterer) ParseURI(log types.Log) (*YoinkURI, error) {
	event := new(YoinkURI)
	if err := _Yoink.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YoinkYoinkedIterator is returned from FilterYoinked and is used to iterate over the raw logs and unpacked data for Yoinked events raised by the Yoink contract.
type YoinkYoinkedIterator struct {
	Event *YoinkYoinked // Event containing the contract specifics and raw log

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
func (it *YoinkYoinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YoinkYoinked)
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
		it.Event = new(YoinkYoinked)
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
func (it *YoinkYoinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YoinkYoinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YoinkYoinked represents a Yoinked event raised by the Yoink contract.
type YoinkYoinked struct {
	By       common.Address
	From     common.Address
	TimeHeld *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterYoinked is a free log retrieval operation binding the contract event 0x0cf9f4581bed778e0b60c8ae447949ed2751b37d702da26074bbdcd4a1c7a85c.
//
// Solidity: event Yoinked(address indexed by, address indexed from, uint256 timeHeld)
func (_Yoink *YoinkFilterer) FilterYoinked(opts *bind.FilterOpts, by []common.Address, from []common.Address) (*YoinkYoinkedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Yoink.contract.FilterLogs(opts, "Yoinked", byRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &YoinkYoinkedIterator{contract: _Yoink.contract, event: "Yoinked", logs: logs, sub: sub}, nil
}

// WatchYoinked is a free log subscription operation binding the contract event 0x0cf9f4581bed778e0b60c8ae447949ed2751b37d702da26074bbdcd4a1c7a85c.
//
// Solidity: event Yoinked(address indexed by, address indexed from, uint256 timeHeld)
func (_Yoink *YoinkFilterer) WatchYoinked(opts *bind.WatchOpts, sink chan<- *YoinkYoinked, by []common.Address, from []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Yoink.contract.WatchLogs(opts, "Yoinked", byRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YoinkYoinked)
				if err := _Yoink.contract.UnpackLog(event, "Yoinked", log); err != nil {
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

// ParseYoinked is a log parse operation binding the contract event 0x0cf9f4581bed778e0b60c8ae447949ed2751b37d702da26074bbdcd4a1c7a85c.
//
// Solidity: event Yoinked(address indexed by, address indexed from, uint256 timeHeld)
func (_Yoink *YoinkFilterer) ParseYoinked(log types.Log) (*YoinkYoinked, error) {
	event := new(YoinkYoinked)
	if err := _Yoink.contract.UnpackLog(event, "Yoinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
