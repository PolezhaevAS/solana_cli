package model

import (
	"errors"

	"github.com/gagliardetto/solana-go"
)

// Swap info struct
type SwapInfo struct {
	IsInitialized       bool
	IsPaused            bool
	Nonce               uint8
	InitialAmpFactor    uint64
	TargetAmpFactor     uint64
	StartRampTs         int64
	StopRampTs          int64
	FutureAdminDeadline int64
	FutureAdminKey      solana.PublicKey
	AdminKey            solana.PublicKey
	TokenAReserve       solana.PublicKey
	TokenBReserve       solana.PublicKey
	PoolTokenMint       solana.PublicKey
	TokenAMint          solana.PublicKey
	TokenBMint          solana.PublicKey
	TokenAFee           solana.PublicKey
	TokenBFee           solana.PublicKey
	Fees                *Fees
}

// Fees struct
type Fees struct {
	AdminTradeFeeNumerator      uint64
	AdminTradeFeeDenominator    uint64
	AdminWithdrawFeeNumerator   uint64
	AdminWithdrawDeeDenominator uint64
	TradeFeeNumerator           uint64
	TradeFeeDenominator         uint64
	WithdrawFeeNumerator        uint64
	WithdrawFeeDenominator      uint64
}

// Token info in swap
type SwapTokenInfo struct {
	TokenMint    solana.PublicKey
	TokenReserve solana.PublicKey
	TokenFee     solana.PublicKey
}

// Check token on swap info and return public key or error
func (s *SwapInfo) HasToken(token solana.PublicKey) (*SwapTokenInfo, error) {
	if token == s.TokenAMint {
		return &SwapTokenInfo{
			TokenMint:    s.TokenAMint,
			TokenReserve: s.TokenAReserve,
			TokenFee:     s.TokenAFee,
		}, nil
	}

	if token == s.TokenBMint {
		return &SwapTokenInfo{
			TokenMint:    s.TokenBMint,
			TokenReserve: s.TokenBReserve,
			TokenFee:     s.TokenBFee,
		}, nil
	}

	return nil, errors.New("cann't find token in swap info")
}
