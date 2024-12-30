// Copyright 2024 EMC
// SPDX-License-Identifier: LGPL-3.0-only
/*
Provides the command-line interface for the chainbridge cosigner application.

*/
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	bridgeConfig "github.com/ChainSafe/ChainBridge/config"
	"github.com/ChainSafe/ChainBridge/cosigner/config"
	vault "github.com/ChainSafe/ChainBridge/vault"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/msg"
	log "github.com/ChainSafe/log15"
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron/cosigner"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var VaultInactiveStatus uint8 = 0
var VaultActiveStatus uint8 = 1
var VaultPassedStatus uint8 = 2
var VaultExecutedStatus uint8 = 3

var coSignerConverter cosigner.CoSignerConverter

var app = cli.NewApp()

var cliFlags = []cli.Flag{
	config.ConfigFileFlag,
	config.KeystorePathFlag,
	config.MetricsFlag,
	config.MetricsPort,
	config.VerbosityFlag,
}

var generateFlags = []cli.Flag{
	config.PasswordFlag,
	config.Sr25519Flag,
	config.Secp256k1Flag,
}

var devFlags = []cli.Flag{}
var chains = make(map[uint8]*config.EthereumChain)

//var importFlags = []cli.Flag{
//	config.Sr25519Flag,
//	config.Secp256k1Flag,
//	config.PasswordFlag,
//}

var (
	Version = "0.0.1"
)

type CustomerContent struct {
	TxKey                  string      `json:"txKey"`
	CoinKey                string      `json:"coinKey"`
	TxAmount               string      `json:"txAmount"`
	TransactionType        string      `json:"transactionType"`
	SourceAccountKey       string      `json:"sourceAccountKey"`
	SourceAccountType      string      `json:"sourceAccountType"`
	SourceAddress          string      `json:"sourceAddress"`
	DestinationAccountKey  string      `json:"destinationAccountKey"`
	DestinationAccountType string      `json:"destinationAccountType"`
	DestinationAddress     string      `json:"destinationAddress"`
	VaultTxDirection       string      `json:"vaultTxDirection"`
	TransactionStatus      string      `json:"transactionStatus"`
	CreateTime             json.Number `json:"createTime"`
	CreatedByUserKey       string      `json:"createdByUserKey"`
	TxFee                  string      `json:"txFee"`
	FeeCoinKey             string      `json:"feeCoinKey"`
	CustomerRefId          string      `json:"customerRefId"`
	AmlLock                string      `json:"amlLock"`
	TransactionDirection   string      `json:"transactionDirection"`
	EstimateFee            string      `json:"estimateFee"`
}
type CoSignerCallBackBizContent struct {
	Type            string          `json:"type"`
	CustomerContent CustomerContent `json:"customerContent"`
}

// init initializes CLI
func init() {
	app.Action = run
	app.Copyright = "Copyright 2019 ChainSafe Systems Authors"
	app.Name = "chainbridge-cosigner-callback"
	app.Usage = "ChainBridge"
	app.Authors = []*cli.Author{{Name: "ChainSafe Systems 2019"}}
	app.Version = Version
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{}

	app.Flags = append(app.Flags, cliFlags...)
	app.Flags = append(app.Flags, devFlags...)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func startLogger(ctx *cli.Context) error {
	logger := log.Root()
	handler := logger.GetHandler()
	var lvl log.Lvl

	if lvlToInt, err := strconv.Atoi(ctx.String(config.VerbosityFlag.Name)); err == nil {
		lvl = log.Lvl(lvlToInt)
	} else if lvl, err = log.LvlFromString(ctx.String(config.VerbosityFlag.Name)); err != nil {
		return err
	}
	log.Root().SetHandler(log.LvlFilterHandler(lvl, handler))

	return nil
}

func run(ctx *cli.Context) error {
	err := startLogger(ctx)
	if err != nil {
		return err
	}

	log.Info("Starting ChainBridge CoSigner Callback server...")

	cfg, err := config.GetConfig(ctx)
	if err != nil {
		return err
	}

	log.Debug("Config on initialization...", "config", *cfg)

	bridgeCfg, err := bridgeConfig.GetConfig(ctx)
	if err != nil {
		return err
	}

	for _, chain := range bridgeCfg.Chains {
		chainId, errr := strconv.Atoi(chain.Id)
		if errr != nil {
			return errr
		}
		chainConfig := &core.ChainConfig{
			Name:           chain.Name,
			Id:             msg.ChainId(chainId),
			Endpoint:       chain.Endpoint,
			From:           chain.From,
			KeystorePath:   cfg.KeystorePath,
			Insecure:       false,
			BlockstorePath: ctx.String(bridgeConfig.BlockstorePathFlag.Name),
			FreshStart:     ctx.Bool(bridgeConfig.FreshStartFlag.Name),
			LatestBlock:    ctx.Bool(bridgeConfig.LatestBlockFlag.Name),
			Opts:           chain.Opts,
		}

		logger := log.Root().New("chain", chainConfig.Name)

		if chain.Type == "ethereum" {
			newChain, err := config.InitializeEthereumChain(chainConfig, logger)
			if err != nil {
				return err
			}
			chains[uint8(chainConfig.Id)] = newChain
		} else if chain.Type == "substrate" {
			return errors.New("unrecognized Chain Type")
		} else {
			return errors.New("unrecognized Chain Type")
		}
	}

	coSignerConverter = cosigner.CoSignerConverter{Config: cosigner.CoSignerConfig{
		ApiPubKey:  cfg.ApiPubKey,
		BizPrivKey: cfg.BizPrivKey,
	}}

	//var ks = cfg.KeystorePath

	http.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		resp := fmt.Sprintf("%s\n", time.Now().String())
		w.Write([]byte(resp))
	})

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/audit", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		_, err = json.Marshal(body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		var coSignerCallBack cosigner.CoSignerCallBack

		if err := json.Unmarshal(body, &coSignerCallBack); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		coSignerBizContent, _ := coSignerConverter.RequestConvert(coSignerCallBack)
		//According to different types of CoSignerCallBack, the customer handles the corresponding type of business logic.
		log.Info(fmt.Sprintf("coSignerBizContent: %s", coSignerBizContent))

		var coSignerCallBackBizContent CoSignerCallBackBizContent

		if err := json.Unmarshal([]byte(coSignerBizContent), &coSignerCallBackBizContent); err != nil {
			http.Error(w, err.Error(), 400)
			log.Error(fmt.Sprintf("Unmarshal coSignerBizContent err : %s", err.Error()))
			return
		}

		var coSignerResponse cosigner.CoSignerResponse
		coSignerResponse.TxKey = coSignerCallBackBizContent.CustomerContent.TxKey
		coSignerResponse.Approve = false

		customerRefId := coSignerCallBackBizContent.CustomerContent.CustomerRefId
		log.Info("CustomerContent", "TxKey", coSignerCallBackBizContent.CustomerContent.TxKey, "customerRefId", customerRefId)

		txCustomerRefId, err := vault.ParseCustomerRefId(customerRefId)
		if err != nil {
			log.Error(fmt.Sprintf("Unmarshal coSignerBizContent err : %s", err.Error()))
		} else {
			chainId := txCustomerRefId.DstChainId
			chain, ok := chains[chainId]
			if !ok {
				log.Error("resource not found")
			} else {
				proposalRecord, err := chain.BridgeContract.GetTxProposalRecord(chain.Conn.CallOpts(), txCustomerRefId.Raw)
				if err != nil {
					log.Error("proposalRecord not found", "TxKey", coSignerCallBackBizContent.CustomerContent.TxKey, "customerRefId", customerRefId)
				} else {
					prop, err := chain.BridgeContract.GetVaultProposal(chain.Conn.CallOpts(), proposalRecord.OriginChainID, proposalRecord.DepositNonce, proposalRecord.DataHash)
					if err != nil {
						log.Error("Failed to check vault proposal existence", "TxKey", coSignerCallBackBizContent.CustomerContent.TxKey, "customerRefId", customerRefId)
					} else {
						if prop.Status != VaultPassedStatus {
							log.Error("Failed to check vaultProposalStatus", "TxKey", coSignerCallBackBizContent.CustomerContent.TxKey, "customerRefId", "vaultProposalStatus", customerRefId, prop.Status)
						} else {
							coSignerResponse.Approve = true
						}
					}
				}
			}
		}
		//chainId = coSignerCallBackBizContent.CustomerContent.

		encryptResponse, _ := coSignerConverter.ResponseConverterWithNewCryptoType(coSignerResponse)
		log.Debug(fmt.Sprintf("encryptResponse: %s", encryptResponse))
		resp, err := json.Marshal(encryptResponse)
		if err != nil {
			http.Error(w, err.Error(), 400)
			log.Error(fmt.Sprintf("Marshal encryptResponse err : %s", err.Error()))
			return
		}
		log.Info(string(resp))
		w.Write(resp)
	})

	err = http.ListenAndServe(fmt.Sprintf(":%d", 60001), nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Info("Health status server is shutting down", err)
	} else {
		log.Error("Error serving metrics", "err", err)
	}

	return nil
}
