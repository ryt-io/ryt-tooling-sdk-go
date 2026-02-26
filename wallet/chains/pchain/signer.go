// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package pchain

import (
	"context"
	"fmt"

	"github.com/ryt-io/ryt-v2/wallet/subnet/primary"

	"github.com/ryt-io/ryt-tooling-sdk-go/wallet/types"

	avagoTxs "github.com/ryt-io/ryt-v2/vms/platformvm/txs"
)

// SignTx signs P-Chain transactions
func SignTx(wallet *primary.Wallet, params types.SignTxParams) (types.SignTxResult, error) {
	// Get the P-Chain transaction from the BuildTxResult
	pChainTx, ok := params.BuildTxResult.GetTx().(*avagoTxs.Tx)
	if !ok {
		return types.SignTxResult{}, fmt.Errorf("expected P-Chain transaction, got %T", params.BuildTxResult.GetTx())
	}

	if err := wallet.P().Signer().Sign(context.Background(), pChainTx); err != nil {
		return types.SignTxResult{}, fmt.Errorf("error signing tx: %w", err)
	}
	pChainSignResult := types.NewPChainSignTxResult(pChainTx)
	return types.SignTxResult{SignTxOutput: pChainSignResult}, nil
}
