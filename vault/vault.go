package vault

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ChainSafe/ChainBridge/safeheron"
	"github.com/ChainSafe/ChainBridge/safeheron/api"
	"github.com/ChainSafe/ChainBridge/vault/secretes"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ChainSafe/log15"
	"github.com/spf13/viper"
	"math"
	"math/big"
	"strconv"
	"strings"
)

const TxStatusCompleted = "COMPLETED"

type Vault struct {
	transactionApi      api.TransactionApi
	vaultBridgeConfig   *VaultBridgeConfig
	log                 log15.Logger
	CustomerRefIdPrefix string
}

type TxCustomerRefId struct {
	Raw                 string
	TxId                string
	CustomerRefIdPrefix string
	SrcChainId          uint8
	DstChainId          uint8
	DepositNonce        uint64
}

func NewVault(log log15.Logger) *Vault {
	viper.SetConfigFile("vault-config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading vault config file, %w", err))
	}

	pswd := keystore.GetPassword(fmt.Sprintf("Enter password for key %s:", viper.GetString("privateKeyPemFile")))
	password := string(pswd)
	apiPrivateKey, err := secretes.LoadEncrypedPrivateKeyFromPath(viper.GetString("privateKeyPemFile"), password)
	if err != nil {
		panic(fmt.Errorf("error reading encrpted privateKey pem file, %w", err))
	}

	sc := safeheron.Client{
		Config: safeheron.ApiConfig{
			BaseUrl:               viper.GetString("baseUrl"),
			ApiKey:                viper.GetString("apiKey"),
			RsaPrivateKey:         viper.GetString("privateKeyPemFile"),
			SafeheronRsaPublicKey: viper.GetString("safeheronPublicKeyPemFile"),
			RequestTimeout:        viper.GetInt64("requestTimeout"),
		},
		ApiPrivateKey: apiPrivateKey}

	vaultBridgeConfig := VaultBridgeConfig{
		BridgeResources: make(map[string]map[uint8]*ResourceChain),
	}

	if err := viper.UnmarshalKey("bridgeResources", &vaultBridgeConfig.BridgeResources); err != nil {
		log.Error("Unable to decode into struct BridgeResources", "err", err)
	}

	if err := viper.UnmarshalKey("customerRefIdPrefix", &vaultBridgeConfig.CustomerRefIdPrefix); err != nil {
		log.Error("Unable to decode into customerRefIdPrefix string", "err", err)
	}

	return &Vault{
		log:                 log,
		transactionApi:      api.TransactionApi{Client: sc},
		vaultBridgeConfig:   &vaultBridgeConfig,
		CustomerRefIdPrefix: vaultBridgeConfig.CustomerRefIdPrefix,
	}
}

func (v *Vault) MakeCustomerRefId(srcChainId uint8, dstChainId uint8, depositNonce uint64) string {
	return fmt.Sprintf("%s-src-%03d-dst-%03d-nonce-%09d", v.CustomerRefIdPrefix, srcChainId, dstChainId, depositNonce)
}

func ParseCustomerRefId(customerRefId string) (*TxCustomerRefId, error) {
	customerRefIdParams := strings.Split(customerRefId, "-")
	srcChainId, err := strconv.ParseUint(customerRefIdParams[2], 10, 8)
	if err != nil {
		return nil, errors.New("error parsing customerRefId")
	}
	dstChainId, err := strconv.ParseUint(customerRefIdParams[4], 10, 8)
	if err != nil {
		return nil, errors.New("error parsing customerRefId")
	}
	nonce, err := strconv.ParseUint(customerRefIdParams[6], 10, 8)
	if err != nil {
		return nil, errors.New("error parsing customerRefId")
	}

	return &TxCustomerRefId{
		Raw:                 customerRefId,
		TxId:                fmt.Sprintf("%s-src-%03d-dst-%03d-nonce-%09d", customerRefIdParams[0], srcChainId, dstChainId, nonce),
		CustomerRefIdPrefix: customerRefIdParams[0],
		SrcChainId:          uint8(srcChainId),
		DstChainId:          uint8(dstChainId),
		DepositNonce:        nonce,
	}, nil
}

func (v *Vault) RetrieveTransaction(txKey string, txId string) (string, string, error) {
	oneTransactionsRequest := api.OneTransactionsRequest{
		CustomerRefId: txId,
		TxKey:         txKey,
	}

	var txResp api.OneTransactionsResponse
	if err := v.transactionApi.OneTransactions(oneTransactionsRequest, &txResp); err != nil {
		v.log.Error("RetrieveTransaction", "err", fmt.Errorf("failed to retrieve transaction, %w", err))
		return "", "", err
	}
	return txResp.TransactionStatus, txResp.TransactionSubStatus, nil
}

func AsStringFromFloat(precision int, amount *big.Float) (string, error) {
	fmtString := fmt.Sprintf("%%.%df", precision)
	return fmt.Sprintf(strings.TrimRight(strings.TrimRight(fmt.Sprintf(fmtString, amount), "0"), ".")), nil
}

func Wei2ethStr(amount *big.Int, decimal int) string {
	ethValue := WeiToEther(amount, decimal)
	stringFromFloat, err := AsStringFromFloat(6, ethValue)
	if err != nil {
		return ""
	}
	return stringFromFloat
}

func WeiToEther(wei *big.Int, decimal int) *big.Float {
	ether := new(big.Float)
	ether.SetInt(wei)
	ethValue := new(big.Float).Quo(ether, big.NewFloat(math.Pow(10, float64(decimal))))
	return ethValue
}

type TransactionNote struct {
	CustomerRefId string
}

func (v *Vault) SendVaultTransaction(chainId uint8, resourceId string, destinationAddress string, customerRefId string, txAmount *big.Int, failOnContract bool) (string, error) {
	var decimal int
	coinKey := ""
	resourceChain, ok := v.vaultBridgeConfig.BridgeResources[resourceId][chainId]
	if !ok {
		return "", errors.New("resource not found")
	}
	decimal = resourceChain.CoinDecimal
	coinKey = resourceChain.CoinKey

	if decimal == 0 || coinKey == "" {
		return "", errors.New("invalid coin")
	}

	txAmountString := Wei2ethStr(txAmount, decimal)
	if txAmountString == "" {
		return "", errors.New("transaction amount is zero")
	}

	txNote := TransactionNote{
		CustomerRefId: customerRefId,
	}
	txNoteJson, _ := json.Marshal(txNote)
	v.log.Info("SendVaultTransaction", "destinationChainId", chainId, "destinationAddress", destinationAddress, "txAmount", txAmountString)

	createTransactionsRequest := api.CreateTransactionsRequest{
		SourceAccountKey:       viper.GetString("accountKey"),
		SourceAccountType:      "VAULT_ACCOUNT",
		DestinationAccountType: "WHITELISTING_ACCOUNT",
		DestinationAddress:     destinationAddress,
		CoinKey:                coinKey,
		TxAmount:               txAmountString,
		TxFeeLevel:             "MIDDLE",
		CustomerRefId:          customerRefId,
		FailOnContract:         &failOnContract,
		Note:                   string(txNoteJson),
	}

	var txKeyResult api.TxKeyResult
	if err := v.transactionApi.CreateTransactions(createTransactionsRequest, &txKeyResult); err != nil {
		v.log.Error("SendTransaction", "err", fmt.Errorf("failed to send transaction, %w", err))
		return "", err
	}

	v.log.Info("transaction has been created", "txKey", txKeyResult.TxKey)
	return txKeyResult.TxKey, nil
}
