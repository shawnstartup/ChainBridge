package vault

import (
	"fmt"
	"github.com/ChainSafe/log15"
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron"
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron/api"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

const TxStatusCompleted = "COMPLETED"

type Vault struct {
	transactionApi api.TransactionApi
	log            log15.Logger
}

func NewVault(log log15.Logger) *Vault {
	viper.SetConfigFile("vault-config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %w", err))
	}

	sc := safeheron.Client{Config: safeheron.ApiConfig{
		BaseUrl:               viper.GetString("baseUrl"),
		ApiKey:                viper.GetString("apiKey"),
		RsaPrivateKey:         viper.GetString("privateKeyPemFile"),
		SafeheronRsaPublicKey: viper.GetString("safeheronPublicKeyPemFile"),
		RequestTimeout:        viper.GetInt64("requestTimeout"),
	}}

	return &Vault{
		log:            log,
		transactionApi: api.TransactionApi{Client: sc},
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
