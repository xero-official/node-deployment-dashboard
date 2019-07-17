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
const NodeMappingBin = `0x608060405234801561001057600080fd5b50604051602080610b128339016040526003805433600160a060020a03199182168117909255600480549091169091179055610ac1806100516000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631240abac8114610066578063ba7249c81461011f578063c9ba0b9a146101c6578063dbebfba61461026b575b600080fd5b34801561007257600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261010b958335600160a060020a031695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061028c9650505050505050565b604080519115158252519081900360200190f35b34801561012b57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101c4958335600160a060020a031695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506106bc9650505050505050565b005b3480156101d257600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101c4958335600160a060020a031695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506108169650505050505050565b34801561027757600080fd5b506101c4600160a060020a0360043516610956565b6000826040516020018082805190602001908083835b602083106102c15780518252601f1990920191602091820191016102a2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106103245780518252601f199092019160209182019101610305565b51815160209384036101000a6000190180199092169116179052604051919093018190038120885190955060009450889391925082918401908083835b602083106103805780518252601f199092019160209182019101610361565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060405160200180828054600181600116156101000203166002900480156104115780601f106103ef576101008083540402835291820191610411565b820191906000526020600020905b8154815290600101906020018083116103fd575b50509150506040516020818303038152906040526040518082805190602001908083835b602083106104545780518252601f199092019160209182019101610435565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614806104ab5750600160a060020a03808516600081815260016020526040902054909116145b806106aa5750816040516020018082805190602001908083835b602083106104e45780518252601f1990920191602091820191016104c5565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106105475780518252601f199092019160209182019101610528565b51815160209384036101000a6000190180199092169116179052604051919093018190038120875190955060029450879391925082918401908083835b602083106105a35780518252601f199092019160209182019101610584565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060405160200180828054600181600116156101000203166002900480156106345780601f10610612576101008083540402835291820191610634565b820191906000526020600020905b815481529060010190602001808311610620575b50509150506040516020818303038152906040526040518082805190602001908083835b602083106106775780518252601f199092019160209182019101610658565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916145b15156106b257fe5b5060019392505050565b600354600160a060020a03163314806106df5750600454600160a060020a031633145b15156106ea57600080fd5b816000836040518082805190602001908083835b6020831061071d5780518252601f1990920191602091820191016106fe565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101909320845161075e95919491909101925090506109b3565b50600160a060020a038316600081815260016020908152604091829020805473ffffffffffffffffffffffffffffffffffffffff19169093179092555182518392600292849290918291908401908083835b602083106107cf5780518252601f1990920191602091820191016107b0565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101909320845161081095919491909101925090506109b3565b50505050565b600354600160a060020a03163314806108395750600454600160a060020a031633145b151561084457600080fd5b6000826040518082805190602001908083835b602083106108765780518252601f199092019160209182019101610857565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206108af925090506000610a31565b600160a060020a038316600090815260016020908152604091829020805473ffffffffffffffffffffffffffffffffffffffff191690559051825160029284929182918401908083835b602083106109185780518252601f1990920191602091820191016108f9565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220610951925090506000610a31565b505050565b600354600160a060020a03163314806109795750600454600160a060020a031633145b151561098457600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109f457805160ff1916838001178555610a21565b82800160010185558215610a21579182015b82811115610a21578251825591602001919060010190610a06565b50610a2d929150610a78565b5090565b50805460018160011615610100020316600290046000825580601f10610a575750610a75565b601f016020900490600052602060002090810190610a759190610a78565b50565b610a9291905b80821115610a2d5760008155600101610a7e565b905600a165627a7a723058207aff8a5c23904ce956802f98c953c903d4cb39fb248fc133d0afa9d09781c0700029`

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
