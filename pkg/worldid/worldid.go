// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package worldid

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

// WorldIDIdentityManagerImplV1RootInfo is an auto generated low-level Go binding around an user-defined struct.
type WorldIDIdentityManagerImplV1RootInfo struct {
	Root                *big.Int
	SupersededTimestamp *big.Int
	IsValid             bool
}

// WorldIdMetaData contains all meta data concerning the WorldId contract.
var WorldIdMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"NO_SUCH_ROOT\",\"inputs\":[],\"outputs\":[{\"name\":\"rootInfo\",\"type\":\"tuple\",\"internalType\":\"structWorldIDIdentityManagerImplV1.RootInfo\",\"components\":[{\"name\":\"root\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supersededTimestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"isValid\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateIdentityRegistrationInputHash\",\"inputs\":[{\"name\":\"startIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"preRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"postRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"identityCommitments\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisterIdentitiesVerifierLookupTableAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootHistoryExpiry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSemaphoreVerifierAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTreeDepth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"identityOperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_treeDepth\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"initialRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_batchInsertionVerifiers\",\"type\":\"address\",\"internalType\":\"contractVerifierLookupTable\"},{\"name\":\"_batchUpdateVerifiers\",\"type\":\"address\",\"internalType\":\"contractVerifierLookupTable\"},{\"name\":\"_semaphoreVerifier\",\"type\":\"address\",\"internalType\":\"contractISemaphoreVerifier\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structWorldIDIdentityManagerImplV1.RootInfo\",\"components\":[{\"name\":\"root\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supersededTimestamp\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"isValid\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerIdentities\",\"inputs\":[{\"name\":\"insertionProof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"},{\"name\":\"preRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"identityCommitments\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"postRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireValidRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setIdentityOperator\",\"inputs\":[{\"name\":\"newIdentityOperator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRegisterIdentitiesVerifierLookupTable\",\"inputs\":[{\"name\":\"newTable\",\"type\":\"address\",\"internalType\":\"contractVerifierLookupTable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRootHistoryExpiry\",\"inputs\":[{\"name\":\"newExpiryTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSemaphoreVerifier\",\"inputs\":[{\"name\":\"newVerifier\",\"type\":\"address\",\"internalType\":\"contractISemaphoreVerifier\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyProof\",\"inputs\":[{\"name\":\"root\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signalHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nullifierHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"externalNullifierHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"uint256[8]\",\"internalType\":\"uint256[8]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DependencyUpdated\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumWorldIDIdentityManagerImplV1.Dependency\"},{\"name\":\"oldAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityOperatorChanged\",\"inputs\":[{\"name\":\"oldOperator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOperator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RootHistoryExpirySet\",\"inputs\":[{\"name\":\"oldExpiryTime\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newExpiryTime\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StateBridgeStateChange\",\"inputs\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TreeChanged\",\"inputs\":[{\"name\":\"preRoot\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumWorldIDIdentityManagerImplV1.TreeChange\"},{\"name\":\"postRoot\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorldIDIdentityManagerImplInitialized\",\"inputs\":[{\"name\":\"_treeDepth\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"initialRoot\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotRenounceOwnership\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpiredRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ImplementationNotInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitment\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStateBridgeAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifierLUT\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedInputLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotLatestRoot\",\"inputs\":[{\"name\":\"providedRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"latestRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ProofValidationFailure\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StateBridgeAlreadyDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StateBridgeAlreadyEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnreducedElement\",\"inputs\":[{\"name\":\"elementType\",\"type\":\"uint8\",\"internalType\":\"enumWorldIDIdentityManagerImplV1.UnreducedElementType\"},{\"name\":\"element\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnsupportedTreeDepth\",\"inputs\":[{\"name\":\"depth\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}]",
}

// WorldIdABI is the input ABI used to generate the binding from.
// Deprecated: Use WorldIdMetaData.ABI instead.
var WorldIdABI = WorldIdMetaData.ABI

// WorldId is an auto generated Go binding around an Ethereum contract.
type WorldId struct {
	WorldIdCaller     // Read-only binding to the contract
	WorldIdTransactor // Write-only binding to the contract
	WorldIdFilterer   // Log filterer for contract events
}

// WorldIdCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorldIdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorldIdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorldIdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorldIdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorldIdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorldIdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorldIdSession struct {
	Contract     *WorldId          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorldIdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorldIdCallerSession struct {
	Contract *WorldIdCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// WorldIdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorldIdTransactorSession struct {
	Contract     *WorldIdTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WorldIdRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorldIdRaw struct {
	Contract *WorldId // Generic contract binding to access the raw methods on
}

// WorldIdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorldIdCallerRaw struct {
	Contract *WorldIdCaller // Generic read-only contract binding to access the raw methods on
}

// WorldIdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorldIdTransactorRaw struct {
	Contract *WorldIdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorldId creates a new instance of WorldId, bound to a specific deployed contract.
func NewWorldId(address common.Address, backend bind.ContractBackend) (*WorldId, error) {
	contract, err := bindWorldId(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorldId{WorldIdCaller: WorldIdCaller{contract: contract}, WorldIdTransactor: WorldIdTransactor{contract: contract}, WorldIdFilterer: WorldIdFilterer{contract: contract}}, nil
}

// NewWorldIdCaller creates a new read-only instance of WorldId, bound to a specific deployed contract.
func NewWorldIdCaller(address common.Address, caller bind.ContractCaller) (*WorldIdCaller, error) {
	contract, err := bindWorldId(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorldIdCaller{contract: contract}, nil
}

// NewWorldIdTransactor creates a new write-only instance of WorldId, bound to a specific deployed contract.
func NewWorldIdTransactor(address common.Address, transactor bind.ContractTransactor) (*WorldIdTransactor, error) {
	contract, err := bindWorldId(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorldIdTransactor{contract: contract}, nil
}

// NewWorldIdFilterer creates a new log filterer instance of WorldId, bound to a specific deployed contract.
func NewWorldIdFilterer(address common.Address, filterer bind.ContractFilterer) (*WorldIdFilterer, error) {
	contract, err := bindWorldId(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorldIdFilterer{contract: contract}, nil
}

// bindWorldId binds a generic wrapper to an already deployed contract.
func bindWorldId(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorldIdMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorldId *WorldIdRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorldId.Contract.WorldIdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorldId *WorldIdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorldId.Contract.WorldIdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorldId *WorldIdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorldId.Contract.WorldIdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorldId *WorldIdCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorldId.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorldId *WorldIdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorldId.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorldId *WorldIdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorldId.Contract.contract.Transact(opts, method, params...)
}

// NOSUCHROOT is a free data retrieval call binding the contract method 0x561f204b.
//
// Solidity: function NO_SUCH_ROOT() pure returns((uint256,uint128,bool) rootInfo)
func (_WorldId *WorldIdCaller) NOSUCHROOT(opts *bind.CallOpts) (WorldIDIdentityManagerImplV1RootInfo, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "NO_SUCH_ROOT")

	if err != nil {
		return *new(WorldIDIdentityManagerImplV1RootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(WorldIDIdentityManagerImplV1RootInfo)).(*WorldIDIdentityManagerImplV1RootInfo)

	return out0, err

}

// NOSUCHROOT is a free data retrieval call binding the contract method 0x561f204b.
//
// Solidity: function NO_SUCH_ROOT() pure returns((uint256,uint128,bool) rootInfo)
func (_WorldId *WorldIdSession) NOSUCHROOT() (WorldIDIdentityManagerImplV1RootInfo, error) {
	return _WorldId.Contract.NOSUCHROOT(&_WorldId.CallOpts)
}

// NOSUCHROOT is a free data retrieval call binding the contract method 0x561f204b.
//
// Solidity: function NO_SUCH_ROOT() pure returns((uint256,uint128,bool) rootInfo)
func (_WorldId *WorldIdCallerSession) NOSUCHROOT() (WorldIDIdentityManagerImplV1RootInfo, error) {
	return _WorldId.Contract.NOSUCHROOT(&_WorldId.CallOpts)
}

// CalculateIdentityRegistrationInputHash is a free data retrieval call binding the contract method 0x8c76a909.
//
// Solidity: function calculateIdentityRegistrationInputHash(uint32 startIndex, uint256 preRoot, uint256 postRoot, uint256[] identityCommitments) view returns(bytes32 hash)
func (_WorldId *WorldIdCaller) CalculateIdentityRegistrationInputHash(opts *bind.CallOpts, startIndex uint32, preRoot *big.Int, postRoot *big.Int, identityCommitments []*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "calculateIdentityRegistrationInputHash", startIndex, preRoot, postRoot, identityCommitments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateIdentityRegistrationInputHash is a free data retrieval call binding the contract method 0x8c76a909.
//
// Solidity: function calculateIdentityRegistrationInputHash(uint32 startIndex, uint256 preRoot, uint256 postRoot, uint256[] identityCommitments) view returns(bytes32 hash)
func (_WorldId *WorldIdSession) CalculateIdentityRegistrationInputHash(startIndex uint32, preRoot *big.Int, postRoot *big.Int, identityCommitments []*big.Int) ([32]byte, error) {
	return _WorldId.Contract.CalculateIdentityRegistrationInputHash(&_WorldId.CallOpts, startIndex, preRoot, postRoot, identityCommitments)
}

// CalculateIdentityRegistrationInputHash is a free data retrieval call binding the contract method 0x8c76a909.
//
// Solidity: function calculateIdentityRegistrationInputHash(uint32 startIndex, uint256 preRoot, uint256 postRoot, uint256[] identityCommitments) view returns(bytes32 hash)
func (_WorldId *WorldIdCallerSession) CalculateIdentityRegistrationInputHash(startIndex uint32, preRoot *big.Int, postRoot *big.Int, identityCommitments []*big.Int) ([32]byte, error) {
	return _WorldId.Contract.CalculateIdentityRegistrationInputHash(&_WorldId.CallOpts, startIndex, preRoot, postRoot, identityCommitments)
}

// GetRegisterIdentitiesVerifierLookupTableAddress is a free data retrieval call binding the contract method 0x8fc22e9f.
//
// Solidity: function getRegisterIdentitiesVerifierLookupTableAddress() view returns(address)
func (_WorldId *WorldIdCaller) GetRegisterIdentitiesVerifierLookupTableAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "getRegisterIdentitiesVerifierLookupTableAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegisterIdentitiesVerifierLookupTableAddress is a free data retrieval call binding the contract method 0x8fc22e9f.
//
// Solidity: function getRegisterIdentitiesVerifierLookupTableAddress() view returns(address)
func (_WorldId *WorldIdSession) GetRegisterIdentitiesVerifierLookupTableAddress() (common.Address, error) {
	return _WorldId.Contract.GetRegisterIdentitiesVerifierLookupTableAddress(&_WorldId.CallOpts)
}

// GetRegisterIdentitiesVerifierLookupTableAddress is a free data retrieval call binding the contract method 0x8fc22e9f.
//
// Solidity: function getRegisterIdentitiesVerifierLookupTableAddress() view returns(address)
func (_WorldId *WorldIdCallerSession) GetRegisterIdentitiesVerifierLookupTableAddress() (common.Address, error) {
	return _WorldId.Contract.GetRegisterIdentitiesVerifierLookupTableAddress(&_WorldId.CallOpts)
}

// GetRootHistoryExpiry is a free data retrieval call binding the contract method 0x43f974cb.
//
// Solidity: function getRootHistoryExpiry() view returns(uint256)
func (_WorldId *WorldIdCaller) GetRootHistoryExpiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "getRootHistoryExpiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRootHistoryExpiry is a free data retrieval call binding the contract method 0x43f974cb.
//
// Solidity: function getRootHistoryExpiry() view returns(uint256)
func (_WorldId *WorldIdSession) GetRootHistoryExpiry() (*big.Int, error) {
	return _WorldId.Contract.GetRootHistoryExpiry(&_WorldId.CallOpts)
}

// GetRootHistoryExpiry is a free data retrieval call binding the contract method 0x43f974cb.
//
// Solidity: function getRootHistoryExpiry() view returns(uint256)
func (_WorldId *WorldIdCallerSession) GetRootHistoryExpiry() (*big.Int, error) {
	return _WorldId.Contract.GetRootHistoryExpiry(&_WorldId.CallOpts)
}

// GetSemaphoreVerifierAddress is a free data retrieval call binding the contract method 0xf2038f95.
//
// Solidity: function getSemaphoreVerifierAddress() view returns(address)
func (_WorldId *WorldIdCaller) GetSemaphoreVerifierAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "getSemaphoreVerifierAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSemaphoreVerifierAddress is a free data retrieval call binding the contract method 0xf2038f95.
//
// Solidity: function getSemaphoreVerifierAddress() view returns(address)
func (_WorldId *WorldIdSession) GetSemaphoreVerifierAddress() (common.Address, error) {
	return _WorldId.Contract.GetSemaphoreVerifierAddress(&_WorldId.CallOpts)
}

// GetSemaphoreVerifierAddress is a free data retrieval call binding the contract method 0xf2038f95.
//
// Solidity: function getSemaphoreVerifierAddress() view returns(address)
func (_WorldId *WorldIdCallerSession) GetSemaphoreVerifierAddress() (common.Address, error) {
	return _WorldId.Contract.GetSemaphoreVerifierAddress(&_WorldId.CallOpts)
}

// GetTreeDepth is a free data retrieval call binding the contract method 0x8e5cdd50.
//
// Solidity: function getTreeDepth() view returns(uint8)
func (_WorldId *WorldIdCaller) GetTreeDepth(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "getTreeDepth")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTreeDepth is a free data retrieval call binding the contract method 0x8e5cdd50.
//
// Solidity: function getTreeDepth() view returns(uint8)
func (_WorldId *WorldIdSession) GetTreeDepth() (uint8, error) {
	return _WorldId.Contract.GetTreeDepth(&_WorldId.CallOpts)
}

// GetTreeDepth is a free data retrieval call binding the contract method 0x8e5cdd50.
//
// Solidity: function getTreeDepth() view returns(uint8)
func (_WorldId *WorldIdCallerSession) GetTreeDepth() (uint8, error) {
	return _WorldId.Contract.GetTreeDepth(&_WorldId.CallOpts)
}

// IdentityOperator is a free data retrieval call binding the contract method 0xf134b6ca.
//
// Solidity: function identityOperator() view returns(address)
func (_WorldId *WorldIdCaller) IdentityOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "identityOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityOperator is a free data retrieval call binding the contract method 0xf134b6ca.
//
// Solidity: function identityOperator() view returns(address)
func (_WorldId *WorldIdSession) IdentityOperator() (common.Address, error) {
	return _WorldId.Contract.IdentityOperator(&_WorldId.CallOpts)
}

// IdentityOperator is a free data retrieval call binding the contract method 0xf134b6ca.
//
// Solidity: function identityOperator() view returns(address)
func (_WorldId *WorldIdCallerSession) IdentityOperator() (common.Address, error) {
	return _WorldId.Contract.IdentityOperator(&_WorldId.CallOpts)
}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(uint256)
func (_WorldId *WorldIdCaller) LatestRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "latestRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(uint256)
func (_WorldId *WorldIdSession) LatestRoot() (*big.Int, error) {
	return _WorldId.Contract.LatestRoot(&_WorldId.CallOpts)
}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(uint256)
func (_WorldId *WorldIdCallerSession) LatestRoot() (*big.Int, error) {
	return _WorldId.Contract.LatestRoot(&_WorldId.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorldId *WorldIdCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorldId *WorldIdSession) Owner() (common.Address, error) {
	return _WorldId.Contract.Owner(&_WorldId.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorldId *WorldIdCallerSession) Owner() (common.Address, error) {
	return _WorldId.Contract.Owner(&_WorldId.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_WorldId *WorldIdCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_WorldId *WorldIdSession) PendingOwner() (common.Address, error) {
	return _WorldId.Contract.PendingOwner(&_WorldId.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_WorldId *WorldIdCallerSession) PendingOwner() (common.Address, error) {
	return _WorldId.Contract.PendingOwner(&_WorldId.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WorldId *WorldIdCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WorldId *WorldIdSession) ProxiableUUID() ([32]byte, error) {
	return _WorldId.Contract.ProxiableUUID(&_WorldId.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WorldId *WorldIdCallerSession) ProxiableUUID() ([32]byte, error) {
	return _WorldId.Contract.ProxiableUUID(&_WorldId.CallOpts)
}

// QueryRoot is a free data retrieval call binding the contract method 0x3f7c178d.
//
// Solidity: function queryRoot(uint256 root) view returns((uint256,uint128,bool))
func (_WorldId *WorldIdCaller) QueryRoot(opts *bind.CallOpts, root *big.Int) (WorldIDIdentityManagerImplV1RootInfo, error) {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "queryRoot", root)

	if err != nil {
		return *new(WorldIDIdentityManagerImplV1RootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(WorldIDIdentityManagerImplV1RootInfo)).(*WorldIDIdentityManagerImplV1RootInfo)

	return out0, err

}

// QueryRoot is a free data retrieval call binding the contract method 0x3f7c178d.
//
// Solidity: function queryRoot(uint256 root) view returns((uint256,uint128,bool))
func (_WorldId *WorldIdSession) QueryRoot(root *big.Int) (WorldIDIdentityManagerImplV1RootInfo, error) {
	return _WorldId.Contract.QueryRoot(&_WorldId.CallOpts, root)
}

// QueryRoot is a free data retrieval call binding the contract method 0x3f7c178d.
//
// Solidity: function queryRoot(uint256 root) view returns((uint256,uint128,bool))
func (_WorldId *WorldIdCallerSession) QueryRoot(root *big.Int) (WorldIDIdentityManagerImplV1RootInfo, error) {
	return _WorldId.Contract.QueryRoot(&_WorldId.CallOpts, root)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_WorldId *WorldIdCaller) RenounceOwnership(opts *bind.CallOpts) error {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "renounceOwnership")

	if err != nil {
		return err
	}

	return err

}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_WorldId *WorldIdSession) RenounceOwnership() error {
	return _WorldId.Contract.RenounceOwnership(&_WorldId.CallOpts)
}

// RenounceOwnership is a free data retrieval call binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() view returns()
func (_WorldId *WorldIdCallerSession) RenounceOwnership() error {
	return _WorldId.Contract.RenounceOwnership(&_WorldId.CallOpts)
}

// RequireValidRoot is a free data retrieval call binding the contract method 0xf2358e1d.
//
// Solidity: function requireValidRoot(uint256 root) view returns()
func (_WorldId *WorldIdCaller) RequireValidRoot(opts *bind.CallOpts, root *big.Int) error {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "requireValidRoot", root)

	if err != nil {
		return err
	}

	return err

}

// RequireValidRoot is a free data retrieval call binding the contract method 0xf2358e1d.
//
// Solidity: function requireValidRoot(uint256 root) view returns()
func (_WorldId *WorldIdSession) RequireValidRoot(root *big.Int) error {
	return _WorldId.Contract.RequireValidRoot(&_WorldId.CallOpts, root)
}

// RequireValidRoot is a free data retrieval call binding the contract method 0xf2358e1d.
//
// Solidity: function requireValidRoot(uint256 root) view returns()
func (_WorldId *WorldIdCallerSession) RequireValidRoot(root *big.Int) error {
	return _WorldId.Contract.RequireValidRoot(&_WorldId.CallOpts, root)
}

// VerifyProof is a free data retrieval call binding the contract method 0x354ca120.
//
// Solidity: function verifyProof(uint256 root, uint256 signalHash, uint256 nullifierHash, uint256 externalNullifierHash, uint256[8] proof) view returns()
func (_WorldId *WorldIdCaller) VerifyProof(opts *bind.CallOpts, root *big.Int, signalHash *big.Int, nullifierHash *big.Int, externalNullifierHash *big.Int, proof [8]*big.Int) error {
	var out []interface{}
	err := _WorldId.contract.Call(opts, &out, "verifyProof", root, signalHash, nullifierHash, externalNullifierHash, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x354ca120.
//
// Solidity: function verifyProof(uint256 root, uint256 signalHash, uint256 nullifierHash, uint256 externalNullifierHash, uint256[8] proof) view returns()
func (_WorldId *WorldIdSession) VerifyProof(root *big.Int, signalHash *big.Int, nullifierHash *big.Int, externalNullifierHash *big.Int, proof [8]*big.Int) error {
	return _WorldId.Contract.VerifyProof(&_WorldId.CallOpts, root, signalHash, nullifierHash, externalNullifierHash, proof)
}

// VerifyProof is a free data retrieval call binding the contract method 0x354ca120.
//
// Solidity: function verifyProof(uint256 root, uint256 signalHash, uint256 nullifierHash, uint256 externalNullifierHash, uint256[8] proof) view returns()
func (_WorldId *WorldIdCallerSession) VerifyProof(root *big.Int, signalHash *big.Int, nullifierHash *big.Int, externalNullifierHash *big.Int, proof [8]*big.Int) error {
	return _WorldId.Contract.VerifyProof(&_WorldId.CallOpts, root, signalHash, nullifierHash, externalNullifierHash, proof)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_WorldId *WorldIdTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorldId.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_WorldId *WorldIdSession) AcceptOwnership() (*types.Transaction, error) {
	return _WorldId.Contract.AcceptOwnership(&_WorldId.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_WorldId *WorldIdTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _WorldId.Contract.AcceptOwnership(&_WorldId.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x38c87065.
//
// Solidity: function initialize(uint8 _treeDepth, uint256 initialRoot, address _batchInsertionVerifiers, address _batchUpdateVerifiers, address _semaphoreVerifier) returns()
func (_WorldId *WorldIdTransactor) Initialize(opts *bind.TransactOpts, _treeDepth uint8, initialRoot *big.Int, _batchInsertionVerifiers common.Address, _batchUpdateVerifiers common.Address, _semaphoreVerifier common.Address) (*types.Transaction, error) {
	return _WorldId.contract.Transact(opts, "initialize", _treeDepth, initialRoot, _batchInsertionVerifiers, _batchUpdateVerifiers, _semaphoreVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x38c87065.
//
// Solidity: function initialize(uint8 _treeDepth, uint256 initialRoot, address _batchInsertionVerifiers, address _batchUpdateVerifiers, address _semaphoreVerifier) returns()
func (_WorldId *WorldIdSession) Initialize(_treeDepth uint8, initialRoot *big.Int, _batchInsertionVerifiers common.Address, _batchUpdateVerifiers common.Address, _semaphoreVerifier common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.Initialize(&_WorldId.TransactOpts, _treeDepth, initialRoot, _batchInsertionVerifiers, _batchUpdateVerifiers, _semaphoreVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x38c87065.
//
// Solidity: function initialize(uint8 _treeDepth, uint256 initialRoot, address _batchInsertionVerifiers, address _batchUpdateVerifiers, address _semaphoreVerifier) returns()
func (_WorldId *WorldIdTransactorSession) Initialize(_treeDepth uint8, initialRoot *big.Int, _batchInsertionVerifiers common.Address, _batchUpdateVerifiers common.Address, _semaphoreVerifier common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.Initialize(&_WorldId.TransactOpts, _treeDepth, initialRoot, _batchInsertionVerifiers, _batchUpdateVerifiers, _semaphoreVerifier)
}

// RegisterIdentities is a paid mutator transaction binding the contract method 0x2217b211.
//
// Solidity: function registerIdentities(uint256[8] insertionProof, uint256 preRoot, uint32 startIndex, uint256[] identityCommitments, uint256 postRoot) returns()
func (_WorldId *WorldIdTransactor) RegisterIdentities(opts *bind.TransactOpts, insertionProof [8]*big.Int, preRoot *big.Int, startIndex uint32, identityCommitments []*big.Int, postRoot *big.Int) (*types.Transaction, error) {
	return _WorldId.contract.Transact(opts, "registerIdentities", insertionProof, preRoot, startIndex, identityCommitments, postRoot)
}

// RegisterIdentities is a paid mutator transaction binding the contract method 0x2217b211.
//
// Solidity: function registerIdentities(uint256[8] insertionProof, uint256 preRoot, uint32 startIndex, uint256[] identityCommitments, uint256 postRoot) returns()
func (_WorldId *WorldIdSession) RegisterIdentities(insertionProof [8]*big.Int, preRoot *big.Int, startIndex uint32, identityCommitments []*big.Int, postRoot *big.Int) (*types.Transaction, error) {
	return _WorldId.Contract.RegisterIdentities(&_WorldId.TransactOpts, insertionProof, preRoot, startIndex, identityCommitments, postRoot)
}

// RegisterIdentities is a paid mutator transaction binding the contract method 0x2217b211.
//
// Solidity: function registerIdentities(uint256[8] insertionProof, uint256 preRoot, uint32 startIndex, uint256[] identityCommitments, uint256 postRoot) returns()
func (_WorldId *WorldIdTransactorSession) RegisterIdentities(insertionProof [8]*big.Int, preRoot *big.Int, startIndex uint32, identityCommitments []*big.Int, postRoot *big.Int) (*types.Transaction, error) {
	return _WorldId.Contract.RegisterIdentities(&_WorldId.TransactOpts, insertionProof, preRoot, startIndex, identityCommitments, postRoot)
}

// SetIdentityOperator is a paid mutator transaction binding the contract method 0xa7bba582.
//
// Solidity: function setIdentityOperator(address newIdentityOperator) returns(address)
func (_WorldId *WorldIdTransactor) SetIdentityOperator(opts *bind.TransactOpts, newIdentityOperator common.Address) (*types.Transaction, error) {
	return _WorldId.contract.Transact(opts, "setIdentityOperator", newIdentityOperator)
}

// SetIdentityOperator is a paid mutator transaction binding the contract method 0xa7bba582.
//
// Solidity: function setIdentityOperator(address newIdentityOperator) returns(address)
func (_WorldId *WorldIdSession) SetIdentityOperator(newIdentityOperator common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.SetIdentityOperator(&_WorldId.TransactOpts, newIdentityOperator)
}

// SetIdentityOperator is a paid mutator transaction binding the contract method 0xa7bba582.
//
// Solidity: function setIdentityOperator(address newIdentityOperator) returns(address)
func (_WorldId *WorldIdTransactorSession) SetIdentityOperator(newIdentityOperator common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.SetIdentityOperator(&_WorldId.TransactOpts, newIdentityOperator)
}

// SetRegisterIdentitiesVerifierLookupTable is a paid mutator transaction binding the contract method 0x2f059fca.
//
// Solidity: function setRegisterIdentitiesVerifierLookupTable(address newTable) returns()
func (_WorldId *WorldIdTransactor) SetRegisterIdentitiesVerifierLookupTable(opts *bind.TransactOpts, newTable common.Address) (*types.Transaction, error) {
	return _WorldId.contract.Transact(opts, "setRegisterIdentitiesVerifierLookupTable", newTable)
}

// SetRegisterIdentitiesVerifierLookupTable is a paid mutator transaction binding the contract method 0x2f059fca.
//
// Solidity: function setRegisterIdentitiesVerifierLookupTable(address newTable) returns()
func (_WorldId *WorldIdSession) SetRegisterIdentitiesVerifierLookupTable(newTable common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.SetRegisterIdentitiesVerifierLookupTable(&_WorldId.TransactOpts, newTable)
}

// SetRegisterIdentitiesVerifierLookupTable is a paid mutator transaction binding the contract method 0x2f059fca.
//
// Solidity: function setRegisterIdentitiesVerifierLookupTable(address newTable) returns()
func (_WorldId *WorldIdTransactorSession) SetRegisterIdentitiesVerifierLookupTable(newTable common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.SetRegisterIdentitiesVerifierLookupTable(&_WorldId.TransactOpts, newTable)
}

// SetRootHistoryExpiry is a paid mutator transaction binding the contract method 0xc70aa727.
//
// Solidity: function setRootHistoryExpiry(uint256 newExpiryTime) returns()
func (_WorldId *WorldIdTransactor) SetRootHistoryExpiry(opts *bind.TransactOpts, newExpiryTime *big.Int) (*types.Transaction, error) {
	return _WorldId.contract.Transact(opts, "setRootHistoryExpiry", newExpiryTime)
}

// SetRootHistoryExpiry is a paid mutator transaction binding the contract method 0xc70aa727.
//
// Solidity: function setRootHistoryExpiry(uint256 newExpiryTime) returns()
func (_WorldId *WorldIdSession) SetRootHistoryExpiry(newExpiryTime *big.Int) (*types.Transaction, error) {
	return _WorldId.Contract.SetRootHistoryExpiry(&_WorldId.TransactOpts, newExpiryTime)
}

// SetRootHistoryExpiry is a paid mutator transaction binding the contract method 0xc70aa727.
//
// Solidity: function setRootHistoryExpiry(uint256 newExpiryTime) returns()
func (_WorldId *WorldIdTransactorSession) SetRootHistoryExpiry(newExpiryTime *big.Int) (*types.Transaction, error) {
	return _WorldId.Contract.SetRootHistoryExpiry(&_WorldId.TransactOpts, newExpiryTime)
}

// SetSemaphoreVerifier is a paid mutator transaction binding the contract method 0x0e3a12f3.
//
// Solidity: function setSemaphoreVerifier(address newVerifier) returns()
func (_WorldId *WorldIdTransactor) SetSemaphoreVerifier(opts *bind.TransactOpts, newVerifier common.Address) (*types.Transaction, error) {
	return _WorldId.contract.Transact(opts, "setSemaphoreVerifier", newVerifier)
}

// SetSemaphoreVerifier is a paid mutator transaction binding the contract method 0x0e3a12f3.
//
// Solidity: function setSemaphoreVerifier(address newVerifier) returns()
func (_WorldId *WorldIdSession) SetSemaphoreVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.SetSemaphoreVerifier(&_WorldId.TransactOpts, newVerifier)
}

// SetSemaphoreVerifier is a paid mutator transaction binding the contract method 0x0e3a12f3.
//
// Solidity: function setSemaphoreVerifier(address newVerifier) returns()
func (_WorldId *WorldIdTransactorSession) SetSemaphoreVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.SetSemaphoreVerifier(&_WorldId.TransactOpts, newVerifier)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorldId *WorldIdTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WorldId.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorldId *WorldIdSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.TransferOwnership(&_WorldId.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorldId *WorldIdTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.TransferOwnership(&_WorldId.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WorldId *WorldIdTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _WorldId.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WorldId *WorldIdSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.UpgradeTo(&_WorldId.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WorldId *WorldIdTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WorldId.Contract.UpgradeTo(&_WorldId.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WorldId *WorldIdTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WorldId.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WorldId *WorldIdSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WorldId.Contract.UpgradeToAndCall(&_WorldId.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WorldId *WorldIdTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WorldId.Contract.UpgradeToAndCall(&_WorldId.TransactOpts, newImplementation, data)
}

// WorldIdAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the WorldId contract.
type WorldIdAdminChangedIterator struct {
	Event *WorldIdAdminChanged // Event containing the contract specifics and raw log

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
func (it *WorldIdAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdAdminChanged)
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
		it.Event = new(WorldIdAdminChanged)
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
func (it *WorldIdAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdAdminChanged represents a AdminChanged event raised by the WorldId contract.
type WorldIdAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WorldId *WorldIdFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*WorldIdAdminChangedIterator, error) {

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &WorldIdAdminChangedIterator{contract: _WorldId.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WorldId *WorldIdFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *WorldIdAdminChanged) (event.Subscription, error) {

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdAdminChanged)
				if err := _WorldId.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WorldId *WorldIdFilterer) ParseAdminChanged(log types.Log) (*WorldIdAdminChanged, error) {
	event := new(WorldIdAdminChanged)
	if err := _WorldId.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldIdBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the WorldId contract.
type WorldIdBeaconUpgradedIterator struct {
	Event *WorldIdBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *WorldIdBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdBeaconUpgraded)
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
		it.Event = new(WorldIdBeaconUpgraded)
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
func (it *WorldIdBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdBeaconUpgraded represents a BeaconUpgraded event raised by the WorldId contract.
type WorldIdBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WorldId *WorldIdFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*WorldIdBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &WorldIdBeaconUpgradedIterator{contract: _WorldId.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WorldId *WorldIdFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *WorldIdBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdBeaconUpgraded)
				if err := _WorldId.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WorldId *WorldIdFilterer) ParseBeaconUpgraded(log types.Log) (*WorldIdBeaconUpgraded, error) {
	event := new(WorldIdBeaconUpgraded)
	if err := _WorldId.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldIdDependencyUpdatedIterator is returned from FilterDependencyUpdated and is used to iterate over the raw logs and unpacked data for DependencyUpdated events raised by the WorldId contract.
type WorldIdDependencyUpdatedIterator struct {
	Event *WorldIdDependencyUpdated // Event containing the contract specifics and raw log

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
func (it *WorldIdDependencyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdDependencyUpdated)
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
		it.Event = new(WorldIdDependencyUpdated)
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
func (it *WorldIdDependencyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdDependencyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdDependencyUpdated represents a DependencyUpdated event raised by the WorldId contract.
type WorldIdDependencyUpdated struct {
	Kind       uint8
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDependencyUpdated is a free log retrieval operation binding the contract event 0xd194b8423e9cb3c7cbebbbc3fe7f79dc2cbe0b40e03270d975abff491504c7b1.
//
// Solidity: event DependencyUpdated(uint8 indexed kind, address indexed oldAddress, address indexed newAddress)
func (_WorldId *WorldIdFilterer) FilterDependencyUpdated(opts *bind.FilterOpts, kind []uint8, oldAddress []common.Address, newAddress []common.Address) (*WorldIdDependencyUpdatedIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "DependencyUpdated", kindRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &WorldIdDependencyUpdatedIterator{contract: _WorldId.contract, event: "DependencyUpdated", logs: logs, sub: sub}, nil
}

// WatchDependencyUpdated is a free log subscription operation binding the contract event 0xd194b8423e9cb3c7cbebbbc3fe7f79dc2cbe0b40e03270d975abff491504c7b1.
//
// Solidity: event DependencyUpdated(uint8 indexed kind, address indexed oldAddress, address indexed newAddress)
func (_WorldId *WorldIdFilterer) WatchDependencyUpdated(opts *bind.WatchOpts, sink chan<- *WorldIdDependencyUpdated, kind []uint8, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "DependencyUpdated", kindRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdDependencyUpdated)
				if err := _WorldId.contract.UnpackLog(event, "DependencyUpdated", log); err != nil {
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

// ParseDependencyUpdated is a log parse operation binding the contract event 0xd194b8423e9cb3c7cbebbbc3fe7f79dc2cbe0b40e03270d975abff491504c7b1.
//
// Solidity: event DependencyUpdated(uint8 indexed kind, address indexed oldAddress, address indexed newAddress)
func (_WorldId *WorldIdFilterer) ParseDependencyUpdated(log types.Log) (*WorldIdDependencyUpdated, error) {
	event := new(WorldIdDependencyUpdated)
	if err := _WorldId.contract.UnpackLog(event, "DependencyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldIdIdentityOperatorChangedIterator is returned from FilterIdentityOperatorChanged and is used to iterate over the raw logs and unpacked data for IdentityOperatorChanged events raised by the WorldId contract.
type WorldIdIdentityOperatorChangedIterator struct {
	Event *WorldIdIdentityOperatorChanged // Event containing the contract specifics and raw log

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
func (it *WorldIdIdentityOperatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdIdentityOperatorChanged)
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
		it.Event = new(WorldIdIdentityOperatorChanged)
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
func (it *WorldIdIdentityOperatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdIdentityOperatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdIdentityOperatorChanged represents a IdentityOperatorChanged event raised by the WorldId contract.
type WorldIdIdentityOperatorChanged struct {
	OldOperator common.Address
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterIdentityOperatorChanged is a free log retrieval operation binding the contract event 0x5a674c516c196404869e5f502b5634ce442416bb016dde54b5de4c812cc019e6.
//
// Solidity: event IdentityOperatorChanged(address indexed oldOperator, address indexed newOperator)
func (_WorldId *WorldIdFilterer) FilterIdentityOperatorChanged(opts *bind.FilterOpts, oldOperator []common.Address, newOperator []common.Address) (*WorldIdIdentityOperatorChangedIterator, error) {

	var oldOperatorRule []interface{}
	for _, oldOperatorItem := range oldOperator {
		oldOperatorRule = append(oldOperatorRule, oldOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "IdentityOperatorChanged", oldOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return &WorldIdIdentityOperatorChangedIterator{contract: _WorldId.contract, event: "IdentityOperatorChanged", logs: logs, sub: sub}, nil
}

// WatchIdentityOperatorChanged is a free log subscription operation binding the contract event 0x5a674c516c196404869e5f502b5634ce442416bb016dde54b5de4c812cc019e6.
//
// Solidity: event IdentityOperatorChanged(address indexed oldOperator, address indexed newOperator)
func (_WorldId *WorldIdFilterer) WatchIdentityOperatorChanged(opts *bind.WatchOpts, sink chan<- *WorldIdIdentityOperatorChanged, oldOperator []common.Address, newOperator []common.Address) (event.Subscription, error) {

	var oldOperatorRule []interface{}
	for _, oldOperatorItem := range oldOperator {
		oldOperatorRule = append(oldOperatorRule, oldOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "IdentityOperatorChanged", oldOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdIdentityOperatorChanged)
				if err := _WorldId.contract.UnpackLog(event, "IdentityOperatorChanged", log); err != nil {
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

// ParseIdentityOperatorChanged is a log parse operation binding the contract event 0x5a674c516c196404869e5f502b5634ce442416bb016dde54b5de4c812cc019e6.
//
// Solidity: event IdentityOperatorChanged(address indexed oldOperator, address indexed newOperator)
func (_WorldId *WorldIdFilterer) ParseIdentityOperatorChanged(log types.Log) (*WorldIdIdentityOperatorChanged, error) {
	event := new(WorldIdIdentityOperatorChanged)
	if err := _WorldId.contract.UnpackLog(event, "IdentityOperatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldIdInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WorldId contract.
type WorldIdInitializedIterator struct {
	Event *WorldIdInitialized // Event containing the contract specifics and raw log

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
func (it *WorldIdInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdInitialized)
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
		it.Event = new(WorldIdInitialized)
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
func (it *WorldIdInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdInitialized represents a Initialized event raised by the WorldId contract.
type WorldIdInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorldId *WorldIdFilterer) FilterInitialized(opts *bind.FilterOpts) (*WorldIdInitializedIterator, error) {

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WorldIdInitializedIterator{contract: _WorldId.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorldId *WorldIdFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WorldIdInitialized) (event.Subscription, error) {

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdInitialized)
				if err := _WorldId.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorldId *WorldIdFilterer) ParseInitialized(log types.Log) (*WorldIdInitialized, error) {
	event := new(WorldIdInitialized)
	if err := _WorldId.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldIdOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the WorldId contract.
type WorldIdOwnershipTransferStartedIterator struct {
	Event *WorldIdOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *WorldIdOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdOwnershipTransferStarted)
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
		it.Event = new(WorldIdOwnershipTransferStarted)
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
func (it *WorldIdOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the WorldId contract.
type WorldIdOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_WorldId *WorldIdFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WorldIdOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WorldIdOwnershipTransferStartedIterator{contract: _WorldId.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_WorldId *WorldIdFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *WorldIdOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdOwnershipTransferStarted)
				if err := _WorldId.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_WorldId *WorldIdFilterer) ParseOwnershipTransferStarted(log types.Log) (*WorldIdOwnershipTransferStarted, error) {
	event := new(WorldIdOwnershipTransferStarted)
	if err := _WorldId.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldIdOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WorldId contract.
type WorldIdOwnershipTransferredIterator struct {
	Event *WorldIdOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WorldIdOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdOwnershipTransferred)
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
		it.Event = new(WorldIdOwnershipTransferred)
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
func (it *WorldIdOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdOwnershipTransferred represents a OwnershipTransferred event raised by the WorldId contract.
type WorldIdOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorldId *WorldIdFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WorldIdOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WorldIdOwnershipTransferredIterator{contract: _WorldId.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorldId *WorldIdFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorldIdOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdOwnershipTransferred)
				if err := _WorldId.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorldId *WorldIdFilterer) ParseOwnershipTransferred(log types.Log) (*WorldIdOwnershipTransferred, error) {
	event := new(WorldIdOwnershipTransferred)
	if err := _WorldId.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldIdRootHistoryExpirySetIterator is returned from FilterRootHistoryExpirySet and is used to iterate over the raw logs and unpacked data for RootHistoryExpirySet events raised by the WorldId contract.
type WorldIdRootHistoryExpirySetIterator struct {
	Event *WorldIdRootHistoryExpirySet // Event containing the contract specifics and raw log

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
func (it *WorldIdRootHistoryExpirySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdRootHistoryExpirySet)
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
		it.Event = new(WorldIdRootHistoryExpirySet)
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
func (it *WorldIdRootHistoryExpirySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdRootHistoryExpirySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdRootHistoryExpirySet represents a RootHistoryExpirySet event raised by the WorldId contract.
type WorldIdRootHistoryExpirySet struct {
	OldExpiryTime *big.Int
	NewExpiryTime *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRootHistoryExpirySet is a free log retrieval operation binding the contract event 0xf62a6f06fde00a303cd5939e6b53762854412c96a196cda26720cedd28af9e70.
//
// Solidity: event RootHistoryExpirySet(uint256 indexed oldExpiryTime, uint256 indexed newExpiryTime)
func (_WorldId *WorldIdFilterer) FilterRootHistoryExpirySet(opts *bind.FilterOpts, oldExpiryTime []*big.Int, newExpiryTime []*big.Int) (*WorldIdRootHistoryExpirySetIterator, error) {

	var oldExpiryTimeRule []interface{}
	for _, oldExpiryTimeItem := range oldExpiryTime {
		oldExpiryTimeRule = append(oldExpiryTimeRule, oldExpiryTimeItem)
	}
	var newExpiryTimeRule []interface{}
	for _, newExpiryTimeItem := range newExpiryTime {
		newExpiryTimeRule = append(newExpiryTimeRule, newExpiryTimeItem)
	}

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "RootHistoryExpirySet", oldExpiryTimeRule, newExpiryTimeRule)
	if err != nil {
		return nil, err
	}
	return &WorldIdRootHistoryExpirySetIterator{contract: _WorldId.contract, event: "RootHistoryExpirySet", logs: logs, sub: sub}, nil
}

// WatchRootHistoryExpirySet is a free log subscription operation binding the contract event 0xf62a6f06fde00a303cd5939e6b53762854412c96a196cda26720cedd28af9e70.
//
// Solidity: event RootHistoryExpirySet(uint256 indexed oldExpiryTime, uint256 indexed newExpiryTime)
func (_WorldId *WorldIdFilterer) WatchRootHistoryExpirySet(opts *bind.WatchOpts, sink chan<- *WorldIdRootHistoryExpirySet, oldExpiryTime []*big.Int, newExpiryTime []*big.Int) (event.Subscription, error) {

	var oldExpiryTimeRule []interface{}
	for _, oldExpiryTimeItem := range oldExpiryTime {
		oldExpiryTimeRule = append(oldExpiryTimeRule, oldExpiryTimeItem)
	}
	var newExpiryTimeRule []interface{}
	for _, newExpiryTimeItem := range newExpiryTime {
		newExpiryTimeRule = append(newExpiryTimeRule, newExpiryTimeItem)
	}

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "RootHistoryExpirySet", oldExpiryTimeRule, newExpiryTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdRootHistoryExpirySet)
				if err := _WorldId.contract.UnpackLog(event, "RootHistoryExpirySet", log); err != nil {
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

// ParseRootHistoryExpirySet is a log parse operation binding the contract event 0xf62a6f06fde00a303cd5939e6b53762854412c96a196cda26720cedd28af9e70.
//
// Solidity: event RootHistoryExpirySet(uint256 indexed oldExpiryTime, uint256 indexed newExpiryTime)
func (_WorldId *WorldIdFilterer) ParseRootHistoryExpirySet(log types.Log) (*WorldIdRootHistoryExpirySet, error) {
	event := new(WorldIdRootHistoryExpirySet)
	if err := _WorldId.contract.UnpackLog(event, "RootHistoryExpirySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldIdStateBridgeStateChangeIterator is returned from FilterStateBridgeStateChange and is used to iterate over the raw logs and unpacked data for StateBridgeStateChange events raised by the WorldId contract.
type WorldIdStateBridgeStateChangeIterator struct {
	Event *WorldIdStateBridgeStateChange // Event containing the contract specifics and raw log

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
func (it *WorldIdStateBridgeStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdStateBridgeStateChange)
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
		it.Event = new(WorldIdStateBridgeStateChange)
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
func (it *WorldIdStateBridgeStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdStateBridgeStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdStateBridgeStateChange represents a StateBridgeStateChange event raised by the WorldId contract.
type WorldIdStateBridgeStateChange struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStateBridgeStateChange is a free log retrieval operation binding the contract event 0xc2481c150ab648d8248c5a0b452a12da07a6a26d8e11b0ed6cee10375e14d5e4.
//
// Solidity: event StateBridgeStateChange(bool indexed isEnabled)
func (_WorldId *WorldIdFilterer) FilterStateBridgeStateChange(opts *bind.FilterOpts, isEnabled []bool) (*WorldIdStateBridgeStateChangeIterator, error) {

	var isEnabledRule []interface{}
	for _, isEnabledItem := range isEnabled {
		isEnabledRule = append(isEnabledRule, isEnabledItem)
	}

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "StateBridgeStateChange", isEnabledRule)
	if err != nil {
		return nil, err
	}
	return &WorldIdStateBridgeStateChangeIterator{contract: _WorldId.contract, event: "StateBridgeStateChange", logs: logs, sub: sub}, nil
}

// WatchStateBridgeStateChange is a free log subscription operation binding the contract event 0xc2481c150ab648d8248c5a0b452a12da07a6a26d8e11b0ed6cee10375e14d5e4.
//
// Solidity: event StateBridgeStateChange(bool indexed isEnabled)
func (_WorldId *WorldIdFilterer) WatchStateBridgeStateChange(opts *bind.WatchOpts, sink chan<- *WorldIdStateBridgeStateChange, isEnabled []bool) (event.Subscription, error) {

	var isEnabledRule []interface{}
	for _, isEnabledItem := range isEnabled {
		isEnabledRule = append(isEnabledRule, isEnabledItem)
	}

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "StateBridgeStateChange", isEnabledRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdStateBridgeStateChange)
				if err := _WorldId.contract.UnpackLog(event, "StateBridgeStateChange", log); err != nil {
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

// ParseStateBridgeStateChange is a log parse operation binding the contract event 0xc2481c150ab648d8248c5a0b452a12da07a6a26d8e11b0ed6cee10375e14d5e4.
//
// Solidity: event StateBridgeStateChange(bool indexed isEnabled)
func (_WorldId *WorldIdFilterer) ParseStateBridgeStateChange(log types.Log) (*WorldIdStateBridgeStateChange, error) {
	event := new(WorldIdStateBridgeStateChange)
	if err := _WorldId.contract.UnpackLog(event, "StateBridgeStateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldIdTreeChangedIterator is returned from FilterTreeChanged and is used to iterate over the raw logs and unpacked data for TreeChanged events raised by the WorldId contract.
type WorldIdTreeChangedIterator struct {
	Event *WorldIdTreeChanged // Event containing the contract specifics and raw log

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
func (it *WorldIdTreeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdTreeChanged)
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
		it.Event = new(WorldIdTreeChanged)
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
func (it *WorldIdTreeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdTreeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdTreeChanged represents a TreeChanged event raised by the WorldId contract.
type WorldIdTreeChanged struct {
	PreRoot  *big.Int
	Kind     uint8
	PostRoot *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTreeChanged is a free log retrieval operation binding the contract event 0x25f6d5cc356ee0b49cf708c13c68197947f5740a878a298765e4b18e4afdaf04.
//
// Solidity: event TreeChanged(uint256 indexed preRoot, uint8 indexed kind, uint256 indexed postRoot)
func (_WorldId *WorldIdFilterer) FilterTreeChanged(opts *bind.FilterOpts, preRoot []*big.Int, kind []uint8, postRoot []*big.Int) (*WorldIdTreeChangedIterator, error) {

	var preRootRule []interface{}
	for _, preRootItem := range preRoot {
		preRootRule = append(preRootRule, preRootItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var postRootRule []interface{}
	for _, postRootItem := range postRoot {
		postRootRule = append(postRootRule, postRootItem)
	}

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "TreeChanged", preRootRule, kindRule, postRootRule)
	if err != nil {
		return nil, err
	}
	return &WorldIdTreeChangedIterator{contract: _WorldId.contract, event: "TreeChanged", logs: logs, sub: sub}, nil
}

// WatchTreeChanged is a free log subscription operation binding the contract event 0x25f6d5cc356ee0b49cf708c13c68197947f5740a878a298765e4b18e4afdaf04.
//
// Solidity: event TreeChanged(uint256 indexed preRoot, uint8 indexed kind, uint256 indexed postRoot)
func (_WorldId *WorldIdFilterer) WatchTreeChanged(opts *bind.WatchOpts, sink chan<- *WorldIdTreeChanged, preRoot []*big.Int, kind []uint8, postRoot []*big.Int) (event.Subscription, error) {

	var preRootRule []interface{}
	for _, preRootItem := range preRoot {
		preRootRule = append(preRootRule, preRootItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}
	var postRootRule []interface{}
	for _, postRootItem := range postRoot {
		postRootRule = append(postRootRule, postRootItem)
	}

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "TreeChanged", preRootRule, kindRule, postRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdTreeChanged)
				if err := _WorldId.contract.UnpackLog(event, "TreeChanged", log); err != nil {
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

// ParseTreeChanged is a log parse operation binding the contract event 0x25f6d5cc356ee0b49cf708c13c68197947f5740a878a298765e4b18e4afdaf04.
//
// Solidity: event TreeChanged(uint256 indexed preRoot, uint8 indexed kind, uint256 indexed postRoot)
func (_WorldId *WorldIdFilterer) ParseTreeChanged(log types.Log) (*WorldIdTreeChanged, error) {
	event := new(WorldIdTreeChanged)
	if err := _WorldId.contract.UnpackLog(event, "TreeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldIdUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the WorldId contract.
type WorldIdUpgradedIterator struct {
	Event *WorldIdUpgraded // Event containing the contract specifics and raw log

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
func (it *WorldIdUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdUpgraded)
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
		it.Event = new(WorldIdUpgraded)
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
func (it *WorldIdUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdUpgraded represents a Upgraded event raised by the WorldId contract.
type WorldIdUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WorldId *WorldIdFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WorldIdUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WorldIdUpgradedIterator{contract: _WorldId.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WorldId *WorldIdFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WorldIdUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdUpgraded)
				if err := _WorldId.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WorldId *WorldIdFilterer) ParseUpgraded(log types.Log) (*WorldIdUpgraded, error) {
	event := new(WorldIdUpgraded)
	if err := _WorldId.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorldIdWorldIDIdentityManagerImplInitializedIterator is returned from FilterWorldIDIdentityManagerImplInitialized and is used to iterate over the raw logs and unpacked data for WorldIDIdentityManagerImplInitialized events raised by the WorldId contract.
type WorldIdWorldIDIdentityManagerImplInitializedIterator struct {
	Event *WorldIdWorldIDIdentityManagerImplInitialized // Event containing the contract specifics and raw log

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
func (it *WorldIdWorldIDIdentityManagerImplInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorldIdWorldIDIdentityManagerImplInitialized)
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
		it.Event = new(WorldIdWorldIDIdentityManagerImplInitialized)
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
func (it *WorldIdWorldIDIdentityManagerImplInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorldIdWorldIDIdentityManagerImplInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorldIdWorldIDIdentityManagerImplInitialized represents a WorldIDIdentityManagerImplInitialized event raised by the WorldId contract.
type WorldIdWorldIDIdentityManagerImplInitialized struct {
	TreeDepth   uint8
	InitialRoot *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWorldIDIdentityManagerImplInitialized is a free log retrieval operation binding the contract event 0xd1bcc66c061c32a21f569d138c2dadef4a38a0636309881954af5f44010ec1d8.
//
// Solidity: event WorldIDIdentityManagerImplInitialized(uint8 _treeDepth, uint256 initialRoot)
func (_WorldId *WorldIdFilterer) FilterWorldIDIdentityManagerImplInitialized(opts *bind.FilterOpts) (*WorldIdWorldIDIdentityManagerImplInitializedIterator, error) {

	logs, sub, err := _WorldId.contract.FilterLogs(opts, "WorldIDIdentityManagerImplInitialized")
	if err != nil {
		return nil, err
	}
	return &WorldIdWorldIDIdentityManagerImplInitializedIterator{contract: _WorldId.contract, event: "WorldIDIdentityManagerImplInitialized", logs: logs, sub: sub}, nil
}

// WatchWorldIDIdentityManagerImplInitialized is a free log subscription operation binding the contract event 0xd1bcc66c061c32a21f569d138c2dadef4a38a0636309881954af5f44010ec1d8.
//
// Solidity: event WorldIDIdentityManagerImplInitialized(uint8 _treeDepth, uint256 initialRoot)
func (_WorldId *WorldIdFilterer) WatchWorldIDIdentityManagerImplInitialized(opts *bind.WatchOpts, sink chan<- *WorldIdWorldIDIdentityManagerImplInitialized) (event.Subscription, error) {

	logs, sub, err := _WorldId.contract.WatchLogs(opts, "WorldIDIdentityManagerImplInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorldIdWorldIDIdentityManagerImplInitialized)
				if err := _WorldId.contract.UnpackLog(event, "WorldIDIdentityManagerImplInitialized", log); err != nil {
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

// ParseWorldIDIdentityManagerImplInitialized is a log parse operation binding the contract event 0xd1bcc66c061c32a21f569d138c2dadef4a38a0636309881954af5f44010ec1d8.
//
// Solidity: event WorldIDIdentityManagerImplInitialized(uint8 _treeDepth, uint256 initialRoot)
func (_WorldId *WorldIdFilterer) ParseWorldIDIdentityManagerImplInitialized(log types.Log) (*WorldIdWorldIDIdentityManagerImplInitialized, error) {
	event := new(WorldIdWorldIDIdentityManagerImplInitialized)
	if err := _WorldId.contract.UnpackLog(event, "WorldIDIdentityManagerImplInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
