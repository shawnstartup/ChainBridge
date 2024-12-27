// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20Handler

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

// ERC20HandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type ERC20HandlerDepositRecord struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}

// ERC20HandlerMetaData contains all meta data concerning the ERC20Handler contract.
var ERC20HandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"burnableContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFromERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"internalType\":\"structERC20Handler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620025fa380380620025fa833981810160405281019062000037919062000489565b81518351146200007e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000759062000632565b60405180910390fd5b836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b8351811015620001165762000108848281518110620000df57fe5b6020026020010151848381518110620000f457fe5b60200260200101516200016560201b60201c565b8080600101915050620000c4565b5060008090505b81518110156200015a576200014c8282815181106200013857fe5b60200260200101516200025760201b60201c565b80806001019150506200011d565b505050505062000757565b806002600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16620002e6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002dd9062000610565b60405180910390fd5b6001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600081519050620003528162000723565b92915050565b600082601f8301126200036a57600080fd5b8151620003816200037b8262000682565b62000654565b91508181835260208401935060208101905083856020840282011115620003a757600080fd5b60005b83811015620003db5781620003c0888262000341565b845260208401935060208301925050600181019050620003aa565b5050505092915050565b600082601f830112620003f757600080fd5b81516200040e6200040882620006ab565b62000654565b915081818352602084019350602081019050838560208402820111156200043457600080fd5b60005b838110156200046857816200044d888262000472565b84526020840193506020830192505060018101905062000437565b5050505092915050565b60008151905062000483816200073d565b92915050565b60008060008060808587031215620004a057600080fd5b6000620004b08782880162000341565b945050602085015167ffffffffffffffff811115620004ce57600080fd5b620004dc87828801620003e5565b935050604085015167ffffffffffffffff811115620004fa57600080fd5b620005088782880162000358565b925050606085015167ffffffffffffffff8111156200052657600080fd5b620005348782880162000358565b91505092959194509250565b60006200054f602483620006d4565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000620005b7603c83620006d4565b91507f696e697469616c5265736f7572636549447320616e6420696e697469616c436f60008301527f6e7472616374416464726573736573206c656e206d69736d61746368000000006020830152604082019050919050565b600060208201905081810360008301526200062b8162000540565b9050919050565b600060208201905081810360008301526200064d81620005a8565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200067857600080fd5b8060405250919050565b600067ffffffffffffffff8211156200069a57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115620006c357600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620006f28262000703565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200072e81620006e5565b81146200073a57600080fd5b50565b6200074881620006f9565b81146200075457600080fd5b50565b611e9380620007676000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806395601f0911610097578063c8ba6c8711610066578063c8ba6c87146102c3578063d9caed12146102f3578063dd71105d1461030f578063e248cff21461032b57610100565b806395601f091461023d578063b8fa373614610259578063ba484c0914610275578063c3b65816146102a557610100565b80634402027f116100d35780634402027f1461018b5780636817031b146101c15780636a70d081146101dd5780637f79bea81461020d57610100565b806307b7ed99146101055780630a6d55d814610121578063318c136e1461015157806338995da91461016f575b600080fd5b61011f600480360381019061011a91906114ab565b610347565b005b61013b600480360381019061013691906115af565b61035b565b6040516101489190611adf565b60405180910390f35b61015961038e565b6040516101669190611adf565b60405180910390f35b6101896004803603810190610184919061166c565b6103b3565b005b6101a560048036038101906101a0919061173a565b610750565b6040516101b89796959493929190611b5a565b60405180910390f35b6101db60048036038101906101d691906114ab565b610891565b005b6101f760048036038101906101f291906114ab565b6108dd565b6040516102049190611bd0565b60405180910390f35b610227600480360381019061022291906114ab565b6108fd565b6040516102349190611bd0565b60405180910390f35b61025760048036038101906102529190611537565b61091d565b005b610273600480360381019061026e91906115d8565b610934565b005b61028f600480360381019061028a91906116fe565b61094a565b60405161029c9190611ca6565b60405180910390f35b6102ad610b3f565b6040516102ba9190611adf565b60405180910390f35b6102dd60048036038101906102d891906114ab565b610b65565b6040516102ea9190611beb565b60405180910390f35b61030d60048036038101906103089190611537565b610b7d565b005b610329600480360381019061032491906114d4565b610b95565b005b61034560048036038101906103409190611614565b610bad565b005b61034f610d21565b61035881610db2565b50565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6103bb610d21565b606060008060c4359150604051925060e4359050808301602001604052600e360360e484376000600260008b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166104a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049b90611c86565b60405180910390fd5b600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561050657610501818885610e99565b61059d565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461058f5761058a8188600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1686610b95565b61059c565b61059b81883086610f11565b5b5b6040518060e001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018360ff1681526020018a60ff1681526020018b81526020018581526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600660008b60ff1660ff16815260200190815260200160002060008a67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff16021790555060408201518160000160156101000a81548160ff021916908360ff1602179055506060820151816001015560808201518160020190805190602001906106ef9291906112cc565b5060a08201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816004015590505050505050505050505050565b6006602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16908060000160159054906101000a900460ff1690806001015490806002018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561085b5780601f106108305761010080835404028352916020019161085b565b820191906000526020600020905b81548152906001019060200180831161083e57829003601f168201915b5050505050908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905087565b610899610d21565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60056020528060005260406000206000915054906101000a900460ff1681565b60046020528060005260406000206000915054906101000a900460ff1681565b600083905061092e81843085610f29565b50505050565b61093c610d21565b6109468282610fb2565b5050565b61095261134c565b600660008360ff1660ff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016000820160159054906101000a900460ff1660ff1660ff16815260200160018201548152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ace5780601f10610aa357610100808354040283529160200191610ace565b820191906000526020600020905b815481529060010190602001808311610ab157829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481525050905092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915090505481565b610b85610d21565b610b908383836110a4565b505050565b6000849050610ba681858585610f29565b5050505050565b610bb5610d21565b60006060606435915060405190506084358082016020016040526084360360848337506000806002600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060208301519150600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610ca4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9b90611c86565b60405180910390fd5b600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610d0957610d04818360601c866110ba565b610d18565b610d17818360601c866110a4565b5b50505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610db0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610da790611c26565b60405180910390fd5b565b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610e3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e3590611c46565b60405180910390fd5b6001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166379cc679084846040518363ffffffff1660e01b8152600401610ed9929190611b31565b600060405180830381600087803b158015610ef357600080fd5b505af1158015610f07573d6000803e3d6000fd5b5050505050505050565b6000849050610f2281858585610f29565b5050505050565b610fac846323b872dd60e01b858585604051602401610f4a93929190611afa565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611132565b50505050565b806002600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008390506110b4818484611246565b50505050565b60008390508073ffffffffffffffffffffffffffffffffffffffff166340c10f1984846040518363ffffffff1660e01b81526004016110fa929190611b31565b600060405180830381600087803b15801561111457600080fd5b505af1158015611128573d6000803e3d6000fd5b5050505050505050565b600060608373ffffffffffffffffffffffffffffffffffffffff168360405161115b9190611ac8565b6000604051808303816000865af19150503d8060008114611198576040519150601f19603f3d011682016040523d82523d6000602084013e61119d565b606091505b5091509150816111e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111d990611c66565b60405180910390fd5b60008151111561124057808060200190518101906112009190611586565b61123f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123690611c06565b60405180910390fd5b5b50505050565b6112c78363a9059cbb60e01b8484604051602401611265929190611b31565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611132565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061130d57805160ff191683800117855561133b565b8280016001018555821561133b579182015b8281111561133a57825182559160200191906001019061131f565b5b50905061134891906113be565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600060ff1681526020016000801916815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6113e091905b808211156113dc5760008160009055506001016113c4565b5090565b90565b6000813590506113f281611dd3565b92915050565b60008151905061140781611dea565b92915050565b60008135905061141c81611e01565b92915050565b60008083601f84011261143457600080fd5b8235905067ffffffffffffffff81111561144d57600080fd5b60208301915083600182028301111561146557600080fd5b9250929050565b60008135905061147b81611e18565b92915050565b60008135905061149081611e2f565b92915050565b6000813590506114a581611e46565b92915050565b6000602082840312156114bd57600080fd5b60006114cb848285016113e3565b91505092915050565b600080600080608085870312156114ea57600080fd5b60006114f8878288016113e3565b9450506020611509878288016113e3565b935050604061151a878288016113e3565b925050606061152b8782880161146c565b91505092959194509250565b60008060006060848603121561154c57600080fd5b600061155a868287016113e3565b935050602061156b868287016113e3565b925050604061157c8682870161146c565b9150509250925092565b60006020828403121561159857600080fd5b60006115a6848285016113f8565b91505092915050565b6000602082840312156115c157600080fd5b60006115cf8482850161140d565b91505092915050565b600080604083850312156115eb57600080fd5b60006115f98582860161140d565b925050602061160a858286016113e3565b9150509250929050565b60008060006040848603121561162957600080fd5b60006116378682870161140d565b935050602084013567ffffffffffffffff81111561165457600080fd5b61166086828701611422565b92509250509250925092565b60008060008060008060a0878903121561168557600080fd5b600061169389828a0161140d565b96505060206116a489828a01611496565b95505060406116b589828a01611481565b94505060606116c689828a016113e3565b935050608087013567ffffffffffffffff8111156116e357600080fd5b6116ef89828a01611422565b92509250509295509295509295565b6000806040838503121561171157600080fd5b600061171f85828601611481565b925050602061173085828601611496565b9150509250929050565b6000806040838503121561174d57600080fd5b600061175b85828601611496565b925050602061176c85828601611481565b9150509250929050565b61177f81611d1c565b82525050565b61178e81611d1c565b82525050565b61179d81611d2e565b82525050565b6117ac81611d3a565b82525050565b6117bb81611d3a565b82525050565b60006117cc82611cd3565b6117d68185611d00565b93506117e6818560208601611d8f565b80840191505092915050565b60006117fd82611cc8565b6118078185611cde565b9350611817818560208601611d8f565b61182081611dc2565b840191505092915050565b600061183682611cc8565b6118408185611cef565b9350611850818560208601611d8f565b61185981611dc2565b840191505092915050565b6000611871602083611d0b565b91507f45524332303a206f7065726174696f6e20646964206e6f7420737563636565646000830152602082019050919050565b60006118b1601e83611d0b565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006118f1602483611d0b565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611957601283611d0b565b91507f45524332303a2063616c6c206661696c656400000000000000000000000000006000830152602082019050919050565b6000611997602883611d0b565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b600060e083016000830151611a086000860182611776565b506020830151611a1b6020860182611aaa565b506040830151611a2e6040860182611aaa565b506060830151611a4160608601826117a3565b5060808301518482036080860152611a5982826117f2565b91505060a0830151611a6e60a0860182611776565b5060c0830151611a8160c0860182611a8c565b508091505092915050565b611a9581611d64565b82525050565b611aa481611d64565b82525050565b611ab381611d82565b82525050565b611ac281611d82565b82525050565b6000611ad482846117c1565b915081905092915050565b6000602082019050611af46000830184611785565b92915050565b6000606082019050611b0f6000830186611785565b611b1c6020830185611785565b611b296040830184611a9b565b949350505050565b6000604082019050611b466000830185611785565b611b536020830184611a9b565b9392505050565b600060e082019050611b6f600083018a611785565b611b7c6020830189611ab9565b611b896040830188611ab9565b611b9660608301876117b2565b8181036080830152611ba8818661182b565b9050611bb760a0830185611785565b611bc460c0830184611a9b565b98975050505050505050565b6000602082019050611be56000830184611794565b92915050565b6000602082019050611c0060008301846117b2565b92915050565b60006020820190508181036000830152611c1f81611864565b9050919050565b60006020820190508181036000830152611c3f816118a4565b9050919050565b60006020820190508181036000830152611c5f816118e4565b9050919050565b60006020820190508181036000830152611c7f8161194a565b9050919050565b60006020820190508181036000830152611c9f8161198a565b9050919050565b60006020820190508181036000830152611cc081846119f0565b905092915050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611d2782611d44565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b60005b83811015611dad578082015181840152602081019050611d92565b83811115611dbc576000848401525b50505050565b6000601f19601f8301169050919050565b611ddc81611d1c565b8114611de757600080fd5b50565b611df381611d2e565b8114611dfe57600080fd5b50565b611e0a81611d3a565b8114611e1557600080fd5b50565b611e2181611d64565b8114611e2c57600080fd5b50565b611e3881611d6e565b8114611e4357600080fd5b50565b611e4f81611d82565b8114611e5a57600080fd5b5056fea2646970667358221220550df429a3140650e4f27823e91fff893194c889ed3c1879124c3b15ec6eefcc64736f6c63430006040033",
}

// ERC20HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20HandlerMetaData.ABI instead.
var ERC20HandlerABI = ERC20HandlerMetaData.ABI

// ERC20HandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20HandlerMetaData.Bin instead.
var ERC20HandlerBin = ERC20HandlerMetaData.Bin

// DeployERC20Handler deploys a new Ethereum contract, binding an instance of ERC20Handler to it.
func DeployERC20Handler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableContractAddresses []common.Address) (common.Address, *types.Transaction, *ERC20Handler, error) {
	parsed, err := ERC20HandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20HandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses, burnableContractAddresses)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// ERC20Handler is an auto generated Go binding around an Ethereum contract.
type ERC20Handler struct {
	ERC20HandlerCaller     // Read-only binding to the contract
	ERC20HandlerTransactor // Write-only binding to the contract
	ERC20HandlerFilterer   // Log filterer for contract events
}

// ERC20HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20HandlerSession struct {
	Contract     *ERC20Handler     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20HandlerCallerSession struct {
	Contract *ERC20HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20HandlerTransactorSession struct {
	Contract     *ERC20HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20HandlerRaw struct {
	Contract *ERC20Handler // Generic contract binding to access the raw methods on
}

// ERC20HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20HandlerCallerRaw struct {
	Contract *ERC20HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20HandlerTransactorRaw struct {
	Contract *ERC20HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Handler creates a new instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20Handler(address common.Address, backend bind.ContractBackend) (*ERC20Handler, error) {
	contract, err := bindERC20Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Handler{ERC20HandlerCaller: ERC20HandlerCaller{contract: contract}, ERC20HandlerTransactor: ERC20HandlerTransactor{contract: contract}, ERC20HandlerFilterer: ERC20HandlerFilterer{contract: contract}}, nil
}

// NewERC20HandlerCaller creates a new read-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerCaller(address common.Address, caller bind.ContractCaller) (*ERC20HandlerCaller, error) {
	contract, err := bindERC20Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerCaller{contract: contract}, nil
}

// NewERC20HandlerTransactor creates a new write-only instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20HandlerTransactor, error) {
	contract, err := bindERC20Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerTransactor{contract: contract}, nil
}

// NewERC20HandlerFilterer creates a new log filterer instance of ERC20Handler, bound to a specific deployed contract.
func NewERC20HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20HandlerFilterer, error) {
	contract, err := bindERC20Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20HandlerFilterer{contract: contract}, nil
}

// bindERC20Handler binds a generic wrapper to an already deployed contract.
func bindERC20Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20HandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.ERC20HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ERC20HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Handler *ERC20HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Handler *ERC20HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Handler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _ERC20Handler.Contract.BridgeAddress(&_ERC20Handler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.BurnList(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_ERC20Handler *ERC20HandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _ERC20Handler.Contract.ContractWhitelist(&_ERC20Handler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	outstruct := new(struct {
		TokenAddress                   common.Address
		LenDestinationRecipientAddress uint8
		DestinationChainID             uint8
		ResourceID                     [32]byte
		DestinationRecipientAddress    []byte
		Depositer                      common.Address
		Amount                         *big.Int
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
	outstruct.Amount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _ERC20Handler.Contract.DepositRecords(&_ERC20Handler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_ERC20Handler *ERC20HandlerCallerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _ERC20Handler.Contract.DepositRecords(&_ERC20Handler.CallOpts, arg0, arg1)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _ERC20Handler.Contract.ResourceIDToTokenContractAddress(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_ERC20Handler *ERC20HandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _ERC20Handler.Contract.TokenContractAddressToResourceID(&_ERC20Handler.CallOpts, arg0)
}

// VaultAddress is a free data retrieval call binding the contract method 0xc3b65816.
//
// Solidity: function _vaultAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCaller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "_vaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAddress is a free data retrieval call binding the contract method 0xc3b65816.
//
// Solidity: function _vaultAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerSession) VaultAddress() (common.Address, error) {
	return _ERC20Handler.Contract.VaultAddress(&_ERC20Handler.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0xc3b65816.
//
// Solidity: function _vaultAddress() view returns(address)
func (_ERC20Handler *ERC20HandlerCallerSession) VaultAddress() (common.Address, error) {
	return _ERC20Handler.Contract.VaultAddress(&_ERC20Handler.CallOpts)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_ERC20Handler *ERC20HandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce uint64, destId uint8) (ERC20HandlerDepositRecord, error) {
	var out []interface{}
	err := _ERC20Handler.contract.Call(opts, &out, "getDepositRecord", depositNonce, destId)

	if err != nil {
		return *new(ERC20HandlerDepositRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20HandlerDepositRecord)).(*ERC20HandlerDepositRecord)

	return out0, err

}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_ERC20Handler *ERC20HandlerSession) GetDepositRecord(depositNonce uint64, destId uint8) (ERC20HandlerDepositRecord, error) {
	return _ERC20Handler.Contract.GetDepositRecord(&_ERC20Handler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_ERC20Handler *ERC20HandlerCallerSession) GetDepositRecord(depositNonce uint64, destId uint8) (ERC20HandlerDepositRecord, error) {
	return _ERC20Handler.Contract.GetDepositRecord(&_ERC20Handler.CallOpts, depositNonce, destId)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Deposit(&_ERC20Handler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteProposal(&_ERC20Handler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _ERC20Handler.Contract.ExecuteProposal(&_ERC20Handler.TransactOpts, resourceID, data)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) FundERC20(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "fundERC20", tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.FundERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// FundERC20 is a paid mutator transaction binding the contract method 0x95601f09.
//
// Solidity: function fundERC20(address tokenAddress, address owner, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) FundERC20(tokenAddress common.Address, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.FundERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, amount)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetBurnable(&_ERC20Handler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetResource(&_ERC20Handler.TransactOpts, resourceID, contractAddress)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vaultAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactor) SetVault(opts *bind.TransactOpts, vaultAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "setVault", vaultAddress)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vaultAddress) returns()
func (_ERC20Handler *ERC20HandlerSession) SetVault(vaultAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetVault(&_ERC20Handler.TransactOpts, vaultAddress)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vaultAddress) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) SetVault(vaultAddress common.Address) (*types.Transaction, error) {
	return _ERC20Handler.Contract.SetVault(&_ERC20Handler.TransactOpts, vaultAddress)
}

// TransferFromERC20 is a paid mutator transaction binding the contract method 0xdd71105d.
//
// Solidity: function transferFromERC20(address tokenAddress, address owner, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) TransferFromERC20(opts *bind.TransactOpts, tokenAddress common.Address, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "transferFromERC20", tokenAddress, owner, recipient, amount)
}

// TransferFromERC20 is a paid mutator transaction binding the contract method 0xdd71105d.
//
// Solidity: function transferFromERC20(address tokenAddress, address owner, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) TransferFromERC20(tokenAddress common.Address, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.TransferFromERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, recipient, amount)
}

// TransferFromERC20 is a paid mutator transaction binding the contract method 0xdd71105d.
//
// Solidity: function transferFromERC20(address tokenAddress, address owner, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) TransferFromERC20(tokenAddress common.Address, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.TransferFromERC20(&_ERC20Handler.TransactOpts, tokenAddress, owner, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.contract.Transact(opts, "withdraw", tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) returns()
func (_ERC20Handler *ERC20HandlerTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Handler.Contract.Withdraw(&_ERC20Handler.TransactOpts, tokenAddress, recipient, amount)
}
