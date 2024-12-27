// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC721Handler

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

// ERC721HandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type ERC721HandlerDepositRecord struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}

// ERC721HandlerMetaData contains all meta data concerning the ERC721Handler contract.
var ERC721HandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"burnableContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"fundERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structERC721Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200291038038062002910833981810160405281019062000037919062000489565b81518351146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000632565b60405180910390fd5b836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8351811015620001165762000108848281518110620000df57fe5b6020026020010151848381518110620000f457fe5b60200260200101516200016560201b60201c565b8080600101915050620000c4565b5060008090505b81518110156200015a576200014c8282815181106200013857fe5b60200260200101516200025760201b60201c565b80806001019150506200011d565b505050505062000757565b806002600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16620002e6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002dd9062000610565b60405180910390fd5b6001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600081519050620003528162000723565b92915050565b600082601f8301126200036a57600080fd5b8151620003816200037b8262000682565b62000654565b91508181835260208401935060208101905083856020840282011115620003a757600080fd5b60005b83811015620003db5781620003c0888262000341565b845260208401935060208301925050600181019050620003aa565b5050505092915050565b600082601f830112620003f757600080fd5b81516200040e6200040882620006ab565b62000654565b915081818352602084019350602081019050838560208402820111156200043457600080fd5b60005b838110156200046857816200044d888262000472565b84526020840193506020830192505060018101905062000437565b5050505092915050565b60008151905062000483816200073d565b92915050565b60008060008060808587031215620004a057600080fd5b6000620004b08782880162000341565b945050602085015167ffffffffffffffff811115620004ce57600080fd5b620004dc87828801620003e5565b935050604085015167ffffffffffffffff811115620004fa57600080fd5b620005088782880162000358565b925050606085015167ffffffffffffffff8111156200052657600080fd5b620005348782880162000358565b91505092959194509250565b60006200054f602483620006d4565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000620005b7603c83620006d4565b91507f696e697469616c5265736f7572636549447320616e6420696e697469616c436f60008301527f6e7472616374416464726573736573206c656e206d69736d61746368000000006020830152604082019050919050565b600060208201905081810360008301526200062b8162000540565b9050919050565b600060208201905081810360008301526200064d81620005a8565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200067857600080fd5b8060405250919050565b600067ffffffffffffffff8211156200069a57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115620006c357600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620006f28262000703565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200072e81620006e5565b81146200073a57600080fd5b50565b6200074881620006f9565b81146200075457600080fd5b50565b6121a980620007676000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80637354298011610097578063c3b6581611610066578063c3b658161461029b578063c8ba6c87146102b9578063d9caed12146102e9578063e248cff214610305576100f5565b806373542980146102035780637f79bea81461021f578063b8fa37361461024f578063ba484c091461026b576100f5565b806338995da9116100d357806338995da9146101645780634402027f146101805780636817031b146101b75780636a70d081146101d3576100f5565b806307b7ed99146100fa5780630a6d55d814610116578063318c136e14610146575b600080fd5b610114600480360381019061010f919061174f565b610321565b005b610130600480360381019061012b91906117f0565b610335565b60405161013d9190611d44565b60405180910390f35b61014e610368565b60405161015b9190611d44565b60405180910390f35b61017e600480360381019061017991906118ad565b61038d565b005b61019a600480360381019061019591906119bc565b610791565b6040516101ae989796959493929190611dd4565b60405180910390f35b6101d160048036038101906101cc919061174f565b610970565b005b6101ed60048036038101906101e8919061174f565b6109bc565b6040516101fa9190611e60565b60405180910390f35b61021d60048036038101906102189190611778565b6109dc565b005b6102396004803603810190610234919061174f565b610a56565b6040516102469190611e60565b60405180910390f35b61026960048036038101906102649190611819565b610a76565b005b61028560048036038101906102809190611980565b610a8c565b6040516102929190611f11565b60405180910390f35b6102a3610d24565b6040516102b09190611d44565b60405180910390f35b6102d360048036038101906102ce919061174f565b610d4a565b6040516102e09190611e7b565b60405180910390f35b61030360048036038101906102fe9190611778565b610d62565b005b61031f600480360381019061031a9190611855565b610d7b565b005b610329610f18565b61033281610fa9565b50565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610395610f18565b60008060608060c435925060e4359350604051915083820160200160405260e4360360e483376000600260008c815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661047f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047690611ef1565b60405180910390fd5b6104af635b5e139f60e01b8273ffffffffffffffffffffffffffffffffffffffff1661109090919063ffffffff16565b1561054b5760008190508073ffffffffffffffffffffffffffffffffffffffff1663c87b56dd866040518263ffffffff1660e01b81526004016104f29190611f33565b60006040518083038186803b15801561050a57600080fd5b505afa15801561051e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610547919061193f565b9250505b600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156105ac576105a781856110b5565b6105b9565b6105b88189308761112a565b5b6040518061010001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018660ff1681526020018b60ff1681526020018c81526020018481526020018973ffffffffffffffffffffffffffffffffffffffff16815260200185815260200183815250600660008c60ff1660ff16815260200190815260200160002060008b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff16021790555060408201518160000160156101000a81548160ff021916908360ff160217905550606082015181600101556080820151816002019080519060200190610712929190611514565b5060a08201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816004015560e0820151816005019080519060200190610780929190611514565b509050505050505050505050505050565b6006602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16908060000160159054906101000a900460ff1690806001015490806002018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561089c5780601f106108715761010080835404028352916020019161089c565b820191906000526020600020905b81548152906001019060200180831161087f57829003601f168201915b5050505050908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806004015490806005018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109665780601f1061093b57610100808354040283529160200191610966565b820191906000526020600020905b81548152906001019060200180831161094957829003601f168201915b5050505050905088565b610978610f18565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60056020528060005260406000206000915054906101000a900460ff1681565b60008390508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b8152600401610a1e93929190611d5f565b600060405180830381600087803b158015610a3857600080fd5b505af1158015610a4c573d6000803e3d6000fd5b5050505050505050565b60046020528060005260406000206000915054906101000a900460ff1681565b610a7e610f18565b610a8882826111a5565b5050565b610a94611594565b600660008360ff1660ff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020604051806101000160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016000820160159054906101000a900460ff1660ff1660ff16815260200160018201548152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c115780601f10610be657610100808354040283529160200191610c11565b820191906000526020600020905b815481529060010190602001808311610bf457829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d135780601f10610ce857610100808354040283529160200191610d13565b820191906000526020600020905b815481529060010190602001808311610cf657829003601f168201915b505050505081525050905092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915090505481565b610d6a610f18565b610d7683308484611297565b505050565b610d83610f18565b6000606080606435925060405191506084358060a401358184016040016040528060200160840136036084853760405192508083016040016040528160a40180360381853750505060006020830151905060006002600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610e98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8f90611ef1565b60405180910390fd5b600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610efe57610ef9818360601c8786611312565b610f0e565b610f0d81308460601c88611297565b5b5050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fa7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9e90611eb1565b60405180910390fd5b565b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611035576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102c90611ed1565b60405180910390fd5b6001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600061109b8361138d565b80156110ad57506110ac83836113c1565b5b905092915050565b60008290508073ffffffffffffffffffffffffffffffffffffffff166342966c68836040518263ffffffff1660e01b81526004016110f39190611f33565b600060405180830381600087803b15801561110d57600080fd5b505af1158015611121573d6000803e3d6000fd5b50505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b815260040161116c93929190611d5f565b600060405180830381600087803b15801561118657600080fd5b505af115801561119a573d6000803e3d6000fd5b505050505050505050565b806002600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8585856040518463ffffffff1660e01b81526004016112d993929190611d5f565b600060405180830381600087803b1580156112f357600080fd5b505af1158015611307573d6000803e3d6000fd5b505050505050505050565b60008490508073ffffffffffffffffffffffffffffffffffffffff1663d3fc98648585856040518463ffffffff1660e01b815260040161135493929190611d96565b600060405180830381600087803b15801561136e57600080fd5b505af1158015611382573d6000803e3d6000fd5b505050505050505050565b60006113a0826301ffc9a760e01b6113c1565b80156113ba57506113b88263ffffffff60e01b6113c1565b155b9050919050565b60008060006113d085856113e8565b915091508180156113de5750805b9250505092915050565b60008060606301ffc9a760e01b846040516024016114069190611e96565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050600060608673ffffffffffffffffffffffffffffffffffffffff16617530846040516114919190611d2d565b6000604051808303818686fa925050503d80600081146114cd576040519150601f19603f3d011682016040523d82523d6000602084013e6114d2565b606091505b50915091506020815110156114f0576000809450945050505061150d565b818180602001905181019061150591906117c7565b945094505050505b9250929050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061155557805160ff1916838001178555611583565b82800160010185558215611583579182015b82811115611582578251825591602001919060010190611567565b5b509050611590919061160e565b5090565b604051806101000160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600060ff1681526020016000801916815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001606081525090565b61163091905b8082111561162c576000816000905550600101611614565b5090565b90565b600081359050611642816120e9565b92915050565b60008151905061165781612100565b92915050565b60008135905061166c81612117565b92915050565b60008083601f84011261168457600080fd5b8235905067ffffffffffffffff81111561169d57600080fd5b6020830191508360018202830111156116b557600080fd5b9250929050565b600082601f8301126116cd57600080fd5b81516116e06116db82611f7b565b611f4e565b915080825260208301602083018583830111156116fc57600080fd5b6117078382846120a5565b50505092915050565b60008135905061171f8161212e565b92915050565b60008135905061173481612145565b92915050565b6000813590506117498161215c565b92915050565b60006020828403121561176157600080fd5b600061176f84828501611633565b91505092915050565b60008060006060848603121561178d57600080fd5b600061179b86828701611633565b93505060206117ac86828701611633565b92505060406117bd86828701611710565b9150509250925092565b6000602082840312156117d957600080fd5b60006117e784828501611648565b91505092915050565b60006020828403121561180257600080fd5b60006118108482850161165d565b91505092915050565b6000806040838503121561182c57600080fd5b600061183a8582860161165d565b925050602061184b85828601611633565b9150509250929050565b60008060006040848603121561186a57600080fd5b60006118788682870161165d565b935050602084013567ffffffffffffffff81111561189557600080fd5b6118a186828701611672565b92509250509250925092565b60008060008060008060a087890312156118c657600080fd5b60006118d489828a0161165d565b96505060206118e589828a0161173a565b95505060406118f689828a01611725565b945050606061190789828a01611633565b935050608087013567ffffffffffffffff81111561192457600080fd5b61193089828a01611672565b92509250509295509295509295565b60006020828403121561195157600080fd5b600082015167ffffffffffffffff81111561196b57600080fd5b611977848285016116bc565b91505092915050565b6000806040838503121561199357600080fd5b60006119a185828601611725565b92505060206119b28582860161173a565b9150509250929050565b600080604083850312156119cf57600080fd5b60006119dd8582860161173a565b92505060206119ee85828601611725565b9150509250929050565b611a0181612006565b82525050565b611a1081612006565b82525050565b611a1f81612018565b82525050565b611a2e81612024565b82525050565b611a3d81612024565b82525050565b611a4c8161202e565b82525050565b6000611a5d82611fb2565b611a678185611fea565b9350611a778185602086016120a5565b80840191505092915050565b6000611a8e82611fa7565b611a988185611fc8565b9350611aa88185602086016120a5565b611ab1816120d8565b840191505092915050565b6000611ac782611fa7565b611ad18185611fd9565b9350611ae18185602086016120a5565b611aea816120d8565b840191505092915050565b6000611b0082611fbd565b611b0a8185611ff5565b9350611b1a8185602086016120a5565b611b23816120d8565b840191505092915050565b6000611b3b601e83611ff5565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b6000611b7b602483611ff5565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611be1602883611ff5565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b600061010083016000830151611c5360008601826119f8565b506020830151611c666020860182611d0f565b506040830151611c796040860182611d0f565b506060830151611c8c6060860182611a25565b5060808301518482036080860152611ca48282611a83565b91505060a0830151611cb960a08601826119f8565b5060c0830151611ccc60c0860182611cf1565b5060e083015184820360e0860152611ce48282611a83565b9150508091505092915050565b611cfa8161207a565b82525050565b611d098161207a565b82525050565b611d1881612098565b82525050565b611d2781612098565b82525050565b6000611d398284611a52565b915081905092915050565b6000602082019050611d596000830184611a07565b92915050565b6000606082019050611d746000830186611a07565b611d816020830185611a07565b611d8e6040830184611d00565b949350505050565b6000606082019050611dab6000830186611a07565b611db86020830185611d00565b8181036040830152611dca8184611af5565b9050949350505050565b600061010082019050611dea600083018b611a07565b611df7602083018a611d1e565b611e046040830189611d1e565b611e116060830188611a34565b8181036080830152611e238187611abc565b9050611e3260a0830186611a07565b611e3f60c0830185611d00565b81810360e0830152611e518184611abc565b90509998505050505050505050565b6000602082019050611e756000830184611a16565b92915050565b6000602082019050611e906000830184611a34565b92915050565b6000602082019050611eab6000830184611a43565b92915050565b60006020820190508181036000830152611eca81611b2e565b9050919050565b60006020820190508181036000830152611eea81611b6e565b9050919050565b60006020820190508181036000830152611f0a81611bd4565b9050919050565b60006020820190508181036000830152611f2b8184611c3a565b905092915050565b6000602082019050611f486000830184611d00565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611f7157600080fd5b8060405250919050565b600067ffffffffffffffff821115611f9257600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006120118261205a565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b60005b838110156120c35780820151818401526020810190506120a8565b838111156120d2576000848401525b50505050565b6000601f19601f8301169050919050565b6120f281612006565b81146120fd57600080fd5b50565b61210981612018565b811461211457600080fd5b50565b61212081612024565b811461212b57600080fd5b50565b6121378161207a565b811461214257600080fd5b50565b61214e81612084565b811461215957600080fd5b50565b61216581612098565b811461217057600080fd5b5056fea2646970667358221220723e97b3950fb46151dc77f0b4bf867c50d26f09ff0a67945609bbc04e15dd8364736f6c63430006040033",
}

// ERC721HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721HandlerMetaData.ABI instead.
var ERC721HandlerABI = ERC721HandlerMetaData.ABI

// ERC721HandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC721HandlerMetaData.Bin instead.
var ERC721HandlerBin = ERC721HandlerMetaData.Bin

// DeployERC721Handler deploys a new Ethereum contract, binding an instance of ERC721Handler to it.
func DeployERC721Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableContractAddresses []common.Address) (common.Address, *types.Transaction, *ERC721Handler, error) {
	parsed, err := ERC721HandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC721HandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses, burnableContractAddresses)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Handler{ERC721HandlerCaller: ERC721HandlerCaller{contract: contract}, ERC721HandlerTransactor: ERC721HandlerTransactor{contract: contract}, ERC721HandlerFilterer: ERC721HandlerFilterer{contract: contract}}, nil
}

// ERC721Handler is an auto generated Go binding around an Ethereum contract.
type ERC721Handler struct {
	ERC721HandlerCaller     // Read-only binding to the contract
	ERC721HandlerTransactor // Write-only binding to the contract
	ERC721HandlerFilterer   // Log filterer for contract events
}

// ERC721HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721HandlerSession struct {
	Contract     *ERC721Handler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721HandlerCallerSession struct {
	Contract *ERC721HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC721HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721HandlerTransactorSession struct {
	Contract     *ERC721HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC721HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721HandlerRaw struct {
	Contract *ERC721Handler // Generic contract binding to access the raw methods on
}

// ERC721HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721HandlerCallerRaw struct {
	Contract *ERC721HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721HandlerTransactorRaw struct {
	Contract *ERC721HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Handler creates a new instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721Handler(address common.Address, backend bind.ContractBackend) (*ERC721Handler, error) {
	contract, err := bindERC721Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Handler{ERC721HandlerCaller: ERC721HandlerCaller{contract: contract}, ERC721HandlerTransactor: ERC721HandlerTransactor{contract: contract}, ERC721HandlerFilterer: ERC721HandlerFilterer{contract: contract}}, nil
}

// NewERC721HandlerCaller creates a new read-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC721HandlerCaller, error) {
	contract, err := bindERC721Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerCaller{contract: contract}, nil
}

// NewERC721HandlerTransactor creates a new write-only instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721HandlerTransactor, error) {
	contract, err := bindERC721Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerTransactor{contract: contract}, nil
}

// NewERC721HandlerFilterer creates a new log filterer instance of ERC721Handler, bound to a specific deployed contract.
func NewERC721HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721HandlerFilterer, error) {
	contract, err := bindERC721Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721HandlerFilterer{contract: contract}, nil
}

// bindERC721Handler binds a generic wrapper to an already deployed contract.
func bindERC721Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC721HandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.ERC721HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ERC721HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Handler *ERC721HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Handler *ERC721HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Handler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC721Handler.Contract.BridgeAddress(&_ERC721Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.BurnList(&_ERC721Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.BurnList(&_ERC721Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.ContractWhitelist(&_ERC721Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC721Handler *ERC721HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC721Handler.Contract.ContractWhitelist(&_ERC721Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	outstruct := new(struct {
		TokenAddress                   common.Address
		LenDestinationRecipientAddress uint8
		DestinationChainID             uint8
		ResourceID                     [32]byte
		DestinationRecipientAddress    []byte
		Depositer                      common.Address
		TokenID                        *big.Int
		MetaData                       []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LenDestinationRecipientAddress = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.DestinationChainID = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ResourceID = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.DestinationRecipientAddress = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.Depositer = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.TokenID = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.MetaData = *abi.ConvertType(out[7], new([]byte)).(*[]byte)

	return *outstruct, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}, error) {
	return _ERC721Handler.Contract.DepositRecords(&_ERC721Handler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _tokenID, bytes _metaData)
func (_ERC721Handler *ERC721HandlerCallerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	TokenID                        *big.Int
	MetaData                       []byte
}, error) {
	return _ERC721Handler.Contract.DepositRecords(&_ERC721Handler.CallOpts, arg0, arg1)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToTokenContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC721Handler.Contract.ResourceIDToTokenContractAddress(&_ERC721Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.TokenContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC721Handler *ERC721HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC721Handler.Contract.TokenContractAddressToResourceID(&_ERC721Handler.CallOpts, arg0)
}

// VaultAddress is a free data retrieval call binding the contract method 0xc3b65816.
//
// Solidity: function _vaultAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCaller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "_vaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAddress is a free data retrieval call binding the contract method 0xc3b65816.
//
// Solidity: function _vaultAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerSession) VaultAddress() (common.Address, error) {
	return _ERC721Handler.Contract.VaultAddress(&_ERC721Handler.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0xc3b65816.
//
// Solidity: function _vaultAddress() view returns(address)
func (_ERC721Handler *ERC721HandlerCallerSession) VaultAddress() (common.Address, error) {
	return _ERC721Handler.Contract.VaultAddress(&_ERC721Handler.CallOpts)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256,bytes))
func (_ERC721Handler *ERC721HandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce uint64, destId uint8) (ERC721HandlerDepositRecord, error) {
	var out []interface{}
	err := _ERC721Handler.contract.Call(opts, &out, "getDepositRecord", depositNonce, destId)

	if err != nil {
		return *new(ERC721HandlerDepositRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC721HandlerDepositRecord)).(*ERC721HandlerDepositRecord)

	return out0, err

}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256,bytes))
func (_ERC721Handler *ERC721HandlerSession) GetDepositRecord(depositNonce uint64, destId uint8) (ERC721HandlerDepositRecord, error) {
	return _ERC721Handler.Contract.GetDepositRecord(&_ERC721Handler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256,bytes))
func (_ERC721Handler *ERC721HandlerCallerSession) GetDepositRecord(depositNonce uint64, destId uint8) (ERC721HandlerDepositRecord, error) {
	return _ERC721Handler.Contract.GetDepositRecord(&_ERC721Handler.CallOpts, depositNonce, destId)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Deposit(&_ERC721Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteProposal(&_ERC721Handler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC721Handler.Contract.ExecuteProposal(&_ERC721Handler.TransactOpts, resourceID, data)
}

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerTransactor) FundERC721(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "fundERC721", tokenAddress, owner, tokenID)
}

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerSession) FundERC721(tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.Contract.FundERC721(&_ERC721Handler.TransactOpts, tokenAddress, owner, tokenID)
}

// FundERC721 is a paid mutator transaction binding the contract method 0x73542980.
//
// Solidity: function fundERC721(address tokenAddress, address owner, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) FundERC721(tokenAddress common.Address, owner common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.Contract.FundERC721(&_ERC721Handler.TransactOpts, tokenAddress, owner, tokenID)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetBurnable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetBurnable(&_ERC721Handler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetResource(&_ERC721Handler.TransactOpts, resourceID, contractAddress)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vaultAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactor) SetVault(opts *bind.TransactOpts, vaultAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "setVault", vaultAddress)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vaultAddress) returns()
func (_ERC721Handler *ERC721HandlerSession) SetVault(vaultAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetVault(&_ERC721Handler.TransactOpts, vaultAddress)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vaultAddress) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) SetVault(vaultAddress common.Address) (*types.Transaction, error) {
	return _ERC721Handler.Contract.SetVault(&_ERC721Handler.TransactOpts, vaultAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.contract.Transact(opts, "withdraw", tokenAddress, recipient, tokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerSession) Withdraw(tokenAddress common.Address, recipient common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Withdraw(&_ERC721Handler.TransactOpts, tokenAddress, recipient, tokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 tokenID) returns()
func (_ERC721Handler *ERC721HandlerTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _ERC721Handler.Contract.Withdraw(&_ERC721Handler.TransactOpts, tokenAddress, recipient, tokenID)
}
