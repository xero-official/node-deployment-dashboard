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
const NodeContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"GetNodePort\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"GetNodeId\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"RemoveNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"GetNodeIp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"ip\",\"type\":\"string\"},{\"name\":\"port\",\"type\":\"string\"}],\"name\":\"AddNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"mappingAddress\",\"type\":\"address\"}],\"name\":\"UpdateNodeMappingAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"setCollateralRequirement\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NodeContractBin is the compiled bytecode used for deploying new contracts.
const NodeContractBin = `0x608060405234801561001057600080fd5b50604051602080610c0c8339810160405251600060025560045560038054600160a060020a03191633179055610bc18061004b6000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631cb3b230811461007c57806326651d4c146101125780632a39f0d0146101335780638ee1d7ab1461015c57806392be66c01461017d578063af05bec014610245575b600080fd5b34801561008857600080fd5b5061009d600160a060020a0360043516610268565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100d75781810151838201526020016100bf565b50505050905090810190601f1680156101045780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561011e57600080fd5b5061009d600160a060020a0360043516610314565b34801561013f57600080fd5b5061014861038c565b604080519115158252519081900360200190f35b34801561016857600080fd5b5061009d600160a060020a036004351661065b565b6040805160206004803580820135601f810184900484028501840190955284845261014894369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506106d09650505050505050565b34801561025157600080fd5b50610266600160a060020a0360043516610a70565b005b600160a060020a0381166000908152602081815260409182902060040180548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156103085780601f106102dd57610100808354040283529160200191610308565b820191906000526020600020905b8154815290600101906020018083116102eb57829003601f168201915b50505050509050919050565b600160a060020a038116600090815260208181526040918290206002908101805484516001821615610100026000190190911692909204601f810184900484028301840190945283825260609391929091908301828280156103085780601f106102dd57610100808354040283529160200191610308565b600454336000908152602081905260408120600501549091670de0b6b3a764000002146103b557fe5b3360008181526020819052604080822060050154905181156108fc0292818181858888f193505050501580156103ef573d6000803e3d6000fd5b50600554336000818152602081905260409081902090517fc9ba0b9a000000000000000000000000000000000000000000000000000000008152600481018381526060602483019081526002808501805460001960018216156101000201169190910460648501819052600160a060020a039097169663c9ba0b9a9695919460039092019392916044810191608490910190869080156104d05780601f106104a5576101008083540402835291602001916104d0565b820191906000526020600020905b8154815290600101906020018083116104b357829003601f168201915b50508381038252845460026000196101006001841615020190911604808252602090910190859080156105445780601f1061051957610100808354040283529160200191610544565b820191906000526020600020905b81548152906001019060200180831161052757829003601f168201915b505095505050505050600060405180830381600087803b15801561056757600080fd5b505af115801561057b573d6000803e3d6000fd5b5050336000818152602081815260408083208054600280546000199081018752600180875285882054600160a060020a0390811689528888528689208590558354830189528188528689205494895286892080549590911673ffffffffffffffffffffffffffffffffffffffff19958616179055825490910187529386208054831690559585529284905283815590810180549092169091559350915061062490830182610ab6565b610632600383016000610ab6565b610640600483016000610ab6565b50600060059190910155506002805460001901905560015b90565b600160a060020a0381166000908152602081815260409182902060030180548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156103085780601f106102dd57610100808354040283529160200191610308565b600080600454670de0b6b3a7640000023414801561085157506005546040517f1240abac0000000000000000000000000000000000000000000000000000000081523360048201818152606060248401908152895160648501528951600160a060020a0390951694631240abac948b938b93909290916044810191608490910190602087019080838360005b8381101561077457818101518382015260200161075c565b50505050905090810190601f1680156107a15780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156107d45781810151838201526020016107bc565b50505050905090810190601f1680156108015780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561082357600080fd5b505af1158015610837573d6000803e3d6000fd5b505050506040513d602081101561084d57600080fd5b5051155b151561085957fe5b503360008181526020818152604090912060018101805473ffffffffffffffffffffffffffffffffffffffff191690931790925534600583015585516108a791600284019190880190610afd565b5083516108bd9060038301906020870190610afd565b5082516108d39060048301906020860190610afd565b506002548082556000908152600160209081526040808320805473ffffffffffffffffffffffffffffffffffffffff19163390811790915560055491517fba7249c8000000000000000000000000000000000000000000000000000000008152600481018281526060602483019081528b5160648401528b51600160a060020a03959095169663ba7249c89694958d958d95604481019360849091019291880191908190849084905b8381101561099457818101518382015260200161097c565b50505050905090810190601f1680156109c15780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156109f45781810151838201526020016109dc565b50505050905090810190601f168015610a215780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015610a4357600080fd5b505af1158015610a57573d6000803e3d6000fd5b5050600280546001908101909155979650505050505050565b600354600160a060020a03163314610a8757600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b50805460018160011615610100020316600290046000825580601f10610adc5750610afa565b601f016020900490600052602060002090810190610afa9190610b7b565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b3e57805160ff1916838001178555610b6b565b82800160010185558215610b6b579182015b82811115610b6b578251825591602001919060010190610b50565b50610b77929150610b7b565b5090565b61065891905b80821115610b775760008155600101610b815600a165627a7a723058207431ed08be66ea666f3e5851f7e3b9d9619f5216af734e5ba8ed8733863ba1cf0029`

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

// UpdateNodeMappingAddress is a paid mutator transaction binding the contract method 0xaf05bec0.
//
// Solidity: function UpdateNodeMappingAddress(address mappingAddress) returns()
func (_NodeContract *NodeContractTransactor) UpdateNodeMappingAddress(opts *bind.TransactOpts, mappingAddress common.Address) (*types.Transaction, error) {
	return _NodeContract.contract.Transact(opts, "UpdateNodeMappingAddress", mappingAddress)
}

// UpdateNodeMappingAddress is a paid mutator transaction binding the contract method 0xaf05bec0.
//
// Solidity: function UpdateNodeMappingAddress(address mappingAddress) returns()
func (_NodeContract *NodeContractSession) UpdateNodeMappingAddress(mappingAddress common.Address) (*types.Transaction, error) {
	return _NodeContract.Contract.UpdateNodeMappingAddress(&_NodeContract.TransactOpts, mappingAddress)
}

// UpdateNodeMappingAddress is a paid mutator transaction binding the contract method 0xaf05bec0.
//
// Solidity: function UpdateNodeMappingAddress(address mappingAddress) returns()
func (_NodeContract *NodeContractTransactorSession) UpdateNodeMappingAddress(mappingAddress common.Address) (*types.Transaction, error) {
	return _NodeContract.Contract.UpdateNodeMappingAddress(&_NodeContract.TransactOpts, mappingAddress)
}

// NodeMappingContractABI is the input ABI used to generate the binding from.
const NodeMappingContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"CheckExistence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"AddNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"RemoveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NodeMappingContractBin is the compiled bytecode used for deploying new contracts.
const NodeMappingContractBin = `0x608060405234801561001057600080fd5b5061020f806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631240abac811461005b578063ba7249c814610121578063c9ba0b9a14610121575b600080fd5b34801561006757600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261010d95833573ffffffffffffffffffffffffffffffffffffffff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506101d59650505050505050565b604080519115158252519081900360200190f35b34801561012d57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101d395833573ffffffffffffffffffffffffffffffffffffffff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506101de9650505050505050565b005b60009392505050565b5050505600a165627a7a72305820491ad5395983a318b46042a328089375d83e29fb3f815b8a6a45dbee1c8c68e80029`

// DeployNodeMappingContract deploys a new Ethereum contract, binding an instance of NodeMappingContract to it.
func DeployNodeMappingContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeMappingContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeMappingContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeMappingContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeMappingContract{NodeMappingContractCaller: NodeMappingContractCaller{contract: contract}, NodeMappingContractTransactor: NodeMappingContractTransactor{contract: contract}, NodeMappingContractFilterer: NodeMappingContractFilterer{contract: contract}}, nil
}

// NodeMappingContract is an auto generated Go binding around an Ethereum contract.
type NodeMappingContract struct {
	NodeMappingContractCaller     // Read-only binding to the contract
	NodeMappingContractTransactor // Write-only binding to the contract
	NodeMappingContractFilterer   // Log filterer for contract events
}

// NodeMappingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeMappingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeMappingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeMappingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeMappingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeMappingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeMappingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeMappingContractSession struct {
	Contract     *NodeMappingContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NodeMappingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeMappingContractCallerSession struct {
	Contract *NodeMappingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// NodeMappingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeMappingContractTransactorSession struct {
	Contract     *NodeMappingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// NodeMappingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeMappingContractRaw struct {
	Contract *NodeMappingContract // Generic contract binding to access the raw methods on
}

// NodeMappingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeMappingContractCallerRaw struct {
	Contract *NodeMappingContractCaller // Generic read-only contract binding to access the raw methods on
}

// NodeMappingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeMappingContractTransactorRaw struct {
	Contract *NodeMappingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeMappingContract creates a new instance of NodeMappingContract, bound to a specific deployed contract.
func NewNodeMappingContract(address common.Address, backend bind.ContractBackend) (*NodeMappingContract, error) {
	contract, err := bindNodeMappingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeMappingContract{NodeMappingContractCaller: NodeMappingContractCaller{contract: contract}, NodeMappingContractTransactor: NodeMappingContractTransactor{contract: contract}, NodeMappingContractFilterer: NodeMappingContractFilterer{contract: contract}}, nil
}

// NewNodeMappingContractCaller creates a new read-only instance of NodeMappingContract, bound to a specific deployed contract.
func NewNodeMappingContractCaller(address common.Address, caller bind.ContractCaller) (*NodeMappingContractCaller, error) {
	contract, err := bindNodeMappingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeMappingContractCaller{contract: contract}, nil
}

// NewNodeMappingContractTransactor creates a new write-only instance of NodeMappingContract, bound to a specific deployed contract.
func NewNodeMappingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeMappingContractTransactor, error) {
	contract, err := bindNodeMappingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeMappingContractTransactor{contract: contract}, nil
}

// NewNodeMappingContractFilterer creates a new log filterer instance of NodeMappingContract, bound to a specific deployed contract.
func NewNodeMappingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeMappingContractFilterer, error) {
	contract, err := bindNodeMappingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeMappingContractFilterer{contract: contract}, nil
}

// bindNodeMappingContract binds a generic wrapper to an already deployed contract.
func bindNodeMappingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeMappingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeMappingContract *NodeMappingContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeMappingContract.Contract.NodeMappingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeMappingContract *NodeMappingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeMappingContract.Contract.NodeMappingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeMappingContract *NodeMappingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeMappingContract.Contract.NodeMappingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeMappingContract *NodeMappingContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeMappingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeMappingContract *NodeMappingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeMappingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeMappingContract *NodeMappingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeMappingContract.Contract.contract.Transact(opts, method, params...)
}

// AddNode is a paid mutator transaction binding the contract method 0xba7249c8.
//
// Solidity: function AddNode(address , string , string ) returns()
func (_NodeMappingContract *NodeMappingContractTransactor) AddNode(opts *bind.TransactOpts, arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _NodeMappingContract.contract.Transact(opts, "AddNode", arg0, arg1, arg2)
}

// AddNode is a paid mutator transaction binding the contract method 0xba7249c8.
//
// Solidity: function AddNode(address , string , string ) returns()
func (_NodeMappingContract *NodeMappingContractSession) AddNode(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _NodeMappingContract.Contract.AddNode(&_NodeMappingContract.TransactOpts, arg0, arg1, arg2)
}

// AddNode is a paid mutator transaction binding the contract method 0xba7249c8.
//
// Solidity: function AddNode(address , string , string ) returns()
func (_NodeMappingContract *NodeMappingContractTransactorSession) AddNode(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _NodeMappingContract.Contract.AddNode(&_NodeMappingContract.TransactOpts, arg0, arg1, arg2)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x1240abac.
//
// Solidity: function CheckExistence(address , string , string ) returns(bool)
func (_NodeMappingContract *NodeMappingContractTransactor) CheckExistence(opts *bind.TransactOpts, arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _NodeMappingContract.contract.Transact(opts, "CheckExistence", arg0, arg1, arg2)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x1240abac.
//
// Solidity: function CheckExistence(address , string , string ) returns(bool)
func (_NodeMappingContract *NodeMappingContractSession) CheckExistence(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _NodeMappingContract.Contract.CheckExistence(&_NodeMappingContract.TransactOpts, arg0, arg1, arg2)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x1240abac.
//
// Solidity: function CheckExistence(address , string , string ) returns(bool)
func (_NodeMappingContract *NodeMappingContractTransactorSession) CheckExistence(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _NodeMappingContract.Contract.CheckExistence(&_NodeMappingContract.TransactOpts, arg0, arg1, arg2)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xc9ba0b9a.
//
// Solidity: function RemoveNode(address , string , string ) returns()
func (_NodeMappingContract *NodeMappingContractTransactor) RemoveNode(opts *bind.TransactOpts, arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _NodeMappingContract.contract.Transact(opts, "RemoveNode", arg0, arg1, arg2)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xc9ba0b9a.
//
// Solidity: function RemoveNode(address , string , string ) returns()
func (_NodeMappingContract *NodeMappingContractSession) RemoveNode(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _NodeMappingContract.Contract.RemoveNode(&_NodeMappingContract.TransactOpts, arg0, arg1, arg2)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xc9ba0b9a.
//
// Solidity: function RemoveNode(address , string , string ) returns()
func (_NodeMappingContract *NodeMappingContractTransactorSession) RemoveNode(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _NodeMappingContract.Contract.RemoveNode(&_NodeMappingContract.TransactOpts, arg0, arg1, arg2)
}
