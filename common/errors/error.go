// Copyright (c) 2017-2022 The Elastos Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import "github.com/coinbase/rosetta-sdk-go/types"

const (
	DefaultError int32 = iota
	TransactionNotExistError
	ReferTransactionNotExistError
	GetCurrentBlockError
	GetCurrentHeightError
	BlockNotExistError
	GetNeighborsError
	GetNodeStateError
	GetMempoolError
	DecodeAddressError
	EncodeToAddressError
	UnsupportNetworkError
	InvalidTransactionError
	GetAddressBalanceError
	GetUnspentUtxoError
	CurveTypeError
	PublicKeyError
	DecodeTransactionError
	DeserializeTransactionError
	DecodePublicKeyError
	SignatureTypeError
	PublishTransactionError
	InvalidCoinChangeError
	InvalidCoinChangeActionError
	InvalidOperationIdentifierError
	InvalidCoinIdentifierError
	InvalidOperationAmountError
	InvalidOperationAccountAddressError
	InvalidAmountError
	InvalidAccountIdentifierError
	NetworkIdentifierError
	InvalidTransactionHashError
	InvalidIncludeMempoolError
)

var (
	TransactionNotExist = &types.Error{
		Code:      TransactionNotExistError,
		Message:   "failed to get transaction by rpc",
		Retriable: true,
	}

	ReferTransactionNotExist = &types.Error{
		Code:      ReferTransactionNotExistError,
		Message:   "failed to get transaction reference by rpc",
		Retriable: true,
	}

	GetCurrentBlockFailed = &types.Error{
		Code:      GetCurrentBlockError,
		Message:   "failed to get current block height",
		Retriable: true,
	}

	GetCurrentHeightFailed = &types.Error{
		Code:      GetCurrentHeightError,
		Message:   "failed to get current height",
		Retriable: false,
	}

	BlockNotExist = &types.Error{
		Code:      BlockNotExistError,
		Message:   "failed to get block by rpc",
		Retriable: true,
	}

	GetNeighborsFailed = &types.Error{
		Code:      GetNeighborsError,
		Message:   "failed to get neighbors",
		Retriable: true,
	}

	GetNodeStateFailed = &types.Error{
		Code:      GetNodeStateError,
		Message:   "failed to get node state",
		Retriable: true,
	}

	GetMempoolFailed = &types.Error{
		Code:      GetMempoolError,
		Message:   "failed to get mempool",
		Retriable: false,
	}

	DecodeAddressFailed = &types.Error{
		Code:      DecodeAddressError,
		Message:   "failed to decode address",
		Retriable: false,
	}

	EncodeToAddressFailed = &types.Error{
		Code:      EncodeToAddressError,
		Message:   "failed to encode to address",
		Retriable: false,
	}

	UnsupportNetwork = &types.Error{
		Code:      UnsupportNetworkError,
		Message:   "unsupport network",
		Retriable: true,
	}

	NoNetworkIdentifier = &types.Error{
		Code:      NetworkIdentifierError,
		Message:   "must have NetworkIdentifier",
		Retriable: false,
	}

	InvalidTransaction = &types.Error{
		Code:      InvalidTransactionError,
		Message:   "invalid transaction data",
		Retriable: false,
	}

	GetAddressBalanceFailed = &types.Error{
		Code:      GetAddressBalanceError,
		Message:   "failed to get address balance",
		Retriable: false,
	}

	GetUnspentUtxoFailed = &types.Error{
		Code:      GetUnspentUtxoError,
		Message:   "failed to get address utxo",
		Retriable: false,
	}

	InvalidCurveType = &types.Error{
		Code:      CurveTypeError,
		Message:   "invalid curve type, need to be secp256r1",
		Retriable: false,
	}

	InvalidPublicKey = &types.Error{
		Code:      PublicKeyError,
		Message:   "invalid public key, decode failed",
		Retriable: false,
	}

	DecodeTransactionFailed = &types.Error{
		Code:      DecodeTransactionError,
		Message:   "failed to decode tx from hexstring",
		Retriable: false,
	}

	DeserializeTransactionFailed = &types.Error{
		Code:      DeserializeTransactionError,
		Message:   "failed to deserialize tx",
		Retriable: false,
	}

	DecodePublicKeyFailed = &types.Error{
		Code:      DecodePublicKeyError,
		Message:   "failed to decode pubkey from hexstring",
		Retriable: false,
	}

	InvalidSignatureType = &types.Error{
		Code:      SignatureTypeError,
		Message:   "invalid signature type, need to be ecdsa",
		Retriable: false,
	}

	PublishTransactionFailed = &types.Error{
		Code:      PublishTransactionError,
		Message:   "publish tx failed",
		Retriable: false,
	}

	InvalidCoinChange = &types.Error{
		Code:      InvalidCoinChangeError,
		Message:   "operation must have coin change",
		Retriable: false,
	}
	InvalidCoinChangeAction = &types.Error{
		Code:      InvalidCoinChangeActionError,
		Message:   "invalid coin change action",
		Retriable: false,
	}
	InvalidOperationIdentifier = &types.Error{
		Code:      InvalidOperationIdentifierError,
		Message:   "operation has no operation identifier",
		Retriable: false,
	}
	InvalidCoinIdentifier = &types.Error{
		Code:      InvalidCoinIdentifierError,
		Message:   "invalid coin identifier",
		Retriable: false,
	}
	InvalidOperationAmount = &types.Error{
		Code:      InvalidOperationAmountError,
		Message:   "invalid operation amount",
		Retriable: false,
	}
	InvalidOperationAccountAddress = &types.Error{
		Code:      InvalidOperationAccountAddressError,
		Message:   "invalid operation address",
		Retriable: false,
	}
	InvalidAmount = &types.Error{
		Code:      InvalidAmountError,
		Message:   "invalid amount",
		Retriable: false,
	}
	InvalidAccountIdentifier = &types.Error{
		Code:      InvalidAccountIdentifierError,
		Message:   "invalid account identifier",
		Retriable: false,
	}
	InvalidTransactionHash = &types.Error{
		Code:      InvalidTransactionHashError,
		Message:   "invalid transaction hash",
		Retriable: false,
	}

	InvalidIncludeMempool = &types.Error{
		Code:      InvalidIncludeMempoolError,
		Message:   "account/coins should set include_mempool to false",
		Retriable: false,
	}
)

var APIErrorMap = []*types.Error{
	TransactionNotExist,
	ReferTransactionNotExist,
	GetCurrentBlockFailed,
	GetCurrentHeightFailed,
	BlockNotExist,
	GetNeighborsFailed,
	GetNodeStateFailed,
	GetMempoolFailed,
	DecodeAddressFailed,
	EncodeToAddressFailed,
	UnsupportNetwork,
	InvalidTransaction,
	GetAddressBalanceFailed,
	GetUnspentUtxoFailed,
	InvalidCurveType,
	InvalidPublicKey,
	DecodeTransactionFailed,
	DecodePublicKeyFailed,
	InvalidSignatureType,
	PublishTransactionFailed,
	InvalidCoinChange,
	InvalidCoinChangeAction,
	InvalidOperationIdentifier,
	InvalidCoinIdentifier,
	InvalidOperationAmount,
	InvalidOperationAccountAddress,
	InvalidAmount,
	InvalidAccountIdentifier,
	InvalidTransactionHash,
	InvalidIncludeMempool,
}
