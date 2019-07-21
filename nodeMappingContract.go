// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NodeMappingABI is the input ABI used to generate the binding from.
const NodeMappingABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"CheckExistence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"AddNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"RemoveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"SetOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"setCollateralRequirement\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NodeMappingBin is the compiled bytecode used for deploying new contracts.
const NodeMappingBin = `0x608060405234801561001057600080fd5b506040516020806106998339016040526003805433600160a060020a03199182168117909255600480549091169091179055610648806100516000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631240abac8114610066578063ba7249c81461011f578063c9ba0b9a146101c6578063dbebfba61461026b575b600080fd5b34801561007257600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261010b958335600160a060020a031695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061028c9650505050505050565b604080519115158252519081900360200190f35b34801561012b57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101c4958335600160a060020a031695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506103959650505050505050565b005b3480156101d257600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101c4958335600160a060020a031695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506104b19650505050505050565b34801561027757600080fd5b506101c4600160a060020a03600435166105bf565b600080836040518082805190602001908083835b602083106102bf5780518252601f1990920191602091820191016102a0565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546104571491508190506103175750600160a060020a038416600090815260016020526040902054610457145b8061038357506002826040518082805190602001908083835b6020831061034f5780518252601f199092019160209182019101610330565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054610457149150505b151561038b57fe5b5060019392505050565b600354600160a060020a03163314806103b85750600454600160a060020a031633145b15156103c357600080fd5b6104576000836040518082805190602001908083835b602083106103f85780518252601f1990920191602091820191016103d9565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201852095909555600160a060020a03881660009081526001825294909420610457908190558551909460029487945092508291908401908083835b6020831061047b5780518252601f19909201916020918201910161045c565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929092555050505050565b600354600160a060020a03163314806104d45750600454600160a060020a031633145b15156104df57600080fd5b6000826040518082805190602001908083835b602083106105115780518252601f1990920191602091820191016104f2565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520600090819055600160a060020a038916815260018352908120558451600294869450925082918401908083835b602083106105895780518252601f19909201916020918201910161056a565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600090555050505050565b600354600160a060020a03163314806105e25750600454600160a060020a031633145b15156105ed57600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820685c6790bd8de0b7500327871d77156eec47bc550538b0d829beb8fc5e7d345e0029`

// DeployNodeMapping deploys a new Ethereum contract, binding an instance of NodeMapping to it.
func DeployNodeMapping(auth *bind.TransactOpts, backend bind.ContractBackend, setCollateralRequirement *big.Int) (common.Address, *types.Transaction, *NodeMapping, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeMappingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeMappingBin), backend, setCollateralRequirement)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeMapping{NodeMappingCaller: NodeMappingCaller{contract: contract}, NodeMappingTransactor: NodeMappingTransactor{contract: contract}, NodeMappingFilterer: NodeMappingFilterer{contract: contract}}, nil
}

// NodeMapping is an auto generated Go binding around an Ethereum contract.
type NodeMapping struct {
	NodeMappingCaller     // Read-only binding to the contract
	NodeMappingTransactor // Write-only binding to the contract
	NodeMappingFilterer   // Log filterer for contract events
}

// NodeMappingCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeMappingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeMappingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeMappingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeMappingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeMappingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeMappingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeMappingSession struct {
	Contract     *NodeMapping      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeMappingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeMappingCallerSession struct {
	Contract *NodeMappingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeMappingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeMappingTransactorSession struct {
	Contract     *NodeMappingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeMappingRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeMappingRaw struct {
	Contract *NodeMapping // Generic contract binding to access the raw methods on
}

// NodeMappingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeMappingCallerRaw struct {
	Contract *NodeMappingCaller // Generic read-only contract binding to access the raw methods on
}

// NodeMappingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeMappingTransactorRaw struct {
	Contract *NodeMappingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeMapping creates a new instance of NodeMapping, bound to a specific deployed contract.
func NewNodeMapping(address common.Address, backend bind.ContractBackend) (*NodeMapping, error) {
	contract, err := bindNodeMapping(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeMapping{NodeMappingCaller: NodeMappingCaller{contract: contract}, NodeMappingTransactor: NodeMappingTransactor{contract: contract}, NodeMappingFilterer: NodeMappingFilterer{contract: contract}}, nil
}

// NewNodeMappingCaller creates a new read-only instance of NodeMapping, bound to a specific deployed contract.
func NewNodeMappingCaller(address common.Address, caller bind.ContractCaller) (*NodeMappingCaller, error) {
	contract, err := bindNodeMapping(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeMappingCaller{contract: contract}, nil
}

// NewNodeMappingTransactor creates a new write-only instance of NodeMapping, bound to a specific deployed contract.
func NewNodeMappingTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeMappingTransactor, error) {
	contract, err := bindNodeMapping(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeMappingTransactor{contract: contract}, nil
}

// NewNodeMappingFilterer creates a new log filterer instance of NodeMapping, bound to a specific deployed contract.
func NewNodeMappingFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeMappingFilterer, error) {
	contract, err := bindNodeMapping(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeMappingFilterer{contract: contract}, nil
}

// bindNodeMapping binds a generic wrapper to an already deployed contract.
func bindNodeMapping(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeMappingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeMapping *NodeMappingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeMapping.Contract.NodeMappingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeMapping *NodeMappingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeMapping.Contract.NodeMappingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeMapping *NodeMappingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeMapping.Contract.NodeMappingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeMapping *NodeMappingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeMapping.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeMapping *NodeMappingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeMapping.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeMapping *NodeMappingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeMapping.Contract.contract.Transact(opts, method, params...)
}

// AddNode is a paid mutator transaction binding the contract method 0xba7249c8.
//
// Solidity: function AddNode(address nodeAddress, string id, string ip) returns()
func (_NodeMapping *NodeMappingTransactor) AddNode(opts *bind.TransactOpts, nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeMapping.contract.Transact(opts, "AddNode", nodeAddress, id, ip)
}

// AddNode is a paid mutator transaction binding the contract method 0xba7249c8.
//
// Solidity: function AddNode(address nodeAddress, string id, string ip) returns()
func (_NodeMapping *NodeMappingSession) AddNode(nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeMapping.Contract.AddNode(&_NodeMapping.TransactOpts, nodeAddress, id, ip)
}

// AddNode is a paid mutator transaction binding the contract method 0xba7249c8.
//
// Solidity: function AddNode(address nodeAddress, string id, string ip) returns()
func (_NodeMapping *NodeMappingTransactorSession) AddNode(nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeMapping.Contract.AddNode(&_NodeMapping.TransactOpts, nodeAddress, id, ip)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x1240abac.
//
// Solidity: function CheckExistence(address nodeAddress, string id, string ip) returns(bool)
func (_NodeMapping *NodeMappingTransactor) CheckExistence(opts *bind.TransactOpts, nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeMapping.contract.Transact(opts, "CheckExistence", nodeAddress, id, ip)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x1240abac.
//
// Solidity: function CheckExistence(address nodeAddress, string id, string ip) returns(bool)
func (_NodeMapping *NodeMappingSession) CheckExistence(nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeMapping.Contract.CheckExistence(&_NodeMapping.TransactOpts, nodeAddress, id, ip)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x1240abac.
//
// Solidity: function CheckExistence(address nodeAddress, string id, string ip) returns(bool)
func (_NodeMapping *NodeMappingTransactorSession) CheckExistence(nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeMapping.Contract.CheckExistence(&_NodeMapping.TransactOpts, nodeAddress, id, ip)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xc9ba0b9a.
//
// Solidity: function RemoveNode(address nodeAddress, string id, string ip) returns()
func (_NodeMapping *NodeMappingTransactor) RemoveNode(opts *bind.TransactOpts, nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeMapping.contract.Transact(opts, "RemoveNode", nodeAddress, id, ip)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xc9ba0b9a.
//
// Solidity: function RemoveNode(address nodeAddress, string id, string ip) returns()
func (_NodeMapping *NodeMappingSession) RemoveNode(nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeMapping.Contract.RemoveNode(&_NodeMapping.TransactOpts, nodeAddress, id, ip)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xc9ba0b9a.
//
// Solidity: function RemoveNode(address nodeAddress, string id, string ip) returns()
func (_NodeMapping *NodeMappingTransactorSession) RemoveNode(nodeAddress common.Address, id string, ip string) (*types.Transaction, error) {
	return _NodeMapping.Contract.RemoveNode(&_NodeMapping.TransactOpts, nodeAddress, id, ip)
}

// SetOperator is a paid mutator transaction binding the contract method 0xdbebfba6.
//
// Solidity: function SetOperator(address newOperator) returns()
func (_NodeMapping *NodeMappingTransactor) SetOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _NodeMapping.contract.Transact(opts, "SetOperator", newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xdbebfba6.
//
// Solidity: function SetOperator(address newOperator) returns()
func (_NodeMapping *NodeMappingSession) SetOperator(newOperator common.Address) (*types.Transaction, error) {
	return _NodeMapping.Contract.SetOperator(&_NodeMapping.TransactOpts, newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xdbebfba6.
//
// Solidity: function SetOperator(address newOperator) returns()
func (_NodeMapping *NodeMappingTransactorSession) SetOperator(newOperator common.Address) (*types.Transaction, error) {
	return _NodeMapping.Contract.SetOperator(&_NodeMapping.TransactOpts, newOperator)
}
