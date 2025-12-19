// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package local

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary"

	"github.com/ava-labs/avalanche-tooling-sdk-go/account"
	"github.com/ava-labs/avalanche-tooling-sdk-go/network"
	"github.com/ava-labs/avalanche-tooling-sdk-go/wallet"
)

// ChainClients is now defined in the wallet package

// LocalWallet represents a local wallet implementation
type LocalWallet struct {
	primaryWallet     *primary.Wallet            // Avalanchego primary wallet for P/X/C operations
	accounts          map[string]account.Account // Named accounts map
	activeAccountName string                     // Currently active account name
	activeNetwork     network.Network            // Active network for operations
	seenSubnetIDs     []ids.ID                   // Subnet IDs seen for active account
}

// Ensure LocalWallet implements Wallet interface
var _ wallet.Wallet = (*LocalWallet)(nil)

// NewLocalWallet creates a new local wallet with the specified network
func NewLocalWallet(network network.Network) (*LocalWallet, error) {
	return &LocalWallet{
		primaryWallet:     nil,
		accounts:          make(map[string]account.Account),
		activeAccountName: "",
		activeNetwork:     network,
		seenSubnetIDs:     []ids.ID{},
	}, nil
}

// SetNetwork sets the active network for wallet operations
func (w *LocalWallet) SetNetwork(network network.Network) {
	w.activeNetwork = network
}

// Network returns the active network for wallet operations
func (w *LocalWallet) Network() network.Network {
	return w.activeNetwork
}

// Primary returns the interface for P/X/C chain operations
func (w *LocalWallet) Primary() wallet.PrimaryNetworkOperations {
	return &primaryNetworkOperations{localWallet: w}
}
