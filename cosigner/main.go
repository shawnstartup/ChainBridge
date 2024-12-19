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
	"github.com/ChainSafe/ChainBridge/cosigner/config"
	"github.com/Safeheron/safeheron-api-sdk-go/safeheron/cosigner"
	"io"
	"net/http"
	"os"
	"time"

	log "github.com/ChainSafe/log15"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"strconv"
)

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

//var importFlags = []cli.Flag{
//	config.Sr25519Flag,
//	config.Secp256k1Flag,
//	config.PasswordFlag,
//}

var (
	Version = "0.0.1"
)

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
		}

		var coSignerCallBack cosigner.CoSignerCallBack

		if err := json.Unmarshal(body, &coSignerCallBack); err != nil {
			http.Error(w, err.Error(), 400)
		}

		coSignerBizContent, _ := coSignerConverter.RequestConvert(coSignerCallBack)
		//According to different types of CoSignerCallBack, the customer handles the corresponding type of business logic.
		log.Info("coSignerBizContent: %s", coSignerBizContent)

		var coSignerResponse cosigner.CoSignerResponse
		coSignerResponse.Approve = true
		coSignerResponse.TxKey = coSignerCallBack.Key
		encryptResponse, _ := coSignerConverter.ResponseConverterWithNewCryptoType(coSignerResponse)
		log.Info("encryptResponse: %s", encryptResponse)
		resp, err := json.Marshal(encryptResponse)
		log.Info(string(resp), err)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
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
