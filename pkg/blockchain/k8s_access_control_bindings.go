// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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

// K8sAccessControlMetaData contains all meta data concerning the K8sAccessControl contract.
var K8sAccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"namespaces\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"verbs\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"resources\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"}],\"name\":\"PermissionGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"PermissionRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"namespaces\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"verbs\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"resources\",\"type\":\"string[]\"}],\"name\":\"PermissionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSION_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPermissionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAuthorizedWalletAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizedWalletsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"getPermissions\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"namespaces\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"verbs\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"resources\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"namespaces\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"verbs\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"resources\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"}],\"name\":\"grantPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"namespace\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verb\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePermissionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"revokePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"namespaces\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"verbs\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"resources\",\"type\":\"string[]\"}],\"name\":\"updatePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001f60003362000080565b506200004c7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217753362000080565b50620000797f50dd73a7b7a643432ced584fcc2cd0a1d5efa85899b62d28c5a52294d183d0e33362000080565b506200012f565b6000828152602081815260408083206001600160a01b038516845290915281205460ff1662000125576000838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055620000dc3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600162000129565b5060005b92915050565b61202f806200013f6000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806375b238fc116100ad578063d547741f11610071578063d547741f14610275578063dd4c93b414610288578063e08951ea1461029b578063ecb80b6f146102c6578063fe9fbb80146102d957600080fd5b806375b238fc1461021857806375f6db4e1461023f57806391d1485414610247578063a217fddf1461025a578063a3bf5b9a1461026257600080fd5b806327f1c1a7116100f457806327f1c1a7146101b75780632f2ff15d146101ca57806336568abe146101df5780635958abb0146101f25780636811ba2f1461020557600080fd5b80630125a4251461012657806301ffc9a71461014e578063160a792514610171578063248a9ca314610194575b600080fd5b61013b600080516020611fda83398151915281565b6040519081526020015b60405180910390f35b61016161015c36600461194d565b6102ec565b6040519015158152602001610145565b61018461017f366004611993565b610323565b6040516101459493929190611a33565b61013b6101a2366004611a7e565b60009081526020819052604090206001015490565b6101616101c5366004611b4e565b6106b2565b6101dd6101d8366004611be7565b610aa4565b005b6101dd6101ed366004611be7565b610acf565b6101dd610200366004611993565b610b07565b6101dd610213366004611cb2565b610b4d565b61013b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b61013b610d9e565b610161610255366004611be7565b610daf565b61013b600081565b6101dd610270366004611993565b610dd8565b6101dd610283366004611be7565b610ec2565b6101dd610296366004611d3f565b610ee7565b6102ae6102a9366004611a7e565b6111a1565b6040516001600160a01b039091168152602001610145565b6101dd6102d4366004611993565b6111fc565b6101616102e7366004611993565b61123e565b60006001600160e01b03198216637965db0b60e01b148061031d57506301ffc9a760e01b6001600160e01b03198316145b92915050565b6001600160a01b03811660009081526001602090815260408083208151815460c09481028201850190935260a0810183815260609586958695919485949390928492849190879085015b8282101561041957838290600052602060002001805461038c90611de0565b80601f01602080910402602001604051908101604052809291908181526020018280546103b890611de0565b80156104055780601f106103da57610100808354040283529160200191610405565b820191906000526020600020905b8154815290600101906020018083116103e857829003601f168201915b50505050508152602001906001019061036d565b50505050815260200160018201805480602002602001604051908101604052809291908181526020016000905b828210156104f257838290600052602060002001805461046590611de0565b80601f016020809104026020016040519081016040528092919081815260200182805461049190611de0565b80156104de5780601f106104b3576101008083540402835291602001916104de565b820191906000526020600020905b8154815290600101906020018083116104c157829003601f168201915b505050505081526020019060010190610446565b50505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b828210156105cb57838290600052602060002001805461053e90611de0565b80601f016020809104026020016040519081016040528092919081815260200182805461056a90611de0565b80156105b75780601f1061058c576101008083540402835291602001916105b7565b820191906000526020600020905b81548152906001019060200180831161059a57829003601f168201915b50505050508152602001906001019061051f565b505050908252506003820154602082015260049091015460ff161515604090910152608081015190915061068d576040805160008082526020820190925290610624565b606081526020019060019003908161060f5790505b506040805160008082526020820190925290610650565b606081526020019060019003908161063b5790505b50604080516000808252602082019092529061067c565b60608152602001906001900390816106675790505b5060009450945094509450506106ab565b80600001518160200151826040015183606001519450945094509450505b9193509193565b6001600160a01b03841660009081526001602090815260408083208151815460c09481028201850190935260a08101838152859491938492849190879085015b8282101561079e57838290600052602060002001805461071190611de0565b80601f016020809104026020016040519081016040528092919081815260200182805461073d90611de0565b801561078a5780601f1061075f5761010080835404028352916020019161078a565b820191906000526020600020905b81548152906001019060200180831161076d57829003601f168201915b5050505050815260200190600101906106f2565b50505050815260200160018201805480602002602001604051908101604052809291908181526020016000905b828210156108775783829060005260206000200180546107ea90611de0565b80601f016020809104026020016040519081016040528092919081815260200182805461081690611de0565b80156108635780601f1061083857610100808354040283529160200191610863565b820191906000526020600020905b81548152906001019060200180831161084657829003601f168201915b5050505050815260200190600101906107cb565b50505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b828210156109505783829060005260206000200180546108c390611de0565b80601f01602080910402602001604051908101604052809291908181526020018280546108ef90611de0565b801561093c5780601f106109115761010080835404028352916020019161093c565b820191906000526020600020905b81548152906001019060200180831161091f57829003601f168201915b5050505050815260200190600101906108a4565b505050908252506003820154602082015260049091015460ff1615156040909101526080810151909150610988576000915050610a9c565b6000816060015111801561099f5750806060015142115b156109ae576000915050610a9c565b80516109ba9086611541565b1580156109eb57506109e98160000151604051806040016040528060018152602001601560f91b815250611541565b155b156109fa576000915050610a9c565b610a08816020015185611541565b158015610a395750610a378160200151604051806040016040528060018152602001601560f91b815250611541565b155b15610a48576000915050610a9c565b610a56816040015184611541565b158015610a875750610a858160400151604051806040016040528060018152602001601560f91b815250611541565b155b15610a96576000915050610a9c565b60019150505b949350505050565b600082815260208190526040902060010154610abf816115a2565b610ac983836115af565b50505050565b6001600160a01b0381163314610af85760405163334bd91960e11b815260040160405180910390fd5b610b028282611641565b505050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775610b31816115a2565b610b49600080516020611fda83398151915283610aa4565b5050565b600080516020611fda833981519152610b65816115a2565b6001600160a01b03851660009081526001602052604090206004015460ff16610bd55760405162461bcd60e51b815260206004820181905260248201527f57616c6c657420686173206e6f20616374697665207065726d697373696f6e7360448201526064015b60405180910390fd5b6000845111610c265760405162461bcd60e51b815260206004820152601f60248201527f4174206c65617374206f6e65206e616d657370616365207265717569726564006044820152606401610bcc565b6000835111610c775760405162461bcd60e51b815260206004820152601a60248201527f4174206c65617374206f6e6520766572622072657175697265640000000000006044820152606401610bcc565b6000825111610cc85760405162461bcd60e51b815260206004820152601e60248201527f4174206c65617374206f6e65207265736f7572636520726571756972656400006044820152606401610bcc565b6001600160a01b03851660009081526001602090815260409091208551610cf192870190611890565b506001600160a01b03851660009081526001602081815260409092208551610d2193919092019190860190611890565b506001600160a01b03851660009081526001602090815260409091208351610d5192600290920191850190611890565b50846001600160a01b03167f259165b83cff236056b54312412ad942d4af8eddf7d3e92032bff6718465f3ba858585604051610d8f93929190611e1a565b60405180910390a25050505050565b6000610daa60026116ac565b905090565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b600080516020611fda833981519152610df0816115a2565b6001600160a01b03821660009081526001602052604090206004015460ff16610e5b5760405162461bcd60e51b815260206004820181905260248201527f57616c6c657420686173206e6f20616374697665207065726d697373696f6e736044820152606401610bcc565b6001600160a01b0382166000908152600160205260409020600401805460ff19169055610e896002836116b6565b506040516001600160a01b038316907f3541f93cbae8c4be65491b824efe1570976e740b18c6aa441db5291f4de4c92190600090a25050565b600082815260208190526040902060010154610edd816115a2565b610ac98383611641565b600080516020611fda833981519152610eff816115a2565b6001600160a01b038616610f4e5760405162461bcd60e51b8152602060048201526016602482015275496e76616c69642077616c6c6574206164647265737360501b6044820152606401610bcc565b6000855111610f9f5760405162461bcd60e51b815260206004820152601f60248201527f4174206c65617374206f6e65206e616d657370616365207265717569726564006044820152606401610bcc565b6000845111610ff05760405162461bcd60e51b815260206004820152601a60248201527f4174206c65617374206f6e6520766572622072657175697265640000000000006044820152606401610bcc565b60008351116110415760405162461bcd60e51b815260206004820152601e60248201527f4174206c65617374206f6e65207265736f7572636520726571756972656400006044820152606401610bcc565b8115611096574282116110965760405162461bcd60e51b815260206004820181905260248201527f45787069726174696f6e206d75737420626520696e20746865206675747572656044820152606401610bcc565b6040805160a0810182528681526020808201879052818301869052606082018590526001608083018190526001600160a01b038a1660009081529082529290922081518051929391926110ec9284920190611890565b5060208281015180516111059260018501920190611890565b5060408201518051611121916002840191602090910190611890565b50606082015160038201556080909101516004909101805460ff19169115159190911790556111516002876116d2565b50856001600160a01b03167f4a58c34854ef823594749e5d55b169493cf1e79b11ba2567eb0053addadf5aa4868686866040516111919493929190611a33565b60405180910390a2505050505050565b60006111ad60026116ac565b82106111f15760405162461bcd60e51b8152602060048201526013602482015272496e646578206f7574206f6620626f756e647360681b6044820152606401610bcc565b61031d6002836116e7565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775611226816115a2565b610b49600080516020611fda83398151915283610ec2565b6001600160a01b03811660009081526001602090815260408083208151815460c09481028201850190935260a08101838152859491938492849190879085015b8282101561132a57838290600052602060002001805461129d90611de0565b80601f01602080910402602001604051908101604052809291908181526020018280546112c990611de0565b80156113165780601f106112eb57610100808354040283529160200191611316565b820191906000526020600020905b8154815290600101906020018083116112f957829003601f168201915b50505050508152602001906001019061127e565b50505050815260200160018201805480602002602001604051908101604052809291908181526020016000905b8282101561140357838290600052602060002001805461137690611de0565b80601f01602080910402602001604051908101604052809291908181526020018280546113a290611de0565b80156113ef5780601f106113c4576101008083540402835291602001916113ef565b820191906000526020600020905b8154815290600101906020018083116113d257829003601f168201915b505050505081526020019060010190611357565b50505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b828210156114dc57838290600052602060002001805461144f90611de0565b80601f016020809104026020016040519081016040528092919081815260200182805461147b90611de0565b80156114c85780601f1061149d576101008083540402835291602001916114c8565b820191906000526020600020905b8154815290600101906020018083116114ab57829003601f168201915b505050505081526020019060010190611430565b505050908252506003820154602082015260049091015460ff16151560409091015260808101519091506115135750600092915050565b6000816060015111801561152a5750806060015142115b156115385750600092915050565b50600192915050565b6000805b835181101561159857828051906020012084828151811061156857611568611e5d565b6020026020010151805190602001200361158657600191505061031d565b8061159081611e89565b915050611545565b5060009392505050565b6115ac81336116f3565b50565b60006115bb8383610daf565b611639576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556115f13390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161031d565b50600061031d565b600061164d8383610daf565b15611639576000838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161031d565b600061031d825490565b60006116cb836001600160a01b03841661172c565b9392505050565b60006116cb836001600160a01b03841661181f565b60006116cb8383611866565b6116fd8282610daf565b610b495760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610bcc565b60008181526001830160205260408120548015611815576000611750600183611ea2565b855490915060009061176490600190611ea2565b90508082146117c957600086600001828154811061178457611784611e5d565b90600052602060002001549050808760000184815481106117a7576117a7611e5d565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806117da576117da611eb5565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061031d565b600091505061031d565b60008181526001830160205260408120546116395750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561031d565b600082600001828154811061187d5761187d611e5d565b9060005260206000200154905092915050565b8280548282559060005260206000209081019282156118d6579160200282015b828111156118d657825182906118c69082611f19565b50916020019190600101906118b0565b506118e29291506118e6565b5090565b808211156118e25760006118fa8282611903565b506001016118e6565b50805461190f90611de0565b6000825580601f1061191f575050565b601f0160209004906000526020600020908101906115ac91905b808211156118e25760008155600101611939565b60006020828403121561195f57600080fd5b81356001600160e01b0319811681146116cb57600080fd5b80356001600160a01b038116811461198e57600080fd5b919050565b6000602082840312156119a557600080fd5b6116cb82611977565b600081518084526020808501808196508360051b810191508286016000805b86811015611a25578385038a5282518051808752835b818110156119fe578281018901518882018a015288016119e3565b5086810188018490529a87019a601f01601f191690950186019450918501916001016119cd565b509298975050505050505050565b608081526000611a4660808301876119ae565b8281036020840152611a5881876119ae565b90508281036040840152611a6c81866119ae565b91505082606083015295945050505050565b600060208284031215611a9057600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611ad657611ad6611a97565b604052919050565b600082601f830112611aef57600080fd5b813567ffffffffffffffff811115611b0957611b09611a97565b611b1c601f8201601f1916602001611aad565b818152846020838601011115611b3157600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215611b6457600080fd5b611b6d85611977565b9350602085013567ffffffffffffffff80821115611b8a57600080fd5b611b9688838901611ade565b94506040870135915080821115611bac57600080fd5b611bb888838901611ade565b93506060870135915080821115611bce57600080fd5b50611bdb87828801611ade565b91505092959194509250565b60008060408385031215611bfa57600080fd5b82359150611c0a60208401611977565b90509250929050565b600082601f830112611c2457600080fd5b8135602067ffffffffffffffff80831115611c4157611c41611a97565b8260051b611c50838201611aad565b9384528581018301938381019088861115611c6a57600080fd5b84880192505b85831015611ca657823584811115611c885760008081fd5b611c968a87838c0101611ade565b8352509184019190840190611c70565b98975050505050505050565b60008060008060808587031215611cc857600080fd5b611cd185611977565b9350602085013567ffffffffffffffff80821115611cee57600080fd5b611cfa88838901611c13565b94506040870135915080821115611d1057600080fd5b611d1c88838901611c13565b93506060870135915080821115611d3257600080fd5b50611bdb87828801611c13565b600080600080600060a08688031215611d5757600080fd5b611d6086611977565b9450602086013567ffffffffffffffff80821115611d7d57600080fd5b611d8989838a01611c13565b95506040880135915080821115611d9f57600080fd5b611dab89838a01611c13565b94506060880135915080821115611dc157600080fd5b50611dce88828901611c13565b95989497509295608001359392505050565b600181811c90821680611df457607f821691505b602082108103611e1457634e487b7160e01b600052602260045260246000fd5b50919050565b606081526000611e2d60608301866119ae565b8281036020840152611e3f81866119ae565b90508281036040840152611e5381856119ae565b9695505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611e9b57611e9b611e73565b5060010190565b8181038181111561031d5761031d611e73565b634e487b7160e01b600052603160045260246000fd5b601f821115610b0257600081815260208120601f850160051c81016020861015611ef25750805b601f850160051c820191505b81811015611f1157828155600101611efe565b505050505050565b815167ffffffffffffffff811115611f3357611f33611a97565b611f4781611f418454611de0565b84611ecb565b602080601f831160018114611f7c5760008415611f645750858301515b600019600386901b1c1916600185901b178555611f11565b600085815260208120601f198616915b82811015611fab57888601518255948401946001909101908401611f8c565b5085821015611fc95787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fe50dd73a7b7a643432ced584fcc2cd0a1d5efa85899b62d28c5a52294d183d0e3a2646970667358221220b148152bc7ca0869bf7b758bc80746912692a196ba87c65fd05f463cf5dba47d64736f6c63430008140033",
}

// K8sAccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use K8sAccessControlMetaData.ABI instead.
var K8sAccessControlABI = K8sAccessControlMetaData.ABI

// K8sAccessControlBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use K8sAccessControlMetaData.Bin instead.
var K8sAccessControlBin = K8sAccessControlMetaData.Bin

// DeployK8sAccessControl deploys a new Ethereum contract, binding an instance of K8sAccessControl to it.
func DeployK8sAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *K8sAccessControl, error) {
	parsed, err := K8sAccessControlMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(K8sAccessControlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &K8sAccessControl{K8sAccessControlCaller: K8sAccessControlCaller{contract: contract}, K8sAccessControlTransactor: K8sAccessControlTransactor{contract: contract}, K8sAccessControlFilterer: K8sAccessControlFilterer{contract: contract}}, nil
}

// K8sAccessControl is an auto generated Go binding around an Ethereum contract.
type K8sAccessControl struct {
	K8sAccessControlCaller     // Read-only binding to the contract
	K8sAccessControlTransactor // Write-only binding to the contract
	K8sAccessControlFilterer   // Log filterer for contract events
}

// K8sAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type K8sAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// K8sAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type K8sAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// K8sAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type K8sAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// K8sAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type K8sAccessControlSession struct {
	Contract     *K8sAccessControl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// K8sAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type K8sAccessControlCallerSession struct {
	Contract *K8sAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// K8sAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type K8sAccessControlTransactorSession struct {
	Contract     *K8sAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// K8sAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type K8sAccessControlRaw struct {
	Contract *K8sAccessControl // Generic contract binding to access the raw methods on
}

// K8sAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type K8sAccessControlCallerRaw struct {
	Contract *K8sAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// K8sAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type K8sAccessControlTransactorRaw struct {
	Contract *K8sAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewK8sAccessControl creates a new instance of K8sAccessControl, bound to a specific deployed contract.
func NewK8sAccessControl(address common.Address, backend bind.ContractBackend) (*K8sAccessControl, error) {
	contract, err := bindK8sAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &K8sAccessControl{K8sAccessControlCaller: K8sAccessControlCaller{contract: contract}, K8sAccessControlTransactor: K8sAccessControlTransactor{contract: contract}, K8sAccessControlFilterer: K8sAccessControlFilterer{contract: contract}}, nil
}

// NewK8sAccessControlCaller creates a new read-only instance of K8sAccessControl, bound to a specific deployed contract.
func NewK8sAccessControlCaller(address common.Address, caller bind.ContractCaller) (*K8sAccessControlCaller, error) {
	contract, err := bindK8sAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &K8sAccessControlCaller{contract: contract}, nil
}

// NewK8sAccessControlTransactor creates a new write-only instance of K8sAccessControl, bound to a specific deployed contract.
func NewK8sAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*K8sAccessControlTransactor, error) {
	contract, err := bindK8sAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &K8sAccessControlTransactor{contract: contract}, nil
}

// NewK8sAccessControlFilterer creates a new log filterer instance of K8sAccessControl, bound to a specific deployed contract.
func NewK8sAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*K8sAccessControlFilterer, error) {
	contract, err := bindK8sAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &K8sAccessControlFilterer{contract: contract}, nil
}

// bindK8sAccessControl binds a generic wrapper to an already deployed contract.
func bindK8sAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := K8sAccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_K8sAccessControl *K8sAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _K8sAccessControl.Contract.K8sAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_K8sAccessControl *K8sAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.K8sAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_K8sAccessControl *K8sAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.K8sAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_K8sAccessControl *K8sAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _K8sAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_K8sAccessControl *K8sAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_K8sAccessControl *K8sAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _K8sAccessControl.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlSession) ADMINROLE() ([32]byte, error) {
	return _K8sAccessControl.Contract.ADMINROLE(&_K8sAccessControl.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlCallerSession) ADMINROLE() ([32]byte, error) {
	return _K8sAccessControl.Contract.ADMINROLE(&_K8sAccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _K8sAccessControl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _K8sAccessControl.Contract.DEFAULTADMINROLE(&_K8sAccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _K8sAccessControl.Contract.DEFAULTADMINROLE(&_K8sAccessControl.CallOpts)
}

// PERMISSIONMANAGERROLE is a free data retrieval call binding the contract method 0x0125a425.
//
// Solidity: function PERMISSION_MANAGER_ROLE() view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlCaller) PERMISSIONMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _K8sAccessControl.contract.Call(opts, &out, "PERMISSION_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMISSIONMANAGERROLE is a free data retrieval call binding the contract method 0x0125a425.
//
// Solidity: function PERMISSION_MANAGER_ROLE() view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlSession) PERMISSIONMANAGERROLE() ([32]byte, error) {
	return _K8sAccessControl.Contract.PERMISSIONMANAGERROLE(&_K8sAccessControl.CallOpts)
}

// PERMISSIONMANAGERROLE is a free data retrieval call binding the contract method 0x0125a425.
//
// Solidity: function PERMISSION_MANAGER_ROLE() view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlCallerSession) PERMISSIONMANAGERROLE() ([32]byte, error) {
	return _K8sAccessControl.Contract.PERMISSIONMANAGERROLE(&_K8sAccessControl.CallOpts)
}

// GetAuthorizedWalletAt is a free data retrieval call binding the contract method 0xe08951ea.
//
// Solidity: function getAuthorizedWalletAt(uint256 index) view returns(address)
func (_K8sAccessControl *K8sAccessControlCaller) GetAuthorizedWalletAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _K8sAccessControl.contract.Call(opts, &out, "getAuthorizedWalletAt", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizedWalletAt is a free data retrieval call binding the contract method 0xe08951ea.
//
// Solidity: function getAuthorizedWalletAt(uint256 index) view returns(address)
func (_K8sAccessControl *K8sAccessControlSession) GetAuthorizedWalletAt(index *big.Int) (common.Address, error) {
	return _K8sAccessControl.Contract.GetAuthorizedWalletAt(&_K8sAccessControl.CallOpts, index)
}

// GetAuthorizedWalletAt is a free data retrieval call binding the contract method 0xe08951ea.
//
// Solidity: function getAuthorizedWalletAt(uint256 index) view returns(address)
func (_K8sAccessControl *K8sAccessControlCallerSession) GetAuthorizedWalletAt(index *big.Int) (common.Address, error) {
	return _K8sAccessControl.Contract.GetAuthorizedWalletAt(&_K8sAccessControl.CallOpts, index)
}

// GetAuthorizedWalletsCount is a free data retrieval call binding the contract method 0x75f6db4e.
//
// Solidity: function getAuthorizedWalletsCount() view returns(uint256)
func (_K8sAccessControl *K8sAccessControlCaller) GetAuthorizedWalletsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _K8sAccessControl.contract.Call(opts, &out, "getAuthorizedWalletsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAuthorizedWalletsCount is a free data retrieval call binding the contract method 0x75f6db4e.
//
// Solidity: function getAuthorizedWalletsCount() view returns(uint256)
func (_K8sAccessControl *K8sAccessControlSession) GetAuthorizedWalletsCount() (*big.Int, error) {
	return _K8sAccessControl.Contract.GetAuthorizedWalletsCount(&_K8sAccessControl.CallOpts)
}

// GetAuthorizedWalletsCount is a free data retrieval call binding the contract method 0x75f6db4e.
//
// Solidity: function getAuthorizedWalletsCount() view returns(uint256)
func (_K8sAccessControl *K8sAccessControlCallerSession) GetAuthorizedWalletsCount() (*big.Int, error) {
	return _K8sAccessControl.Contract.GetAuthorizedWalletsCount(&_K8sAccessControl.CallOpts)
}

// GetPermissions is a free data retrieval call binding the contract method 0x160a7925.
//
// Solidity: function getPermissions(address wallet) view returns(string[] namespaces, string[] verbs, string[] resources, uint256 expiresAt)
func (_K8sAccessControl *K8sAccessControlCaller) GetPermissions(opts *bind.CallOpts, wallet common.Address) (struct {
	Namespaces []string
	Verbs      []string
	Resources  []string
	ExpiresAt  *big.Int
}, error) {
	var out []interface{}
	err := _K8sAccessControl.contract.Call(opts, &out, "getPermissions", wallet)

	outstruct := new(struct {
		Namespaces []string
		Verbs      []string
		Resources  []string
		ExpiresAt  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Namespaces = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Verbs = *abi.ConvertType(out[1], new([]string)).(*[]string)
	outstruct.Resources = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.ExpiresAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPermissions is a free data retrieval call binding the contract method 0x160a7925.
//
// Solidity: function getPermissions(address wallet) view returns(string[] namespaces, string[] verbs, string[] resources, uint256 expiresAt)
func (_K8sAccessControl *K8sAccessControlSession) GetPermissions(wallet common.Address) (struct {
	Namespaces []string
	Verbs      []string
	Resources  []string
	ExpiresAt  *big.Int
}, error) {
	return _K8sAccessControl.Contract.GetPermissions(&_K8sAccessControl.CallOpts, wallet)
}

// GetPermissions is a free data retrieval call binding the contract method 0x160a7925.
//
// Solidity: function getPermissions(address wallet) view returns(string[] namespaces, string[] verbs, string[] resources, uint256 expiresAt)
func (_K8sAccessControl *K8sAccessControlCallerSession) GetPermissions(wallet common.Address) (struct {
	Namespaces []string
	Verbs      []string
	Resources  []string
	ExpiresAt  *big.Int
}, error) {
	return _K8sAccessControl.Contract.GetPermissions(&_K8sAccessControl.CallOpts, wallet)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _K8sAccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _K8sAccessControl.Contract.GetRoleAdmin(&_K8sAccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_K8sAccessControl *K8sAccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _K8sAccessControl.Contract.GetRoleAdmin(&_K8sAccessControl.CallOpts, role)
}

// HasPermission is a free data retrieval call binding the contract method 0x27f1c1a7.
//
// Solidity: function hasPermission(address wallet, string namespace, string verb, string resource) view returns(bool)
func (_K8sAccessControl *K8sAccessControlCaller) HasPermission(opts *bind.CallOpts, wallet common.Address, namespace string, verb string, resource string) (bool, error) {
	var out []interface{}
	err := _K8sAccessControl.contract.Call(opts, &out, "hasPermission", wallet, namespace, verb, resource)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x27f1c1a7.
//
// Solidity: function hasPermission(address wallet, string namespace, string verb, string resource) view returns(bool)
func (_K8sAccessControl *K8sAccessControlSession) HasPermission(wallet common.Address, namespace string, verb string, resource string) (bool, error) {
	return _K8sAccessControl.Contract.HasPermission(&_K8sAccessControl.CallOpts, wallet, namespace, verb, resource)
}

// HasPermission is a free data retrieval call binding the contract method 0x27f1c1a7.
//
// Solidity: function hasPermission(address wallet, string namespace, string verb, string resource) view returns(bool)
func (_K8sAccessControl *K8sAccessControlCallerSession) HasPermission(wallet common.Address, namespace string, verb string, resource string) (bool, error) {
	return _K8sAccessControl.Contract.HasPermission(&_K8sAccessControl.CallOpts, wallet, namespace, verb, resource)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_K8sAccessControl *K8sAccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _K8sAccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_K8sAccessControl *K8sAccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _K8sAccessControl.Contract.HasRole(&_K8sAccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_K8sAccessControl *K8sAccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _K8sAccessControl.Contract.HasRole(&_K8sAccessControl.CallOpts, role, account)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address wallet) view returns(bool)
func (_K8sAccessControl *K8sAccessControlCaller) IsAuthorized(opts *bind.CallOpts, wallet common.Address) (bool, error) {
	var out []interface{}
	err := _K8sAccessControl.contract.Call(opts, &out, "isAuthorized", wallet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address wallet) view returns(bool)
func (_K8sAccessControl *K8sAccessControlSession) IsAuthorized(wallet common.Address) (bool, error) {
	return _K8sAccessControl.Contract.IsAuthorized(&_K8sAccessControl.CallOpts, wallet)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address wallet) view returns(bool)
func (_K8sAccessControl *K8sAccessControlCallerSession) IsAuthorized(wallet common.Address) (bool, error) {
	return _K8sAccessControl.Contract.IsAuthorized(&_K8sAccessControl.CallOpts, wallet)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_K8sAccessControl *K8sAccessControlCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _K8sAccessControl.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_K8sAccessControl *K8sAccessControlSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _K8sAccessControl.Contract.SupportsInterface(&_K8sAccessControl.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_K8sAccessControl *K8sAccessControlCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _K8sAccessControl.Contract.SupportsInterface(&_K8sAccessControl.CallOpts, interfaceId)
}

// AddPermissionManager is a paid mutator transaction binding the contract method 0x5958abb0.
//
// Solidity: function addPermissionManager(address account) returns()
func (_K8sAccessControl *K8sAccessControlTransactor) AddPermissionManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.contract.Transact(opts, "addPermissionManager", account)
}

// AddPermissionManager is a paid mutator transaction binding the contract method 0x5958abb0.
//
// Solidity: function addPermissionManager(address account) returns()
func (_K8sAccessControl *K8sAccessControlSession) AddPermissionManager(account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.AddPermissionManager(&_K8sAccessControl.TransactOpts, account)
}

// AddPermissionManager is a paid mutator transaction binding the contract method 0x5958abb0.
//
// Solidity: function addPermissionManager(address account) returns()
func (_K8sAccessControl *K8sAccessControlTransactorSession) AddPermissionManager(account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.AddPermissionManager(&_K8sAccessControl.TransactOpts, account)
}

// GrantPermission is a paid mutator transaction binding the contract method 0xdd4c93b4.
//
// Solidity: function grantPermission(address wallet, string[] namespaces, string[] verbs, string[] resources, uint256 expiresAt) returns()
func (_K8sAccessControl *K8sAccessControlTransactor) GrantPermission(opts *bind.TransactOpts, wallet common.Address, namespaces []string, verbs []string, resources []string, expiresAt *big.Int) (*types.Transaction, error) {
	return _K8sAccessControl.contract.Transact(opts, "grantPermission", wallet, namespaces, verbs, resources, expiresAt)
}

// GrantPermission is a paid mutator transaction binding the contract method 0xdd4c93b4.
//
// Solidity: function grantPermission(address wallet, string[] namespaces, string[] verbs, string[] resources, uint256 expiresAt) returns()
func (_K8sAccessControl *K8sAccessControlSession) GrantPermission(wallet common.Address, namespaces []string, verbs []string, resources []string, expiresAt *big.Int) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.GrantPermission(&_K8sAccessControl.TransactOpts, wallet, namespaces, verbs, resources, expiresAt)
}

// GrantPermission is a paid mutator transaction binding the contract method 0xdd4c93b4.
//
// Solidity: function grantPermission(address wallet, string[] namespaces, string[] verbs, string[] resources, uint256 expiresAt) returns()
func (_K8sAccessControl *K8sAccessControlTransactorSession) GrantPermission(wallet common.Address, namespaces []string, verbs []string, resources []string, expiresAt *big.Int) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.GrantPermission(&_K8sAccessControl.TransactOpts, wallet, namespaces, verbs, resources, expiresAt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_K8sAccessControl *K8sAccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_K8sAccessControl *K8sAccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.GrantRole(&_K8sAccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_K8sAccessControl *K8sAccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.GrantRole(&_K8sAccessControl.TransactOpts, role, account)
}

// RemovePermissionManager is a paid mutator transaction binding the contract method 0xecb80b6f.
//
// Solidity: function removePermissionManager(address account) returns()
func (_K8sAccessControl *K8sAccessControlTransactor) RemovePermissionManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.contract.Transact(opts, "removePermissionManager", account)
}

// RemovePermissionManager is a paid mutator transaction binding the contract method 0xecb80b6f.
//
// Solidity: function removePermissionManager(address account) returns()
func (_K8sAccessControl *K8sAccessControlSession) RemovePermissionManager(account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.RemovePermissionManager(&_K8sAccessControl.TransactOpts, account)
}

// RemovePermissionManager is a paid mutator transaction binding the contract method 0xecb80b6f.
//
// Solidity: function removePermissionManager(address account) returns()
func (_K8sAccessControl *K8sAccessControlTransactorSession) RemovePermissionManager(account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.RemovePermissionManager(&_K8sAccessControl.TransactOpts, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_K8sAccessControl *K8sAccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_K8sAccessControl *K8sAccessControlSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.RenounceRole(&_K8sAccessControl.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_K8sAccessControl *K8sAccessControlTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.RenounceRole(&_K8sAccessControl.TransactOpts, role, callerConfirmation)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa3bf5b9a.
//
// Solidity: function revokePermission(address wallet) returns()
func (_K8sAccessControl *K8sAccessControlTransactor) RevokePermission(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.contract.Transact(opts, "revokePermission", wallet)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa3bf5b9a.
//
// Solidity: function revokePermission(address wallet) returns()
func (_K8sAccessControl *K8sAccessControlSession) RevokePermission(wallet common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.RevokePermission(&_K8sAccessControl.TransactOpts, wallet)
}

// RevokePermission is a paid mutator transaction binding the contract method 0xa3bf5b9a.
//
// Solidity: function revokePermission(address wallet) returns()
func (_K8sAccessControl *K8sAccessControlTransactorSession) RevokePermission(wallet common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.RevokePermission(&_K8sAccessControl.TransactOpts, wallet)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_K8sAccessControl *K8sAccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_K8sAccessControl *K8sAccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.RevokeRole(&_K8sAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_K8sAccessControl *K8sAccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.RevokeRole(&_K8sAccessControl.TransactOpts, role, account)
}

// UpdatePermission is a paid mutator transaction binding the contract method 0x6811ba2f.
//
// Solidity: function updatePermission(address wallet, string[] namespaces, string[] verbs, string[] resources) returns()
func (_K8sAccessControl *K8sAccessControlTransactor) UpdatePermission(opts *bind.TransactOpts, wallet common.Address, namespaces []string, verbs []string, resources []string) (*types.Transaction, error) {
	return _K8sAccessControl.contract.Transact(opts, "updatePermission", wallet, namespaces, verbs, resources)
}

// UpdatePermission is a paid mutator transaction binding the contract method 0x6811ba2f.
//
// Solidity: function updatePermission(address wallet, string[] namespaces, string[] verbs, string[] resources) returns()
func (_K8sAccessControl *K8sAccessControlSession) UpdatePermission(wallet common.Address, namespaces []string, verbs []string, resources []string) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.UpdatePermission(&_K8sAccessControl.TransactOpts, wallet, namespaces, verbs, resources)
}

// UpdatePermission is a paid mutator transaction binding the contract method 0x6811ba2f.
//
// Solidity: function updatePermission(address wallet, string[] namespaces, string[] verbs, string[] resources) returns()
func (_K8sAccessControl *K8sAccessControlTransactorSession) UpdatePermission(wallet common.Address, namespaces []string, verbs []string, resources []string) (*types.Transaction, error) {
	return _K8sAccessControl.Contract.UpdatePermission(&_K8sAccessControl.TransactOpts, wallet, namespaces, verbs, resources)
}

// K8sAccessControlPermissionGrantedIterator is returned from FilterPermissionGranted and is used to iterate over the raw logs and unpacked data for PermissionGranted events raised by the K8sAccessControl contract.
type K8sAccessControlPermissionGrantedIterator struct {
	Event *K8sAccessControlPermissionGranted // Event containing the contract specifics and raw log

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
func (it *K8sAccessControlPermissionGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(K8sAccessControlPermissionGranted)
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
		it.Event = new(K8sAccessControlPermissionGranted)
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
func (it *K8sAccessControlPermissionGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *K8sAccessControlPermissionGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// K8sAccessControlPermissionGranted represents a PermissionGranted event raised by the K8sAccessControl contract.
type K8sAccessControlPermissionGranted struct {
	Wallet     common.Address
	Namespaces []string
	Verbs      []string
	Resources  []string
	ExpiresAt  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionGranted is a free log retrieval operation binding the contract event 0x4a58c34854ef823594749e5d55b169493cf1e79b11ba2567eb0053addadf5aa4.
//
// Solidity: event PermissionGranted(address indexed wallet, string[] namespaces, string[] verbs, string[] resources, uint256 expiresAt)
func (_K8sAccessControl *K8sAccessControlFilterer) FilterPermissionGranted(opts *bind.FilterOpts, wallet []common.Address) (*K8sAccessControlPermissionGrantedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _K8sAccessControl.contract.FilterLogs(opts, "PermissionGranted", walletRule)
	if err != nil {
		return nil, err
	}
	return &K8sAccessControlPermissionGrantedIterator{contract: _K8sAccessControl.contract, event: "PermissionGranted", logs: logs, sub: sub}, nil
}

// WatchPermissionGranted is a free log subscription operation binding the contract event 0x4a58c34854ef823594749e5d55b169493cf1e79b11ba2567eb0053addadf5aa4.
//
// Solidity: event PermissionGranted(address indexed wallet, string[] namespaces, string[] verbs, string[] resources, uint256 expiresAt)
func (_K8sAccessControl *K8sAccessControlFilterer) WatchPermissionGranted(opts *bind.WatchOpts, sink chan<- *K8sAccessControlPermissionGranted, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _K8sAccessControl.contract.WatchLogs(opts, "PermissionGranted", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(K8sAccessControlPermissionGranted)
				if err := _K8sAccessControl.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
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

// ParsePermissionGranted is a log parse operation binding the contract event 0x4a58c34854ef823594749e5d55b169493cf1e79b11ba2567eb0053addadf5aa4.
//
// Solidity: event PermissionGranted(address indexed wallet, string[] namespaces, string[] verbs, string[] resources, uint256 expiresAt)
func (_K8sAccessControl *K8sAccessControlFilterer) ParsePermissionGranted(log types.Log) (*K8sAccessControlPermissionGranted, error) {
	event := new(K8sAccessControlPermissionGranted)
	if err := _K8sAccessControl.contract.UnpackLog(event, "PermissionGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// K8sAccessControlPermissionRevokedIterator is returned from FilterPermissionRevoked and is used to iterate over the raw logs and unpacked data for PermissionRevoked events raised by the K8sAccessControl contract.
type K8sAccessControlPermissionRevokedIterator struct {
	Event *K8sAccessControlPermissionRevoked // Event containing the contract specifics and raw log

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
func (it *K8sAccessControlPermissionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(K8sAccessControlPermissionRevoked)
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
		it.Event = new(K8sAccessControlPermissionRevoked)
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
func (it *K8sAccessControlPermissionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *K8sAccessControlPermissionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// K8sAccessControlPermissionRevoked represents a PermissionRevoked event raised by the K8sAccessControl contract.
type K8sAccessControlPermissionRevoked struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPermissionRevoked is a free log retrieval operation binding the contract event 0x3541f93cbae8c4be65491b824efe1570976e740b18c6aa441db5291f4de4c921.
//
// Solidity: event PermissionRevoked(address indexed wallet)
func (_K8sAccessControl *K8sAccessControlFilterer) FilterPermissionRevoked(opts *bind.FilterOpts, wallet []common.Address) (*K8sAccessControlPermissionRevokedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _K8sAccessControl.contract.FilterLogs(opts, "PermissionRevoked", walletRule)
	if err != nil {
		return nil, err
	}
	return &K8sAccessControlPermissionRevokedIterator{contract: _K8sAccessControl.contract, event: "PermissionRevoked", logs: logs, sub: sub}, nil
}

// WatchPermissionRevoked is a free log subscription operation binding the contract event 0x3541f93cbae8c4be65491b824efe1570976e740b18c6aa441db5291f4de4c921.
//
// Solidity: event PermissionRevoked(address indexed wallet)
func (_K8sAccessControl *K8sAccessControlFilterer) WatchPermissionRevoked(opts *bind.WatchOpts, sink chan<- *K8sAccessControlPermissionRevoked, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _K8sAccessControl.contract.WatchLogs(opts, "PermissionRevoked", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(K8sAccessControlPermissionRevoked)
				if err := _K8sAccessControl.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
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

// ParsePermissionRevoked is a log parse operation binding the contract event 0x3541f93cbae8c4be65491b824efe1570976e740b18c6aa441db5291f4de4c921.
//
// Solidity: event PermissionRevoked(address indexed wallet)
func (_K8sAccessControl *K8sAccessControlFilterer) ParsePermissionRevoked(log types.Log) (*K8sAccessControlPermissionRevoked, error) {
	event := new(K8sAccessControlPermissionRevoked)
	if err := _K8sAccessControl.contract.UnpackLog(event, "PermissionRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// K8sAccessControlPermissionUpdatedIterator is returned from FilterPermissionUpdated and is used to iterate over the raw logs and unpacked data for PermissionUpdated events raised by the K8sAccessControl contract.
type K8sAccessControlPermissionUpdatedIterator struct {
	Event *K8sAccessControlPermissionUpdated // Event containing the contract specifics and raw log

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
func (it *K8sAccessControlPermissionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(K8sAccessControlPermissionUpdated)
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
		it.Event = new(K8sAccessControlPermissionUpdated)
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
func (it *K8sAccessControlPermissionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *K8sAccessControlPermissionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// K8sAccessControlPermissionUpdated represents a PermissionUpdated event raised by the K8sAccessControl contract.
type K8sAccessControlPermissionUpdated struct {
	Wallet     common.Address
	Namespaces []string
	Verbs      []string
	Resources  []string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPermissionUpdated is a free log retrieval operation binding the contract event 0x259165b83cff236056b54312412ad942d4af8eddf7d3e92032bff6718465f3ba.
//
// Solidity: event PermissionUpdated(address indexed wallet, string[] namespaces, string[] verbs, string[] resources)
func (_K8sAccessControl *K8sAccessControlFilterer) FilterPermissionUpdated(opts *bind.FilterOpts, wallet []common.Address) (*K8sAccessControlPermissionUpdatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _K8sAccessControl.contract.FilterLogs(opts, "PermissionUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return &K8sAccessControlPermissionUpdatedIterator{contract: _K8sAccessControl.contract, event: "PermissionUpdated", logs: logs, sub: sub}, nil
}

// WatchPermissionUpdated is a free log subscription operation binding the contract event 0x259165b83cff236056b54312412ad942d4af8eddf7d3e92032bff6718465f3ba.
//
// Solidity: event PermissionUpdated(address indexed wallet, string[] namespaces, string[] verbs, string[] resources)
func (_K8sAccessControl *K8sAccessControlFilterer) WatchPermissionUpdated(opts *bind.WatchOpts, sink chan<- *K8sAccessControlPermissionUpdated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _K8sAccessControl.contract.WatchLogs(opts, "PermissionUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(K8sAccessControlPermissionUpdated)
				if err := _K8sAccessControl.contract.UnpackLog(event, "PermissionUpdated", log); err != nil {
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

// ParsePermissionUpdated is a log parse operation binding the contract event 0x259165b83cff236056b54312412ad942d4af8eddf7d3e92032bff6718465f3ba.
//
// Solidity: event PermissionUpdated(address indexed wallet, string[] namespaces, string[] verbs, string[] resources)
func (_K8sAccessControl *K8sAccessControlFilterer) ParsePermissionUpdated(log types.Log) (*K8sAccessControlPermissionUpdated, error) {
	event := new(K8sAccessControlPermissionUpdated)
	if err := _K8sAccessControl.contract.UnpackLog(event, "PermissionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// K8sAccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the K8sAccessControl contract.
type K8sAccessControlRoleAdminChangedIterator struct {
	Event *K8sAccessControlRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *K8sAccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(K8sAccessControlRoleAdminChanged)
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
		it.Event = new(K8sAccessControlRoleAdminChanged)
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
func (it *K8sAccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *K8sAccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// K8sAccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the K8sAccessControl contract.
type K8sAccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_K8sAccessControl *K8sAccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*K8sAccessControlRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _K8sAccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &K8sAccessControlRoleAdminChangedIterator{contract: _K8sAccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_K8sAccessControl *K8sAccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *K8sAccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _K8sAccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(K8sAccessControlRoleAdminChanged)
				if err := _K8sAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_K8sAccessControl *K8sAccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*K8sAccessControlRoleAdminChanged, error) {
	event := new(K8sAccessControlRoleAdminChanged)
	if err := _K8sAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// K8sAccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the K8sAccessControl contract.
type K8sAccessControlRoleGrantedIterator struct {
	Event *K8sAccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *K8sAccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(K8sAccessControlRoleGranted)
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
		it.Event = new(K8sAccessControlRoleGranted)
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
func (it *K8sAccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *K8sAccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// K8sAccessControlRoleGranted represents a RoleGranted event raised by the K8sAccessControl contract.
type K8sAccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_K8sAccessControl *K8sAccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*K8sAccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _K8sAccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &K8sAccessControlRoleGrantedIterator{contract: _K8sAccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_K8sAccessControl *K8sAccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *K8sAccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _K8sAccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(K8sAccessControlRoleGranted)
				if err := _K8sAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_K8sAccessControl *K8sAccessControlFilterer) ParseRoleGranted(log types.Log) (*K8sAccessControlRoleGranted, error) {
	event := new(K8sAccessControlRoleGranted)
	if err := _K8sAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// K8sAccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the K8sAccessControl contract.
type K8sAccessControlRoleRevokedIterator struct {
	Event *K8sAccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *K8sAccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(K8sAccessControlRoleRevoked)
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
		it.Event = new(K8sAccessControlRoleRevoked)
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
func (it *K8sAccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *K8sAccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// K8sAccessControlRoleRevoked represents a RoleRevoked event raised by the K8sAccessControl contract.
type K8sAccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_K8sAccessControl *K8sAccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*K8sAccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _K8sAccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &K8sAccessControlRoleRevokedIterator{contract: _K8sAccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_K8sAccessControl *K8sAccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *K8sAccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _K8sAccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(K8sAccessControlRoleRevoked)
				if err := _K8sAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_K8sAccessControl *K8sAccessControlFilterer) ParseRoleRevoked(log types.Log) (*K8sAccessControlRoleRevoked, error) {
	event := new(K8sAccessControlRoleRevoked)
	if err := _K8sAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
