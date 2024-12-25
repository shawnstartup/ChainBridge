package vault

import (
	"errors"
	"fmt"
	"github.com/ChainSafe/log15"
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron/api"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"math"
	"math/big"
	"strings"
)

const TxStatusCompleted = "COMPLETED"

type Vault struct {
	transactionApi    api.TransactionApi
	vaultBridgeConfig *VaultBridgeConfig
	log               log15.Logger
}

func NewVault(log log15.Logger) *Vault {
	viper.SetConfigFile("vault-config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading vault config file, %w", err))
	}

	sc := safeheron.Client{Config: safeheron.ApiConfig{
		BaseUrl:               viper.GetString("baseUrl"),
		ApiKey:                viper.GetString("apiKey"),
		RsaPrivateKey:         viper.GetString("privateKeyPemFile"),
		SafeheronRsaPublicKey: viper.GetString("safeheronPublicKeyPemFile"),
		RequestTimeout:        viper.GetInt64("requestTimeout"),
	}}

	vaultBridgeConfig := VaultBridgeConfig{
		BridgeResources: make(map[string]map[uint8]*ResourceChain),
	}

	if err := viper.UnmarshalKey("bridgeResources", &vaultBridgeConfig.BridgeResources); err != nil {
		log.Error("Unable to decode into struct", "err", err)
	}

	return &Vault{
		log:               log,
		transactionApi:    api.TransactionApi{Client: sc},
		vaultBridgeConfig: &vaultBridgeConfig,
	}
}

func (v *Vault) RetrieveTransaction(txKey string) (string, string, error) {
	oneTransactionsRequest := api.OneTransactionsRequest{
		TxKey: txKey,
	}

	var txResp api.OneTransactionsResponse
	if err := v.transactionApi.OneTransactions(oneTransactionsRequest, &txResp); err != nil {
		v.log.Error("RetrieveTransaction", "err", fmt.Errorf("failed to retrieve transaction, %w", err))
		return "", "", err
	}
	return txResp.TransactionStatus, txResp.TransactionSubStatus, nil
}

func (v *Vault) SendTransaction(destinationAddress string, coinKey string, txAmount string, failOnContract bool) (string, error) {
	createTransactionsRequest := api.CreateTransactionsRequest{
		SourceAccountKey:       viper.GetString("accountKey"),
		SourceAccountType:      "VAULT_ACCOUNT",
		DestinationAccountType: "WHITELISTING_ACCOUNT",
		DestinationAddress:     destinationAddress,
		CoinKey:                coinKey,
		TxAmount:               txAmount,
		TxFeeLevel:             "MIDDLE",
		CustomerRefId:          uuid.New().String(),
		FailOnContract:         &failOnContract,
	}

	var txKeyResult api.TxKeyResult
	if err := v.transactionApi.CreateTransactions(createTransactionsRequest, &txKeyResult); err != nil {
		v.log.Error("SendTransaction", "err", fmt.Errorf("failed to send transaction, %w", err))
		return "", err
	}

	v.log.Info("transaction has been created", "txKey", txKeyResult.TxKey)
	return txKeyResult.TxKey, nil
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

func (v *Vault) SendVaultTransaction(chainId uint8, resourceId string, destinationAddress string, txAmount *big.Int, failOnContract bool) (string, error) {
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
	v.log.Info("SendVaultTransaction", "destinationChainId", chainId, "destinationAddress", destinationAddress, "txAmount", txAmountString)

	createTransactionsRequest := api.CreateTransactionsRequest{
		SourceAccountKey:       viper.GetString("accountKey"),
		SourceAccountType:      "VAULT_ACCOUNT",
		DestinationAccountType: "WHITELISTING_ACCOUNT",
		DestinationAddress:     destinationAddress,
		CoinKey:                coinKey,
		TxAmount:               txAmountString,
		TxFeeLevel:             "MIDDLE",
		CustomerRefId:          uuid.New().String(),
		FailOnContract:         &failOnContract,
	}

	var txKeyResult api.TxKeyResult
	if err := v.transactionApi.CreateTransactions(createTransactionsRequest, &txKeyResult); err != nil {
		v.log.Error("SendTransaction", "err", fmt.Errorf("failed to send transaction, %w", err))
		return "", err
	}

	v.log.Info("transaction has been created", "txKey", txKeyResult.TxKey)
	return txKeyResult.TxKey, nil
}
