// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	erc20Handler "github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc721Handler "github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	"github.com/ChainSafe/ChainBridge/bindings/GenericHandler"
	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	"github.com/ChainSafe/ChainBridge/connections/ethereum/egs"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
	"math/big"
	"os"
	"path/filepath"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 20000000000
const DefaultBlockConfirmations = 10
const DefaultGasMultiplier = 1

const DefaultConfigPath = "./cosigner-config.json"
const DefaultKeystorePath = "./keys"

// Chain specific options
var (
	BridgeOpt             = "bridge"
	Erc20HandlerOpt       = "erc20Handler"
	Erc721HandlerOpt      = "erc721Handler"
	GenericHandlerOpt     = "genericHandler"
	MaxGasPriceOpt        = "maxGasPrice"
	GasLimitOpt           = "gasLimit"
	GasMultiplier         = "gasMultiplier"
	HttpOpt               = "http"
	StartBlockOpt         = "startBlock"
	BlockConfirmationsOpt = "blockConfirmations"
	EGSApiKey             = "egsApiKey"
	EGSSpeed              = "egsSpeed"
)

type Config struct {
	KeystorePath string `json:"keystorePath,omitempty"`
	BizPrivKey   string `json:"bizPrivKey,omitempty"`
	ApiPubKey    string `json:"apiPubKey,omitempty"`
}

// RawChainConfig is parsed directly from the config file and should be using to construct the core.ChainConfig
type RawChainConfig struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Id       string            `json:"id"`       // ChainID
	Endpoint string            `json:"endpoint"` // url for rpc endpoint
	From     string            `json:"from"`     // address of key to use
	Opts     map[string]string `json:"opts"`
}

type EthereumChain struct {
	Conn                   *connection.Connection
	BridgeContract         *bridge.Bridge
	Cfg                    *EthereumConfig
	Erc20HandlerContract   *erc20Handler.ERC20Handler
	Erc721HandlerContract  *erc721Handler.ERC721Handler
	GenericHandlerContract *GenericHandler.GenericHandler
}

type EthereumConfig struct {
	Name                   string      // Human-readable chain name
	Id                     msg.ChainId // ChainID
	Endpoint               string      // url for rpc endpoint
	From                   string      // address of key to use
	KeystorePath           string      // Location of keyfiles
	BlockstorePath         string
	FreshStart             bool // Disables loading from blockstore at start
	BridgeContract         common.Address
	Erc20HandlerContract   common.Address
	Erc721HandlerContract  common.Address
	GenericHandlerContract common.Address
	GasLimit               *big.Int
	MaxGasPrice            *big.Int
	GasMultiplier          *big.Float
	Http                   bool // Config for type of connection
	StartBlock             *big.Int
	BlockConfirmations     *big.Int
	EgsApiKey              string // API key for ethgasstation to query gas prices
	EgsSpeed               string // The speed which a transaction should be processed: average, fast, fastest. Default: fast}
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ToJSON(file string) *os.File {
	var (
		newFile *os.File
		err     error
	)

	var raw []byte
	if raw, err = json.Marshal(*c); err != nil {
		log.Warn("error marshalling json", "err", err)
		os.Exit(1)
	}

	newFile, err = os.Create(file)
	if err != nil {
		log.Warn("error creating config file", "err", err)
	}
	_, err = newFile.Write(raw)
	if err != nil {
		log.Warn("error writing to config file", "err", err)
	}

	if err := newFile.Close(); err != nil {
		log.Warn("error closing file", "err", err)
	}
	return newFile
}

func (c *Config) validate() error {
	//for _, chain := range c.Chains {
	//	if chain.Type == "" {
	//		return fmt.Errorf("required field chain.Type empty for chain %s", chain.Id)
	//	}
	//	if chain.Endpoint == "" {
	//		return fmt.Errorf("required field chain.Endpoint empty for chain %s", chain.Id)
	//	}
	//	if chain.Name == "" {
	//		return fmt.Errorf("required field chain.Name empty for chain %s", chain.Id)
	//	}
	//	if chain.Id == "" {
	//		return fmt.Errorf("required field chain.Id empty for chain %s", chain.Id)
	//	}
	//	if chain.From == "" {
	//		return fmt.Errorf("required field chain.From empty for chain %s", chain.Id)
	//	}
	//}
	return nil
}

func GetConfig(ctx *cli.Context) (*Config, error) {
	var fig Config
	path := DefaultConfigPath
	if file := ctx.String(ConfigFileFlag.Name); file != "" {
		path = file
	}
	err := loadConfig(path, &fig)
	if err != nil {
		log.Warn("err loading json file", "err", err.Error())
		return &fig, err
	}
	if ksPath := ctx.String(KeystorePathFlag.Name); ksPath != "" {
		fig.KeystorePath = ksPath
	}
	log.Debug("Loaded config", "path", path)
	err = fig.validate()
	if err != nil {
		return nil, err
	}
	return &fig, nil
}

func loadConfig(file string, config *Config) error {
	ext := filepath.Ext(file)
	fp, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	log.Debug("Loading configuration", "path", filepath.Clean(fp))

	f, err := os.Open(filepath.Clean(fp))
	if err != nil {
		return err
	}

	if ext == ".json" {
		if err = json.NewDecoder(f).Decode(&config); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unrecognized extention: %s", ext)
	}

	return nil
}

func parseChainConfig(chainCfg *core.ChainConfig) (*EthereumConfig, error) {

	config := &EthereumConfig{
		Name:                   chainCfg.Name,
		Id:                     chainCfg.Id,
		Endpoint:               chainCfg.Endpoint,
		From:                   chainCfg.From,
		KeystorePath:           chainCfg.KeystorePath,
		BlockstorePath:         chainCfg.BlockstorePath,
		FreshStart:             chainCfg.FreshStart,
		BridgeContract:         utils.ZeroAddress,
		Erc20HandlerContract:   utils.ZeroAddress,
		Erc721HandlerContract:  utils.ZeroAddress,
		GenericHandlerContract: utils.ZeroAddress,
		GasLimit:               big.NewInt(DefaultGasLimit),
		MaxGasPrice:            big.NewInt(DefaultGasPrice),
		GasMultiplier:          big.NewFloat(DefaultGasMultiplier),
		Http:                   false,
		StartBlock:             big.NewInt(0),
		BlockConfirmations:     big.NewInt(0),
		EgsApiKey:              "",
		EgsSpeed:               "",
	}

	if contract, ok := chainCfg.Opts[BridgeOpt]; ok && contract != "" {
		config.BridgeContract = common.HexToAddress(contract)
		delete(chainCfg.Opts, BridgeOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.bridge field for ethereum config")
	}

	if contract, ok := chainCfg.Opts[Erc20HandlerOpt]; ok {
		config.Erc20HandlerContract = common.HexToAddress(contract)
		delete(chainCfg.Opts, Erc20HandlerOpt)
	}

	if contract, ok := chainCfg.Opts[Erc721HandlerOpt]; ok {
		config.Erc721HandlerContract = common.HexToAddress(contract)
		delete(chainCfg.Opts, Erc721HandlerOpt)
	}

	if contract, ok := chainCfg.Opts[GenericHandlerOpt]; ok {
		config.GenericHandlerContract = common.HexToAddress(contract)
		delete(chainCfg.Opts, GenericHandlerOpt)
	}

	if gasPrice, ok := chainCfg.Opts[MaxGasPriceOpt]; ok {
		price := big.NewInt(0)
		_, pass := price.SetString(gasPrice, 10)
		if pass {
			config.MaxGasPrice = price
			delete(chainCfg.Opts, MaxGasPriceOpt)
		} else {
			return nil, errors.New("unable to parse max gas price")
		}
	}

	if gasLimit, ok := chainCfg.Opts[GasLimitOpt]; ok {
		limit := big.NewInt(0)
		_, pass := limit.SetString(gasLimit, 10)
		if pass {
			config.GasLimit = limit
			delete(chainCfg.Opts, GasLimitOpt)
		} else {
			return nil, errors.New("unable to parse gas limit")
		}
	}

	if gasMultiplier, ok := chainCfg.Opts[GasMultiplier]; ok {
		multilier := big.NewFloat(1)
		_, pass := multilier.SetString(gasMultiplier)
		if pass {
			config.GasMultiplier = multilier
			delete(chainCfg.Opts, GasMultiplier)
		} else {
			return nil, errors.New("unable to parse gasMultiplier to float")
		}
	}

	if HTTP, ok := chainCfg.Opts[HttpOpt]; ok && HTTP == "true" {
		config.Http = true
		delete(chainCfg.Opts, HttpOpt)
	} else if HTTP, ok := chainCfg.Opts[HttpOpt]; ok && HTTP == "false" {
		config.Http = false
		delete(chainCfg.Opts, HttpOpt)
	}

	if startBlock, ok := chainCfg.Opts[StartBlockOpt]; ok && startBlock != "" {
		block := big.NewInt(0)
		_, pass := block.SetString(startBlock, 10)
		if pass {
			config.StartBlock = block
			delete(chainCfg.Opts, StartBlockOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", StartBlockOpt)
		}
	}

	if blockConfirmations, ok := chainCfg.Opts[BlockConfirmationsOpt]; ok && blockConfirmations != "" {
		val := big.NewInt(DefaultBlockConfirmations)
		_, pass := val.SetString(blockConfirmations, 10)
		if pass {
			config.BlockConfirmations = val
			delete(chainCfg.Opts, BlockConfirmationsOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", BlockConfirmationsOpt)
		}
	} else {
		config.BlockConfirmations = big.NewInt(DefaultBlockConfirmations)
		delete(chainCfg.Opts, BlockConfirmationsOpt)
	}

	if gsnApiKey, ok := chainCfg.Opts[EGSApiKey]; ok && gsnApiKey != "" {
		config.EgsApiKey = gsnApiKey
		delete(chainCfg.Opts, EGSApiKey)
	}

	if speed, ok := chainCfg.Opts[EGSSpeed]; ok && speed == egs.Average || speed == egs.Fast || speed == egs.Fastest {
		config.EgsSpeed = speed
		delete(chainCfg.Opts, EGSSpeed)
	} else {
		// Default to "fast"
		config.EgsSpeed = egs.Fast
		delete(chainCfg.Opts, EGSSpeed)
	}

	if len(chainCfg.Opts) != 0 {
		return nil, fmt.Errorf("unknown Opts Encountered: %#v", chainCfg.Opts)
	}

	return config, nil
}

func InitializeEthereumChain(chainCfg *core.ChainConfig, logger log15.Logger) (*EthereumChain, error) {
	cfg, err := parseChainConfig(chainCfg)
	if err != nil {
		return nil, err
	}

	kpI, err := keystore.KeypairFromAddress(cfg.From, keystore.EthChain, cfg.KeystorePath, chainCfg.Insecure)
	if err != nil {
		return nil, err
	}
	kp, _ := kpI.(*secp256k1.Keypair)

	conn := connection.NewConnection(cfg.Endpoint, cfg.Http, kp, logger, cfg.GasLimit, cfg.MaxGasPrice, cfg.GasMultiplier, cfg.EgsApiKey, cfg.EgsSpeed)
	err = conn.Connect()
	if err != nil {
		return nil, err
	}
	err = conn.EnsureHasBytecode(cfg.BridgeContract)
	if err != nil {
		return nil, err
	}

	if cfg.Erc20HandlerContract != utils.ZeroAddress {
		err = conn.EnsureHasBytecode(cfg.Erc20HandlerContract)
		if err != nil {
			return nil, err
		}
	}

	if cfg.GenericHandlerContract != utils.ZeroAddress {
		err = conn.EnsureHasBytecode(cfg.GenericHandlerContract)
		if err != nil {
			return nil, err
		}
	}

	bridgeContract, err := bridge.NewBridge(cfg.BridgeContract, conn.Client())
	if err != nil {
		return nil, err
	}

	chainId, err := bridgeContract.ChainID(conn.CallOpts())
	if err != nil {
		return nil, err
	}

	if chainId != uint8(chainCfg.Id) {
		return nil, fmt.Errorf("chainId (%d) and configuration chainId (%d) do not match", chainId, chainCfg.Id)
	}

	erc20HandlerContract, err := erc20Handler.NewERC20Handler(cfg.Erc20HandlerContract, conn.Client())
	if err != nil {
		return nil, err
	}

	erc721HandlerContract, err := erc721Handler.NewERC721Handler(cfg.Erc721HandlerContract, conn.Client())
	if err != nil {
		return nil, err
	}

	genericHandlerContract, err := GenericHandler.NewGenericHandler(cfg.GenericHandlerContract, conn.Client())
	if err != nil {
		return nil, err
	}

	if chainCfg.LatestBlock {
		curr, err := conn.LatestBlock()
		if err != nil {
			return nil, err
		}
		cfg.StartBlock = curr
	}

	return &EthereumChain{
		Conn:                   conn,
		Cfg:                    cfg,
		BridgeContract:         bridgeContract,
		Erc20HandlerContract:   erc20HandlerContract,
		Erc721HandlerContract:  erc721HandlerContract,
		GenericHandlerContract: genericHandlerContract,
	}, nil
}
