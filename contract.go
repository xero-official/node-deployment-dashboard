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

// NodeContractABI is the input ABI used to generate the binding from.
const NodeContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"GetNodePort\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"GetNodeId\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"RemoveNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"GetNodeIp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"ip\",\"type\":\"string\"},{\"name\":\"port\",\"type\":\"string\"}],\"name\":\"AddNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"setCollateralRequirement\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NodeContractBin is the compiled bytecode used for deploying new contracts.
const NodeContractBin = `0x608060405234801561001057600080fd5b506040516020806107438339810160405251600060025560045560038054600160a060020a031916331790556106f88061004b6000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631cb3b230811461007157806326651d4c146101075780632a39f0d0146101285780638ee1d7ab1461015157806392be66c014610172575b600080fd5b34801561007d57600080fd5b50610092600160a060020a036004351661023a565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100cc5781810151838201526020016100b4565b50505050905090810190601f1680156100f95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561011357600080fd5b50610092600160a060020a03600435166102e6565b34801561013457600080fd5b5061013d61035e565b604080519115158252519081900360200190f35b34801561015d57600080fd5b50610092600160a060020a036004351661049d565b6040805160206004803580820135601f810184900484028501840190955284845261013d94369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506105129650505050505050565b600160a060020a0381166000908152602081815260409182902060040180548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156102da5780601f106102af576101008083540402835291602001916102da565b820191906000526020600020905b8154815290600101906020018083116102bd57829003601f168201915b50505050509050919050565b600160a060020a038116600090815260208181526040918290206002908101805484516001821615610100026000190190911692909204601f810184900484028301840190945283825260609391929091908301828280156102da5780601f106102af576101008083540402835291602001916102da565b600454336000908152602081905260408120600501549091670de0b6b3a7640000021461038757fe5b3360008181526020819052604080822060050154905181156108fc0292818181858888f193505050501580156103c1573d6000803e3d6000fd5b50336000818152602081815260408083208054600280546000199081018752600180875285882054600160a060020a0390811689528888528689208590558354830189528188528689205494895286892080549590911673ffffffffffffffffffffffffffffffffffffffff199586161790558254909101875293862080548316905595855292849052838155908101805490921690915591610466908301826105ed565b6104746003830160006105ed565b6104826004830160006105ed565b50600060059190910155506002805460001901905560015b90565b600160a060020a0381166000908152602081815260409182902060030180548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156102da5780601f106102af576101008083540402835291602001916102da565b600080600454670de0b6b3a7640000023414151561052c57fe5b503360008181526020818152604090912060018101805473ffffffffffffffffffffffffffffffffffffffff1916909317909255346005830155855161057a91600284019190880190610634565b5083516105909060038301906020870190610634565b5082516105a69060048301906020860190610634565b50600280548083556000908152600160208190526040909120805473ffffffffffffffffffffffffffffffffffffffff191633179055815481019091559150509392505050565b50805460018160011615610100020316600290046000825580601f106106135750610631565b601f01602090049060005260206000209081019061063191906106b2565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061067557805160ff19168380011785556106a2565b828001600101855582156106a2579182015b828111156106a2578251825591602001919060010190610687565b506106ae9291506106b2565b5090565b61049a91905b808211156106ae57600081556001016106b85600a165627a7a723058202a8101b172804df3376d7a464da769b90a6f2601c0ec45403c6da633e8b412570029`

// DeployNodeContract deploys a new Ethereum contract, binding an instance of NodeContract to it.
func DeployNodeContract(auth *bind.TransactOpts, backend bind.ContractBackend, setCollateralRequirement *big.Int) (common.Address, *types.Transaction, *NodeContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeContractBin), backend, setCollateralRequirement)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeContract{NodeContractCaller: NodeContractCaller{contract: contract}, NodeContractTransactor: NodeContractTransactor{contract: contract}, NodeContractFilterer: NodeContractFilterer{contract: contract}}, nil
}

// NodeContract is an auto generated Go binding around an Ethereum contract.
type NodeContract struct {
	NodeContractCaller     // Read-only binding to the contract
	NodeContractTransactor // Write-only binding to the contract
	NodeContractFilterer   // Log filterer for contract events
}

// NodeContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeContractSession struct {
	Contract     *NodeContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeContractCallerSession struct {
	Contract *NodeContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NodeContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeContractTransactorSession struct {
	Contract     *NodeContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NodeContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeContractRaw struct {
	Contract *NodeContract // Generic contract binding to access the raw methods on
}

// NodeContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeContractCallerRaw struct {
	Contract *NodeContractCaller // Generic read-only contract binding to access the raw methods on
}

// NodeContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeContractTransactorRaw struct {
	Contract *NodeContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeContract creates a new instance of NodeContract, bound to a specific deployed contract.
func NewNodeContract(address common.Address, backend bind.ContractBackend) (*NodeContract, error) {
	contract, err := bindNodeContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeContract{NodeContractCaller: NodeContractCaller{contract: contract}, NodeContractTransactor: NodeContractTransactor{contract: contract}, NodeContractFilterer: NodeContractFilterer{contract: contract}}, nil
}

// NewNodeContractCaller creates a new read-only instance of NodeContract, bound to a specific deployed contract.
func NewNodeContractCaller(address common.Address, caller bind.ContractCaller) (*NodeContractCaller, error) {
	contract, err := bindNodeContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeContractCaller{contract: contract}, nil
}

// NewNodeContractTransactor creates a new write-only instance of NodeContract, bound to a specific deployed contract.
func NewNodeContractTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeContractTransactor, error) {
	contract, err := bindNodeContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeContractTransactor{contract: contract}, nil
}

// NewNodeContractFilterer creates a new log filterer instance of NodeContract, bound to a specific deployed contract.
func NewNodeContractFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeContractFilterer, error) {
	contract, err := bindNodeContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeContractFilterer{contract: contract}, nil
}

// bindNodeContract binds a generic wrapper to an already deployed contract.
func bindNodeContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeContract *NodeContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeContract.Contract.NodeContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeContract *NodeContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeContract.Contract.NodeContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeContract *NodeContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeContract.Contract.NodeContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeContract *NodeContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeContract *NodeContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeContract *NodeContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeContract.Contract.contract.Transact(opts, method, params...)
}

// GetNodeId is a free data retrieval call binding the contract method 0x26651d4c.
//
// Solidity: function GetNodeId(address nodeAddress) constant returns(string)
func (_NodeContract *NodeContractCaller) GetNodeId(opts *bind.CallOpts, nodeAddress common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NodeContract.contract.Call(opts, out, "GetNodeId", nodeAddress)
	return *ret0, err
}

// GetNodeId is a free data retrieval call binding the contract method 0x26651d4c.
//
// Solidity: function GetNodeId(address nodeAddress) constant returns(string)
func (_NodeContract *NodeContractSession) GetNodeId(nodeAddress common.Address) (string, error) {
	return _NodeContract.Contract.GetNodeId(&_NodeContract.CallOpts, nodeAddress)
}

// GetNodeId is a free data retrieval call binding the contract method 0x26651d4c.
//
// Solidity: function GetNodeId(address nodeAddress) constant returns(string)
func (_NodeContract *NodeContractCallerSession) GetNodeId(nodeAddress common.Address) (string, error) {
	return _NodeContract.Contract.GetNodeId(&_NodeContract.CallOpts, nodeAddress)
}

// GetNodeIp is a free data retrieval call binding the contract method 0x8ee1d7ab.
//
// Solidity: function GetNodeIp(address nodeAddress) constant returns(string)
func (_NodeContract *NodeContractCaller) GetNodeIp(opts *bind.CallOpts, nodeAddress common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NodeContract.contract.Call(opts, out, "GetNodeIp", nodeAddress)
	return *ret0, err
}

// GetNodeIp is a free data retrieval call binding the contract method 0x8ee1d7ab.
//
// Solidity: function GetNodeIp(address nodeAddress) constant returns(string)
func (_NodeContract *NodeContractSession) GetNodeIp(nodeAddress common.Address) (string, error) {
	return _NodeContract.Contract.GetNodeIp(&_NodeContract.CallOpts, nodeAddress)
}

// GetNodeIp is a free data retrieval call binding the contract method 0x8ee1d7ab.
//
// Solidity: function GetNodeIp(address nodeAddress) constant returns(string)
func (_NodeContract *NodeContractCallerSession) GetNodeIp(nodeAddress common.Address) (string, error) {
	return _NodeContract.Contract.GetNodeIp(&_NodeContract.CallOpts, nodeAddress)
}

// GetNodePort is a free data retrieval call binding the contract method 0x1cb3b230.
//
// Solidity: function GetNodePort(address nodeAddress) constant returns(string)
func (_NodeContract *NodeContractCaller) GetNodePort(opts *bind.CallOpts, nodeAddress common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NodeContract.contract.Call(opts, out, "GetNodePort", nodeAddress)
	return *ret0, err
}

// GetNodePort is a free data retrieval call binding the contract method 0x1cb3b230.
//
// Solidity: function GetNodePort(address nodeAddress) constant returns(string)
func (_NodeContract *NodeContractSession) GetNodePort(nodeAddress common.Address) (string, error) {
	return _NodeContract.Contract.GetNodePort(&_NodeContract.CallOpts, nodeAddress)
}

// GetNodePort is a free data retrieval call binding the contract method 0x1cb3b230.
//
// Solidity: function GetNodePort(address nodeAddress) constant returns(string)
func (_NodeContract *NodeContractCallerSession) GetNodePort(nodeAddress common.Address) (string, error) {
	return _NodeContract.Contract.GetNodePort(&_NodeContract.CallOpts, nodeAddress)
}

// AddNode is a paid mutator transaction binding the contract method 0x92be66c0.
//
// Solidity: function AddNode(string id, string ip, string port) returns(bool)
func (_NodeContract *NodeContractTransactor) AddNode(opts *bind.TransactOpts, id string, ip string, port string) (*types.Transaction, error) {
	return _NodeContract.contract.Transact(opts, "AddNode", id, ip, port)
}

// AddNode is a paid mutator transaction binding the contract method 0x92be66c0.
//
// Solidity: function AddNode(string id, string ip, string port) returns(bool)
func (_NodeContract *NodeContractSession) AddNode(id string, ip string, port string) (*types.Transaction, error) {
	return _NodeContract.Contract.AddNode(&_NodeContract.TransactOpts, id, ip, port)
}

// AddNode is a paid mutator transaction binding the contract method 0x92be66c0.
//
// Solidity: function AddNode(string id, string ip, string port) returns(bool)
func (_NodeContract *NodeContractTransactorSession) AddNode(id string, ip string, port string) (*types.Transaction, error) {
	return _NodeContract.Contract.AddNode(&_NodeContract.TransactOpts, id, ip, port)
}

// RemoveNode is a paid mutator transaction binding the contract method 0x2a39f0d0.
//
// Solidity: function RemoveNode() returns(bool)
func (_NodeContract *NodeContractTransactor) RemoveNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeContract.contract.Transact(opts, "RemoveNode")
}

// RemoveNode is a paid mutator transaction binding the contract method 0x2a39f0d0.
//
// Solidity: function RemoveNode() returns(bool)
func (_NodeContract *NodeContractSession) RemoveNode() (*types.Transaction, error) {
	return _NodeContract.Contract.RemoveNode(&_NodeContract.TransactOpts)
}

// RemoveNode is a paid mutator transaction binding the contract method 0x2a39f0d0.
//
// Solidity: function RemoveNode() returns(bool)
func (_NodeContract *NodeContractTransactorSession) RemoveNode() (*types.Transaction, error) {
	return _NodeContract.Contract.RemoveNode(&_NodeContract.TransactOpts)
}
