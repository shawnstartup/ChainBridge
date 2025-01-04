// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"fmt"
	"github.com/ChainSafe/ChainBridge/config"
	"github.com/ChainSafe/ChainBridge/vault/secrets"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/urfave/cli/v2"
)

func handleEncryptCmd(ctx *cli.Context) error {
	var err error
	var srcFlag = ""
	var dstFlag = ""
	var password []byte = nil
	if srcFlag = ctx.String(config.SrcFilepathFlag.Name); srcFlag == "" {
		return fmt.Errorf("must provide a src filepath to encrypt")
	}
	if dstFlag = ctx.String(config.DstFilepathFlag.Name); dstFlag == "" {
		return fmt.Errorf("must provide a dst filepath to store encrypted pem file")
	}
	if password = keystore.GetPassword("Enter password to encrypt keystore file:"); len(password) == 0 {
		return fmt.Errorf("must provide a password to encrypt pem file")
	}

	err = secrets.GenerateEncrypedPEMFromPath(srcFlag, dstFlag, string(password))
	if err != nil {
		return fmt.Errorf("failed to encrypt pem file: %w", err)
	}

	return nil
}

func handleDecryptCmd(ctx *cli.Context) error {
	var err error
	var srcFlag = ""
	var dstFlag = ""
	var password []byte = nil
	if srcFlag = ctx.String(config.SrcFilepathFlag.Name); srcFlag == "" {
		return fmt.Errorf("must provide a src filepath to decrypt")
	}
	if dstFlag = ctx.String(config.DstFilepathFlag.Name); dstFlag == "" {
		return fmt.Errorf("must provide a dst filepath to store decrypted pem file")
	}
	if password = keystore.GetPassword("Enter password to decrypt keystore file:"); len(password) == 0 {
		return fmt.Errorf("must provide a password to decrypt pem file")
	}

	err = secrets.GenerateEncrypedPEMFromPath(srcFlag, dstFlag, string(password))
	if err != nil {
		return fmt.Errorf("failed to encrypt pem file: %w", err)
	}

	return nil
}
