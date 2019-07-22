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

// MappingReferenceABI is the input ABI used to generate the binding from.
const MappingReferenceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"CheckExistence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"AddNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"RemoveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MappingReferenceBin is the compiled bytecode used for deploying new contracts.
const MappingReferenceBin = `0x608060405234801561001057600080fd5b5061020f806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631240abac811461005b578063ba7249c814610121578063c9ba0b9a14610121575b600080fd5b34801561006757600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261010d95833573ffffffffffffffffffffffffffffffffffffffff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506101d59650505050505050565b604080519115158252519081900360200190f35b34801561012d57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101d395833573ffffffffffffffffffffffffffffffffffffffff1695369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506101de9650505050505050565b005b60009392505050565b5050505600a165627a7a72305820967d78b7249e8509299bf25516584862579494150d5b0e50c3b55ad021e09c2f0029`

// DeployMappingReference deploys a new Ethereum contract, binding an instance of MappingReference to it.
func DeployMappingReference(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MappingReference, error) {
	parsed, err := abi.JSON(strings.NewReader(MappingReferenceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MappingReferenceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MappingReference{MappingReferenceCaller: MappingReferenceCaller{contract: contract}, MappingReferenceTransactor: MappingReferenceTransactor{contract: contract}, MappingReferenceFilterer: MappingReferenceFilterer{contract: contract}}, nil
}

// MappingReference is an auto generated Go binding around an Ethereum contract.
type MappingReference struct {
	MappingReferenceCaller     // Read-only binding to the contract
	MappingReferenceTransactor // Write-only binding to the contract
	MappingReferenceFilterer   // Log filterer for contract events
}

// MappingReferenceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MappingReferenceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MappingReferenceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MappingReferenceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MappingReferenceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MappingReferenceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MappingReferenceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MappingReferenceSession struct {
	Contract     *MappingReference // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MappingReferenceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MappingReferenceCallerSession struct {
	Contract *MappingReferenceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MappingReferenceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MappingReferenceTransactorSession struct {
	Contract     *MappingReferenceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MappingReferenceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MappingReferenceRaw struct {
	Contract *MappingReference // Generic contract binding to access the raw methods on
}

// MappingReferenceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MappingReferenceCallerRaw struct {
	Contract *MappingReferenceCaller // Generic read-only contract binding to access the raw methods on
}

// MappingReferenceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MappingReferenceTransactorRaw struct {
	Contract *MappingReferenceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMappingReference creates a new instance of MappingReference, bound to a specific deployed contract.
func NewMappingReference(address common.Address, backend bind.ContractBackend) (*MappingReference, error) {
	contract, err := bindMappingReference(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MappingReference{MappingReferenceCaller: MappingReferenceCaller{contract: contract}, MappingReferenceTransactor: MappingReferenceTransactor{contract: contract}, MappingReferenceFilterer: MappingReferenceFilterer{contract: contract}}, nil
}

// NewMappingReferenceCaller creates a new read-only instance of MappingReference, bound to a specific deployed contract.
func NewMappingReferenceCaller(address common.Address, caller bind.ContractCaller) (*MappingReferenceCaller, error) {
	contract, err := bindMappingReference(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MappingReferenceCaller{contract: contract}, nil
}

// NewMappingReferenceTransactor creates a new write-only instance of MappingReference, bound to a specific deployed contract.
func NewMappingReferenceTransactor(address common.Address, transactor bind.ContractTransactor) (*MappingReferenceTransactor, error) {
	contract, err := bindMappingReference(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MappingReferenceTransactor{contract: contract}, nil
}

// NewMappingReferenceFilterer creates a new log filterer instance of MappingReference, bound to a specific deployed contract.
func NewMappingReferenceFilterer(address common.Address, filterer bind.ContractFilterer) (*MappingReferenceFilterer, error) {
	contract, err := bindMappingReference(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MappingReferenceFilterer{contract: contract}, nil
}

// bindMappingReference binds a generic wrapper to an already deployed contract.
func bindMappingReference(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MappingReferenceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MappingReference *MappingReferenceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MappingReference.Contract.MappingReferenceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MappingReference *MappingReferenceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MappingReference.Contract.MappingReferenceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MappingReference *MappingReferenceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MappingReference.Contract.MappingReferenceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MappingReference *MappingReferenceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MappingReference.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MappingReference *MappingReferenceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MappingReference.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MappingReference *MappingReferenceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MappingReference.Contract.contract.Transact(opts, method, params...)
}

// AddNode is a paid mutator transaction binding the contract method 0xba7249c8.
//
// Solidity: function AddNode(address , string , string ) returns()
func (_MappingReference *MappingReferenceTransactor) AddNode(opts *bind.TransactOpts, arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _MappingReference.contract.Transact(opts, "AddNode", arg0, arg1, arg2)
}

// AddNode is a paid mutator transaction binding the contract method 0xba7249c8.
//
// Solidity: function AddNode(address , string , string ) returns()
func (_MappingReference *MappingReferenceSession) AddNode(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _MappingReference.Contract.AddNode(&_MappingReference.TransactOpts, arg0, arg1, arg2)
}

// AddNode is a paid mutator transaction binding the contract method 0xba7249c8.
//
// Solidity: function AddNode(address , string , string ) returns()
func (_MappingReference *MappingReferenceTransactorSession) AddNode(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _MappingReference.Contract.AddNode(&_MappingReference.TransactOpts, arg0, arg1, arg2)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x1240abac.
//
// Solidity: function CheckExistence(address , string , string ) returns(bool)
func (_MappingReference *MappingReferenceTransactor) CheckExistence(opts *bind.TransactOpts, arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _MappingReference.contract.Transact(opts, "CheckExistence", arg0, arg1, arg2)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x1240abac.
//
// Solidity: function CheckExistence(address , string , string ) returns(bool)
func (_MappingReference *MappingReferenceSession) CheckExistence(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _MappingReference.Contract.CheckExistence(&_MappingReference.TransactOpts, arg0, arg1, arg2)
}

// CheckExistence is a paid mutator transaction binding the contract method 0x1240abac.
//
// Solidity: function CheckExistence(address , string , string ) returns(bool)
func (_MappingReference *MappingReferenceTransactorSession) CheckExistence(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _MappingReference.Contract.CheckExistence(&_MappingReference.TransactOpts, arg0, arg1, arg2)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xc9ba0b9a.
//
// Solidity: function RemoveNode(address , string , string ) returns()
func (_MappingReference *MappingReferenceTransactor) RemoveNode(opts *bind.TransactOpts, arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _MappingReference.contract.Transact(opts, "RemoveNode", arg0, arg1, arg2)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xc9ba0b9a.
//
// Solidity: function RemoveNode(address , string , string ) returns()
func (_MappingReference *MappingReferenceSession) RemoveNode(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _MappingReference.Contract.RemoveNode(&_MappingReference.TransactOpts, arg0, arg1, arg2)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xc9ba0b9a.
//
// Solidity: function RemoveNode(address , string , string ) returns()
func (_MappingReference *MappingReferenceTransactorSession) RemoveNode(arg0 common.Address, arg1 string, arg2 string) (*types.Transaction, error) {
	return _MappingReference.Contract.RemoveNode(&_MappingReference.TransactOpts, arg0, arg1, arg2)
}

// NodeContractABI is the input ABI used to generate the binding from.
const NodeContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"GetNodePort\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"GetNodeId\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"RemoveNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetNodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"GetNodeIp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"ip\",\"type\":\"string\"},{\"name\":\"port\",\"type\":\"string\"}],\"name\":\"AddNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"mappingAddress\",\"type\":\"address\"}],\"name\":\"UpdateNodeMappingAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetNodeMappingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"setCollateralRequirement\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// NodeContractBin is the compiled bytecode used for deploying new contracts.
const NodeContractBin = `0x608060405234801561001057600080fd5b50604051602080610c8f8339810160405251600060025560045560038054600160a060020a03191633179055610c448061004b6000396000f30060806040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631cb3b230811461009257806326651d4c146101285780632a39f0d01461014957806370f2a971146101725780638ee1d7ab1461019957806392be66c0146101ba578063af05bec014610282578063de28eed5146102a5575b600080fd5b34801561009e57600080fd5b506100b3600160a060020a03600435166102d6565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100ed5781810151838201526020016100d5565b50505050905090810190601f16801561011a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561013457600080fd5b506100b3600160a060020a0360043516610382565b34801561015557600080fd5b5061015e6103fa565b604080519115158252519081900360200190f35b34801561017e57600080fd5b506101876106c9565b60408051918252519081900360200190f35b3480156101a557600080fd5b506100b3600160a060020a03600435166106cf565b6040805160206004803580820135601f810184900484028501840190955284845261015e94369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506107449650505050505050565b34801561028e57600080fd5b506102a3600160a060020a0360043516610ae4565b005b3480156102b157600080fd5b506102ba610b2a565b60408051600160a060020a039092168252519081900360200190f35b600160a060020a0381166000908152602081815260409182902060040180548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156103765780601f1061034b57610100808354040283529160200191610376565b820191906000526020600020905b81548152906001019060200180831161035957829003601f168201915b50505050509050919050565b600160a060020a038116600090815260208181526040918290206002908101805484516001821615610100026000190190911692909204601f810184900484028301840190945283825260609391929091908301828280156103765780601f1061034b57610100808354040283529160200191610376565b600454336000908152602081905260408120600501549091670de0b6b3a7640000021461042357fe5b3360008181526020819052604080822060050154905181156108fc0292818181858888f1935050505015801561045d573d6000803e3d6000fd5b50600554336000818152602081905260409081902090517fc9ba0b9a000000000000000000000000000000000000000000000000000000008152600481018381526060602483019081526002808501805460001960018216156101000201169190910460648501819052600160a060020a039097169663c9ba0b9a96959194600390920193929160448101916084909101908690801561053e5780601f106105135761010080835404028352916020019161053e565b820191906000526020600020905b81548152906001019060200180831161052157829003601f168201915b50508381038252845460026000196101006001841615020190911604808252602090910190859080156105b25780601f10610587576101008083540402835291602001916105b2565b820191906000526020600020905b81548152906001019060200180831161059557829003601f168201915b505095505050505050600060405180830381600087803b1580156105d557600080fd5b505af11580156105e9573d6000803e3d6000fd5b5050336000818152602081815260408083208054600280546000199081018752600180875285882054600160a060020a0390811689528888528689208590558354830189528188528689205494895286892080549590911673ffffffffffffffffffffffffffffffffffffffff19958616179055825490910187529386208054831690559585529284905283815590810180549092169091559350915061069290830182610b39565b6106a0600383016000610b39565b6106ae600483016000610b39565b50600060059190910155506002805460001901905560015b90565b60025490565b600160a060020a0381166000908152602081815260409182902060030180548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156103765780601f1061034b57610100808354040283529160200191610376565b600080600454670de0b6b3a764000002341480156108c557506005546040517f1240abac0000000000000000000000000000000000000000000000000000000081523360048201818152606060248401908152895160648501528951600160a060020a0390951694631240abac948b938b93909290916044810191608490910190602087019080838360005b838110156107e85781810151838201526020016107d0565b50505050905090810190601f1680156108155780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610848578181015183820152602001610830565b50505050905090810190601f1680156108755780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561089757600080fd5b505af11580156108ab573d6000803e3d6000fd5b505050506040513d60208110156108c157600080fd5b5051155b15156108cd57fe5b503360008181526020818152604090912060018101805473ffffffffffffffffffffffffffffffffffffffff1916909317909255346005830155855161091b91600284019190880190610b80565b5083516109319060038301906020870190610b80565b5082516109479060048301906020860190610b80565b506002548082556000908152600160209081526040808320805473ffffffffffffffffffffffffffffffffffffffff19163390811790915560055491517fba7249c8000000000000000000000000000000000000000000000000000000008152600481018281526060602483019081528b5160648401528b51600160a060020a03959095169663ba7249c89694958d958d95604481019360849091019291880191908190849084905b83811015610a085781810151838201526020016109f0565b50505050905090810190601f168015610a355780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610a68578181015183820152602001610a50565b50505050905090810190601f168015610a955780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015610ab757600080fd5b505af1158015610acb573d6000803e3d6000fd5b5050600280546001908101909155979650505050505050565b600354600160a060020a03163314610afb57600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600554600160a060020a031690565b50805460018160011615610100020316600290046000825580601f10610b5f5750610b7d565b601f016020900490600052602060002090810190610b7d9190610bfe565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610bc157805160ff1916838001178555610bee565b82800160010185558215610bee579182015b82811115610bee578251825591602001919060010190610bd3565b50610bfa929150610bfe565b5090565b6106c691905b80821115610bfa5760008155600101610c045600a165627a7a7230582080c5f92ff2997e4da7ecd22d248ce7ac3d3bb71ac0d4725eda7a7cd8518a9a5c0029`

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

// GetNodeCount is a free data retrieval call binding the contract method 0x70f2a971.
//
// Solidity: function GetNodeCount() constant returns(uint256)
func (_NodeContract *NodeContractCaller) GetNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeContract.contract.Call(opts, out, "GetNodeCount")
	return *ret0, err
}

// GetNodeCount is a free data retrieval call binding the contract method 0x70f2a971.
//
// Solidity: function GetNodeCount() constant returns(uint256)
func (_NodeContract *NodeContractSession) GetNodeCount() (*big.Int, error) {
	return _NodeContract.Contract.GetNodeCount(&_NodeContract.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x70f2a971.
//
// Solidity: function GetNodeCount() constant returns(uint256)
func (_NodeContract *NodeContractCallerSession) GetNodeCount() (*big.Int, error) {
	return _NodeContract.Contract.GetNodeCount(&_NodeContract.CallOpts)
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

// GetNodeMappingAddress is a free data retrieval call binding the contract method 0xde28eed5.
//
// Solidity: function GetNodeMappingAddress() constant returns(address)
func (_NodeContract *NodeContractCaller) GetNodeMappingAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NodeContract.contract.Call(opts, out, "GetNodeMappingAddress")
	return *ret0, err
}

// GetNodeMappingAddress is a free data retrieval call binding the contract method 0xde28eed5.
//
// Solidity: function GetNodeMappingAddress() constant returns(address)
func (_NodeContract *NodeContractSession) GetNodeMappingAddress() (common.Address, error) {
	return _NodeContract.Contract.GetNodeMappingAddress(&_NodeContract.CallOpts)
}

// GetNodeMappingAddress is a free data retrieval call binding the contract method 0xde28eed5.
//
// Solidity: function GetNodeMappingAddress() constant returns(address)
func (_NodeContract *NodeContractCallerSession) GetNodeMappingAddress() (common.Address, error) {
	return _NodeContract.Contract.GetNodeMappingAddress(&_NodeContract.CallOpts)
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
