package instructions

import (
	"bytes"

	ag_binary "github.com/gagliardetto/binary"
)

// Swap instruction data
type SwapData struct {
	Prog      uint8
	AmountIn  uint64
	AmountOut uint64
}

// Get new swap data
func NewSwapData(in, out uint64) *SwapData {
	return &SwapData{
		Prog:      1,
		AmountIn:  in,
		AmountOut: out,
	}
}

// Get swap data bytes
func (s *SwapData) GetBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := ag_binary.NewBinEncoder(buf).Encode(s)
	return buf.Bytes(), err
}

// Withdraw instruction data
type WithdrawData struct {
	Prog                uint8
	PoolTokenAmount     uint64
	MinimumTokenAmountA uint64
	MinumumTokenAmountB uint64
}

// Get new withdraw data
func NewWithdrawData(pool, amountA, amountB uint64) *WithdrawData {
	return &WithdrawData{
		Prog:                3,
		PoolTokenAmount:     pool,
		MinimumTokenAmountA: amountA,
		MinumumTokenAmountB: amountB,
	}
}

// Get withdraw data bytes
func (w *WithdrawData) GetBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := ag_binary.NewBinEncoder(buf).Encode(w)
	return buf.Bytes(), err
}

// Withdraw One instruction data
type WithdrawOneData struct {
	Prog               uint8
	PoolTokenAmount    uint64
	MinimumTokenAmount uint64
}

// Get new withdraw one data
func NewWithdrawOneData(pool, amount uint64) *WithdrawOneData {
	return &WithdrawOneData{
		Prog:               4,
		PoolTokenAmount:    pool,
		MinimumTokenAmount: amount,
	}
}

// Get withdraw one data bytes
func (w *WithdrawOneData) GetBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := ag_binary.NewBinEncoder(buf).Encode(w)
	return buf.Bytes(), err
}
