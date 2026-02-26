// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package precompiles

import (
	"github.com/ryt-io/ryt-v2/ids"
	"github.com/ryt-io/libevm/common"

	"github.com/ryt-io/ryt-tooling-sdk-go/evm/contract"
)

func WarpPrecompileGetBlockchainID(
	rpcURL string,
) (ids.ID, error) {
	out, err := contract.CallToMethod(
		rpcURL,
		common.Address{},
		WarpPrecompile,
		"getBlockchainID()->(bytes32)",
		nil,
	)
	if err != nil {
		return ids.Empty, err
	}
	return contract.GetSmartContractCallResult[[32]byte]("getBlockchainID", out)
}
