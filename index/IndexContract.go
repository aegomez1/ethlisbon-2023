// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package index

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

// IndexMetaData contains all meta data concerning the Index contract.
var IndexMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipfsCID\",\"type\":\"string\"}],\"name\":\"SubmissionAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"addSubmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"submissionMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsCID\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061086a8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80636bae848414610038578063ad22f70b14610054575b5f80fd5b610052600480360381019061004d9190610406565b610085565b005b61006e60048036038101906100699190610460565b61019b565b60405161007c929190610514565b60405180910390f35b60405180604001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001828152505f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816101459190610745565b509050508173ffffffffffffffffffffffffffffffffffffffff167f3840d88f03a647ee04696cbde6b92a0acb6ba6dd37470c27b86536392f8f7e478260405161018f9190610814565b60405180910390a25050565b5f602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010180546101de9061056f565b80601f016020809104026020016040519081016040528092919081815260200182805461020a9061056f565b80156102555780601f1061022c57610100808354040283529160200191610255565b820191905f5260205f20905b81548152906001019060200180831161023857829003601f168201915b5050505050905082565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61029982610270565b9050919050565b6102a98161028f565b81146102b3575f80fd5b50565b5f813590506102c4816102a0565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610318826102d2565b810181811067ffffffffffffffff82111715610337576103366102e2565b5b80604052505050565b5f61034961025f565b9050610355828261030f565b919050565b5f67ffffffffffffffff821115610374576103736102e2565b5b61037d826102d2565b9050602081019050919050565b828183375f83830152505050565b5f6103aa6103a58461035a565b610340565b9050828152602081018484840111156103c6576103c56102ce565b5b6103d184828561038a565b509392505050565b5f82601f8301126103ed576103ec6102ca565b5b81356103fd848260208601610398565b91505092915050565b5f806040838503121561041c5761041b610268565b5b5f610429858286016102b6565b925050602083013567ffffffffffffffff81111561044a5761044961026c565b5b610456858286016103d9565b9150509250929050565b5f6020828403121561047557610474610268565b5b5f610482848285016102b6565b91505092915050565b6104948161028f565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156104d15780820151818401526020810190506104b6565b5f8484015250505050565b5f6104e68261049a565b6104f081856104a4565b93506105008185602086016104b4565b610509816102d2565b840191505092915050565b5f6040820190506105275f83018561048b565b818103602083015261053981846104dc565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061058657607f821691505b60208210810361059957610598610542565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026105fb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826105c0565b61060586836105c0565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61064961064461063f8461061d565b610626565b61061d565b9050919050565b5f819050919050565b6106628361062f565b61067661066e82610650565b8484546105cc565b825550505050565b5f90565b61068a61067e565b610695818484610659565b505050565b5b818110156106b8576106ad5f82610682565b60018101905061069b565b5050565b601f8211156106fd576106ce8161059f565b6106d7846105b1565b810160208510156106e6578190505b6106fa6106f2856105b1565b83018261069a565b50505b505050565b5f82821c905092915050565b5f61071d5f1984600802610702565b1980831691505092915050565b5f610735838361070e565b9150826002028217905092915050565b61074e8261049a565b67ffffffffffffffff811115610767576107666102e2565b5b610771825461056f565b61077c8282856106bc565b5f60209050601f8311600181146107ad575f841561079b578287015190505b6107a5858261072a565b86555061080c565b601f1984166107bb8661059f565b5f5b828110156107e2578489015182556001820191506020850194506020810190506107bd565b868310156107ff57848901516107fb601f89168261070e565b8355505b6001600288020188555050505b505050505050565b5f6020820190508181035f83015261082c81846104dc565b90509291505056fea2646970667358221220068031f96e7471b4728d63d3de211ee0d259825f3d2237902d5672ea3ab4154764736f6c63430008140033",
}

// IndexABI is the input ABI used to generate the binding from.
// Deprecated: Use IndexMetaData.ABI instead.
var IndexABI = IndexMetaData.ABI

// IndexBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IndexMetaData.Bin instead.
var IndexBin = IndexMetaData.Bin

// DeployIndex deploys a new Ethereum contract, binding an instance of Index to it.
func DeployIndex(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Index, error) {
	parsed, err := IndexMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IndexBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Index{IndexCaller: IndexCaller{contract: contract}, IndexTransactor: IndexTransactor{contract: contract}, IndexFilterer: IndexFilterer{contract: contract}}, nil
}

// Index is an auto generated Go binding around an Ethereum contract.
type Index struct {
	IndexCaller     // Read-only binding to the contract
	IndexTransactor // Write-only binding to the contract
	IndexFilterer   // Log filterer for contract events
}

// IndexCaller is an auto generated read-only Go binding around an Ethereum contract.
type IndexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IndexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IndexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IndexSession struct {
	Contract     *Index            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IndexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IndexCallerSession struct {
	Contract *IndexCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IndexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IndexTransactorSession struct {
	Contract     *IndexTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IndexRaw is an auto generated low-level Go binding around an Ethereum contract.
type IndexRaw struct {
	Contract *Index // Generic contract binding to access the raw methods on
}

// IndexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IndexCallerRaw struct {
	Contract *IndexCaller // Generic read-only contract binding to access the raw methods on
}

// IndexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IndexTransactorRaw struct {
	Contract *IndexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIndex creates a new instance of Index, bound to a specific deployed contract.
func NewIndex(address common.Address, backend bind.ContractBackend) (*Index, error) {
	contract, err := bindIndex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Index{IndexCaller: IndexCaller{contract: contract}, IndexTransactor: IndexTransactor{contract: contract}, IndexFilterer: IndexFilterer{contract: contract}}, nil
}

// NewIndexCaller creates a new read-only instance of Index, bound to a specific deployed contract.
func NewIndexCaller(address common.Address, caller bind.ContractCaller) (*IndexCaller, error) {
	contract, err := bindIndex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IndexCaller{contract: contract}, nil
}

// NewIndexTransactor creates a new write-only instance of Index, bound to a specific deployed contract.
func NewIndexTransactor(address common.Address, transactor bind.ContractTransactor) (*IndexTransactor, error) {
	contract, err := bindIndex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IndexTransactor{contract: contract}, nil
}

// NewIndexFilterer creates a new log filterer instance of Index, bound to a specific deployed contract.
func NewIndexFilterer(address common.Address, filterer bind.ContractFilterer) (*IndexFilterer, error) {
	contract, err := bindIndex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IndexFilterer{contract: contract}, nil
}

// bindIndex binds a generic wrapper to an already deployed contract.
func bindIndex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IndexMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Index *IndexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Index.Contract.IndexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Index *IndexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Index.Contract.IndexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Index *IndexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Index.Contract.IndexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Index *IndexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Index.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Index *IndexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Index.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Index *IndexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Index.Contract.contract.Transact(opts, method, params...)
}

// SubmissionMap is a free data retrieval call binding the contract method 0xad22f70b.
//
// Solidity: function submissionMap(address ) view returns(address contractAddress, string ipfsCID)
func (_Index *IndexCaller) SubmissionMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	ContractAddress common.Address
	IpfsCID         string
}, error) {
	var out []interface{}
	err := _Index.contract.Call(opts, &out, "submissionMap", arg0)

	outstruct := new(struct {
		ContractAddress common.Address
		IpfsCID         string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ContractAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IpfsCID = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// SubmissionMap is a free data retrieval call binding the contract method 0xad22f70b.
//
// Solidity: function submissionMap(address ) view returns(address contractAddress, string ipfsCID)
func (_Index *IndexSession) SubmissionMap(arg0 common.Address) (struct {
	ContractAddress common.Address
	IpfsCID         string
}, error) {
	return _Index.Contract.SubmissionMap(&_Index.CallOpts, arg0)
}

// SubmissionMap is a free data retrieval call binding the contract method 0xad22f70b.
//
// Solidity: function submissionMap(address ) view returns(address contractAddress, string ipfsCID)
func (_Index *IndexCallerSession) SubmissionMap(arg0 common.Address) (struct {
	ContractAddress common.Address
	IpfsCID         string
}, error) {
	return _Index.Contract.SubmissionMap(&_Index.CallOpts, arg0)
}

// AddSubmission is a paid mutator transaction binding the contract method 0x6bae8484.
//
// Solidity: function addSubmission(address contractAddr, string cid) returns()
func (_Index *IndexTransactor) AddSubmission(opts *bind.TransactOpts, contractAddr common.Address, cid string) (*types.Transaction, error) {
	return _Index.contract.Transact(opts, "addSubmission", contractAddr, cid)
}

// AddSubmission is a paid mutator transaction binding the contract method 0x6bae8484.
//
// Solidity: function addSubmission(address contractAddr, string cid) returns()
func (_Index *IndexSession) AddSubmission(contractAddr common.Address, cid string) (*types.Transaction, error) {
	return _Index.Contract.AddSubmission(&_Index.TransactOpts, contractAddr, cid)
}

// AddSubmission is a paid mutator transaction binding the contract method 0x6bae8484.
//
// Solidity: function addSubmission(address contractAddr, string cid) returns()
func (_Index *IndexTransactorSession) AddSubmission(contractAddr common.Address, cid string) (*types.Transaction, error) {
	return _Index.Contract.AddSubmission(&_Index.TransactOpts, contractAddr, cid)
}

// IndexSubmissionAddedIterator is returned from FilterSubmissionAdded and is used to iterate over the raw logs and unpacked data for SubmissionAdded events raised by the Index contract.
type IndexSubmissionAddedIterator struct {
	Event *IndexSubmissionAdded // Event containing the contract specifics and raw log

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
func (it *IndexSubmissionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexSubmissionAdded)
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
		it.Event = new(IndexSubmissionAdded)
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
func (it *IndexSubmissionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexSubmissionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexSubmissionAdded represents a SubmissionAdded event raised by the Index contract.
type IndexSubmissionAdded struct {
	ContractAddress common.Address
	IpfsCID         string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSubmissionAdded is a free log retrieval operation binding the contract event 0x3840d88f03a647ee04696cbde6b92a0acb6ba6dd37470c27b86536392f8f7e47.
//
// Solidity: event SubmissionAdded(address indexed contractAddress, string ipfsCID)
func (_Index *IndexFilterer) FilterSubmissionAdded(opts *bind.FilterOpts, contractAddress []common.Address) (*IndexSubmissionAddedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Index.contract.FilterLogs(opts, "SubmissionAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &IndexSubmissionAddedIterator{contract: _Index.contract, event: "SubmissionAdded", logs: logs, sub: sub}, nil
}

// WatchSubmissionAdded is a free log subscription operation binding the contract event 0x3840d88f03a647ee04696cbde6b92a0acb6ba6dd37470c27b86536392f8f7e47.
//
// Solidity: event SubmissionAdded(address indexed contractAddress, string ipfsCID)
func (_Index *IndexFilterer) WatchSubmissionAdded(opts *bind.WatchOpts, sink chan<- *IndexSubmissionAdded, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Index.contract.WatchLogs(opts, "SubmissionAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexSubmissionAdded)
				if err := _Index.contract.UnpackLog(event, "SubmissionAdded", log); err != nil {
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

// ParseSubmissionAdded is a log parse operation binding the contract event 0x3840d88f03a647ee04696cbde6b92a0acb6ba6dd37470c27b86536392f8f7e47.
//
// Solidity: event SubmissionAdded(address indexed contractAddress, string ipfsCID)
func (_Index *IndexFilterer) ParseSubmissionAdded(log types.Log) (*IndexSubmissionAdded, error) {
	event := new(IndexSubmissionAdded)
	if err := _Index.contract.UnpackLog(event, "SubmissionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
