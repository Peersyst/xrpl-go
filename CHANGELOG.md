# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v1.0.0]

### Added

#### binary-codec

- Updated `definitions`.
- New `DecodeLedgerData` function.
- `Quality` encoding/decoding functions.
- New `XChainBridge` and `Issue` types.

#### address-codec

- Address validation with `IsValidAddress`, `IsValidClassicAddress` and `IsValidXAddress`.
- Address conversion with `XAddressToClassicAddress` and `ClassicAddressToXAddress`.
- X-Address encoding/decoding with `EncodeXAddress` and `DecodeXAddress`.

#### keypairs

- New `DeriveNodeAddress` function.

#### xrpl

- New `AccountRoot`, `Amendments`, `Bridge`, `DID`, `DirectoryNode`, `Oracle`, `RippleState`, `XChainOwnedClaimID`, `XChainOwnedCreateAccountClaimID` ledger entry types.
- New `Multisign` utility function.
- New `NftHistory`, `NftsByIssuer`, `LedgerData`, `Check`, `BookOffers`, `PathFind`, `FeatureOne`, `FeatureAll` queries.
- New `SubmitMultisigned` request.
- New `AMMBid`, `AMMCreate`, `AMMDelete`, `AMMDeposit`, `AMMVote`, `AMMWithdraw` amm transactions.
- New `CheckCancel`, `CheckCash`, `CheckCreate` check transactions.
- New `DepositPreauth` transaction.
- New `DIDSet` and `DIDDelete` transactions.
- New `EscrowCreate`, `EscrowFinish`, `EscrowCancel` escrow transactions.
- New `OracleSet` and `OracleDelete` oracle transactions.
- New `XChainAccountCreateCommitment`, `XChainAddAccountCreateAttestation`, `XChainAddClaimAttestation`, `XChainClaim`, `XChainCommit`, `XChainCreateBridge`, `XChainCreateClaimID` and `XChainModifyBridge` cross-chain transactions.
- New `Multisign` wallet method.
- Ripple time conversion utility functions.
- Added query methods for websocket and rpc clients.
- New `SubmitMultisigned`, `AutofillMultisigned` and `SubmitAndWait` methods for both clients.
- Added `Autofill` method for rpc client.
- New `MaxRetries` and `RetryDelay` config options for both clients.

#### Other

- Implemented `secp256k1` algorithm.

### Changed

#### binary-codec

- Exported `FieldInstance` type.
- Updated `NewBinaryParser` constructor to accept `definitions.Definitions` as a parameter.
- Updated `NewSerializer` to `NewBinarySerializer` constructor.
- Refactored `FieldIDCodec` to be a struct with `Encode` and `Decode` methods.
- `FromJson` methods to `FromJSON`.
- `ToJson` methods to `ToJSON`.

#### address-codec

No changes were made.

#### keypairs

- Decoupled `ed25519` and `secp256k1` algorithms from `keypairs` package.
- Decoupled `der` parsing from `keypairs` package.

#### xrpl

- Renamed `CurrencyStringToHex` to `ConvertStringToHex` and `CurrencyHexToString` to `ConvertHexToString`.
- Renamed `HashSignedTx` to `TxBlob`.
- Wallet API methods have been renamed for better usability.
- Renamed `SendRequest` to `Request` methods for websocket and rpc clients.

### Fixed

#### xrpl

- Some queries did not have proper fields. All queries have been updated with the fields that are required by the XRP Ledger.
- Some transaction types did not have proper fields. All transaction types have been updated with the fields that are required by the XRP Ledger.
