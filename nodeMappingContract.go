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
const NodeMappingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"GetOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"CheckExistence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetNodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"AddNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"RemoveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetOperator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"SetOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"setCollateralRequirement\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NodeMappingBin is the compiled bytecode used for deploying new contracts.
const NodeMappingBin = `0x608060405234801561001057600080fd5b506040516020806107638339016040526003805433600160a060020a03199182168117909255600480549091169091179055610712806100516000396000f3006080604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ae50a3981146100875780631240abac146100b857806370f2a97114610171578063ba7249c814610198578063c9ba0b9a1461023f578063cededfb3146102e4578063dbebfba6146102f9575b600080fd5b34801561009357600080fd5b5061009c61031a565b60408051600160a060020a039092168252519081900360200190f35b3480156100c457600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261015d958335600160a060020a031695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506103299650505050505050565b604080519115158252519081900360200190f35b34801561017d57600080fd5b50610186610432565b60408051918252519081900360200190f35b3480156101a457600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261023d958335600160a060020a031695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506104389650505050505050565b005b34801561024b57600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261023d958335600160a060020a031695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061055d9650505050505050565b3480156102f057600080fd5b5061009c61067a565b34801561030557600080fd5b5061023d600160a060020a0360043516610689565b600354600160a060020a031690565b600080836040518082805190602001908083835b6020831061035c5780518252601f19909201916020918201910161033d565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546104571491508190506103b45750600160a060020a038416600090815260016020526040902054610457145b8061042057506002826040518082805190602001908083835b602083106103ec5780518252601f1990920191602091820191016103cd565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054610457149150505b151561042857fe5b5060019392505050565b60055490565b600354600160a060020a031633148061045b5750600454600160a060020a031633145b151561046657600080fd5b6104576000836040518082805190602001908083835b6020831061049b5780518252601f19909201916020918201910161047c565b51815160209384036101000a600019018019909216911617905292019485525060408051948590038201852095909555600160a060020a03881660009081526001825294909420610457908190558551909460029487945092508291908401908083835b6020831061051e5780518252601f1990920191602091820191016104ff565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929092555050600580546001019055505050565b600354600160a060020a03163314806105805750600454600160a060020a031633145b151561058b57600080fd5b6000826040518082805190602001908083835b602083106105bd5780518252601f19909201916020918201910161059e565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520600090819055600160a060020a038916815260018352908120558451600294869450925082918401908083835b602083106106355780518252601f199092019160209182019101610616565b51815160001960209485036101000a81019182169119929092161790915293909101958652604051958690030190942060009055600580549091019055505050505050565b600454600160a060020a031690565b600354600160a060020a03163314806106ac5750600454600160a060020a031633145b15156106b757600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820dabad8272b7e4ab40121ec55ad747808791c096c0d194423463c4e515bb9453f0029`

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

// GetNodeCount is a free data retrieval call binding the contract method 0x70f2a971.
//
// Solidity: function GetNodeCount() constant returns(uint256)
func (_NodeMapping *NodeMappingCaller) GetNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeMapping.contract.Call(opts, out, "GetNodeCount")
	return *ret0, err
}

// GetNodeCount is a free data retrieval call binding the contract method 0x70f2a971.
//
// Solidity: function GetNodeCount() constant returns(uint256)
func (_NodeMapping *NodeMappingSession) GetNodeCount() (*big.Int, error) {
	return _NodeMapping.Contract.GetNodeCount(&_NodeMapping.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x70f2a971.
//
// Solidity: function GetNodeCount() constant returns(uint256)
func (_NodeMapping *NodeMappingCallerSession) GetNodeCount() (*big.Int, error) {
	return _NodeMapping.Contract.GetNodeCount(&_NodeMapping.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xcededfb3.
//
// Solidity: function GetOperator() constant returns(address)
func (_NodeMapping *NodeMappingCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NodeMapping.contract.Call(opts, out, "GetOperator")
	return *ret0, err
}

// GetOperator is a free data retrieval call binding the contract method 0xcededfb3.
//
// Solidity: function GetOperator() constant returns(address)
func (_NodeMapping *NodeMappingSession) GetOperator() (common.Address, error) {
	return _NodeMapping.Contract.GetOperator(&_NodeMapping.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xcededfb3.
//
// Solidity: function GetOperator() constant returns(address)
func (_NodeMapping *NodeMappingCallerSession) GetOperator() (common.Address, error) {
	return _NodeMapping.Contract.GetOperator(&_NodeMapping.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ae50a39.
//
// Solidity: function GetOwner() constant returns(address)
func (_NodeMapping *NodeMappingCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NodeMapping.contract.Call(opts, out, "GetOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x0ae50a39.
//
// Solidity: function GetOwner() constant returns(address)
func (_NodeMapping *NodeMappingSession) GetOwner() (common.Address, error) {
	return _NodeMapping.Contract.GetOwner(&_NodeMapping.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ae50a39.
//
// Solidity: function GetOwner() constant returns(address)
func (_NodeMapping *NodeMappingCallerSession) GetOwner() (common.Address, error) {
	return _NodeMapping.Contract.GetOwner(&_NodeMapping.CallOpts)
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
