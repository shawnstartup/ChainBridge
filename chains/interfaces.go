// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package chains

import (
	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/msg"
)

type Router interface {
	Send(message msg.Message) error
}

type ProposalStatus int

const (
	Inactive ProposalStatus = iota // 从0开始递增
	Active
	Passed
	Executed
	Cancelled
)

var proposalStatusName = [...]string{"Inactive", "Active", "Passed", "Executed", "Cancelled"}

func (s ProposalStatus) String() string {
	if s < 0 || int(s) >= len(proposalStatusName) {
		return "Unknown"
	}
	return proposalStatusName[s]
}

type ChainHandler interface {
	core.Chain
	QueryErc20DepositRecord(destId msg.ChainId, nonce msg.Nonce) (ERC20Handler.ERC20HandlerDepositRecord, error)
	HandleErc20DepositedRecord(erc20DepositRecored ERC20Handler.ERC20HandlerDepositRecord, nonce msg.Nonce) error
	GetProposal(record ERC20Handler.ERC20HandlerDepositRecord, sourceChainId uint8, nonce msg.Nonce) (Bridge.BridgeProposal, error)
	GetVaultProposal(record ERC20Handler.ERC20HandlerDepositRecord, sourceChainId uint8, nonce msg.Nonce) (Bridge.BridgeVaultProposal, error)
}
