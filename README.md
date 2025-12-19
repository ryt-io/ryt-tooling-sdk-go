# Avalanche Tooling Go SDK

> **⚠️ Maintenance Mode Notice**
>
> As of December 2025, Avalanche Tooling Go SDK has entered maintenance mode. This means:
>
> - No new features will be developed by the Ava Labs team
> - GitHub issues are expected to be addressed by community members
>
> **External contributions are very much welcome!** PRs will be reviewed but may not be prioritized. We encourage the community to continue building and improving this project.

## Getting Started

### Installing
Use `go get` to retrieve the SDK to add it to your project's Go module dependencies.

	go get github.com/ava-labs/avalanche-tooling-sdk-go

To update the SDK use `go get -u` to retrieve the latest version of the SDK.

	go get -u github.com/ava-labs/avalanche-tooling-sdk-go

## Wallet Interface

The SDK provides a `Wallet` interface for interacting with Avalanche networks. The `local.LocalWallet` is its main implementation.

```go
// Create a wallet for Fuji testnet
w, err := local.NewLocalWallet(logger, network.FujiNetwork())

// Import an account from private key
acc, _ := account.ImportFromPrivateKey("myaccount", privateKey)
w.ImportAccount(acc)

// P-Chain operations (create subnet, add validators, etc.)
result, err := w.Primary().SubmitTx(ctx, types.SubmitTxParams{
    OperationType: types.CreateSubnet,
})

// EVM operations on any chain
w.SetChain("https://my-l1.example.com/rpc")  // or "C" for C-Chain
balance, err := w.Balance()
tx, receipt, err := w.WriteContract(contractAddr, nil, wallet.Method("transfer(address,uint256)", to, amount))
```

See [wallet/wallet.go](wallet/wallet.go) for the full interface definition.

### Examples

#### P-Chain Operations

- [create_subnet.go](examples/create_subnet.go) - Create a new subnet on Fuji
- [create_chain.go](examples/create_chain.go) - Create a blockchain on an existing subnet
- [convert_subnet.go](examples/convert_subnet.go) - Convert a subnet to L1
- [create_subnet_separate_steps.go](examples/create_subnet_separate_steps.go) - Create subnet using separate BuildTx, SignTx, and SendTx steps

#### EVM Operations

- [evm/erc20_deploy_and_interact.go](examples/evm/erc20_deploy_and_interact.go) - Deploy and interact with an ERC20 token contract
- [evm/query_chain_info.go](examples/evm/query_chain_info.go) - Query chain information (balance, nonce, block number)
- [evm/get_blockchain_id.go](examples/evm/get_blockchain_id.go) - Get blockchain ID using Warp precompile
- [evm/query_validator_manager.go](examples/evm/query_validator_manager.go) - Query validator manager contract

## Avalanchego Keychains

The SDK supports different keychain implementations for signing transactions:

- [Ledger](keychain/ledger/README.md) - Hardware wallet integration for secure key management
- [CubeSigner](keychain/cubesigner/README.md) - Cloud-based key management service

### Examples

#### Ledger Integration

- [ledger/create-subnet.go](examples/ledger/create-subnet.go) - Create subnet using Ledger hardware wallet
- [ledger/cross-chain-transfer.go](examples/ledger/cross-chain-transfer.go) - Cross-chain transfer using Ledger
- [ledger/validate-ledger-txs.go](examples/ledger/validate-ledger-txs.go) - Validate transactions with Ledger

#### CubeSigner Integration

- [cubesigner/create-subnet.go](examples/cubesigner/create-subnet.go) - Create subnet using CubeSigner
- [cubesigner/cross-chain-transfer.go](examples/cubesigner/cross-chain-transfer.go) - Cross-chain transfer using CubeSigner
