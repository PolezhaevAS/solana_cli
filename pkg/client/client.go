package client

import (
	"context"
	"solana/pkg/instructions"
	"solana/pkg/model"

	"github.com/gagliardetto/solana-go"
	a "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

type Client struct {
	rpc *rpc.Client
	ws  *ws.Client
}

func NewClient(ctx context.Context, cluster rpc.Cluster) (*Client, error) {
	rpc := rpc.New(cluster.RPC)
	ws, err := ws.Connect(ctx, cluster.WS)
	if err != nil {
		return nil, err
	}

	return &Client{
		rpc: rpc,
		ws:  ws,
	}, nil
}

func (c *Client) Close() {
	c.ws.Close()
}

func (c *Client) Recent(ctx context.Context) (solana.Hash, error) {
	r, err := c.rpc.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return solana.Hash{}, err
	}

	return r.Value.Blockhash, nil
}

func (c *Client) SendInstructions(ctx context.Context, instr []solana.Instruction, wallet *solana.Wallet) (solana.Signature, error) {

	recent, err := c.Recent(ctx)
	if err != nil {
		return solana.Signature{}, err
	}

	tx, err := solana.NewTransaction(
		instr, recent, solana.TransactionPayer(wallet.PublicKey()),
	)
	if err != nil {
		return solana.Signature{}, err
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if wallet.PublicKey().Equals(key) {
				return &wallet.PrivateKey
			}
			return nil
		},
	)
	if err != nil {
		return solana.Signature{}, err
	}

	sig, err := confirm.SendAndConfirmTransaction(
		ctx, c.rpc, c.ws, tx,
	)
	if err != nil {
		return solana.Signature{}, err
	}

	return sig, nil
}

func (c *Client) GetTokenAccount(ctx context.Context, pubKey solana.PublicKey, mint solana.PublicKey) (solana.PublicKey, *a.Instruction, error) {
	// get account
	account, _, err := solana.FindAssociatedTokenAddress(pubKey, mint)
	if err != nil {
		return solana.PublicKeyFromBytes([]byte{}), nil, err
	}

	// find account
	_, err = c.rpc.GetAccountInfo(ctx, account)
	if err != nil && err.Error() == "not found" {
		instr, err := a.NewCreateInstruction(pubKey, pubKey, mint).ValidateAndBuild()
		if err != nil {
			return account, nil, err
		}
		return account, instr, err
	} else if err != nil {
		return account, nil, err
	} else {
		return account, nil, err
	}
}

func (c *Client) Airdrop(ctx context.Context, pubKey solana.PublicKey, sol uint64) (solana.Signature, error) {
	sig, err := c.rpc.RequestAirdrop(ctx, pubKey, sol*solana.LAMPORTS_PER_SOL, rpc.CommitmentFinalized)
	if err != nil {
		return solana.Signature{}, err
	}

	return sig, nil
}

func (c *Client) Balance(ctx context.Context, pubKey solana.PublicKey) (uint64, error) {
	out, err := c.rpc.GetBalance(ctx, pubKey, rpc.CommitmentFinalized)
	if err != nil {
		return 0, err
	}

	return out.Value, nil
}

func (c *Client) SwapInfo(ctx context.Context, account solana.PublicKey) (*model.SwapInfo, error) {
	var swapInfo model.SwapInfo
	err := c.rpc.GetAccountDataInto(ctx, account, &swapInfo)
	if err != nil {
		return nil, err
	}
	return &swapInfo, nil
}

func (c *Client) Swap(ctx context.Context,
	programId, swapAccount, tokenA, tokenB solana.PublicKey,
	wallet *solana.Wallet,
	swapData *instructions.SwapData,
	showAccounts bool) (solana.Signature, error) {

	instrs := []solana.Instruction{}

	swapInfo, err := c.SwapInfo(ctx, swapAccount)
	if err != nil {
		return solana.Signature{}, err
	}

	swapTokenA, err := swapInfo.HasToken(tokenA)
	if err != nil {
		return solana.Signature{}, err
	}

	swapTokenB, err := swapInfo.HasToken(tokenB)
	if err != nil {
		return solana.Signature{}, err
	}

	swapAuthority, err := solana.CreateProgramAddress([][]byte{swapAccount.Bytes(), {swapInfo.Nonce}}, programId)
	if err != nil {
		return solana.Signature{}, err
	}

	userTokenA, instrTokenA, err := c.GetTokenAccount(ctx, wallet.PublicKey(), swapTokenA.TokenMint)
	if err != nil {
		return solana.Signature{}, err
	}

	if instrTokenA != nil {
		instrs = append(instrs, instrTokenA)
	}

	userTokenB, instrTokenB, err := c.GetTokenAccount(ctx, wallet.PublicKey(), swapTokenB.TokenMint)
	if err != nil {
		return solana.Signature{}, err
	}

	if instrTokenB != nil {
		instrs = append(instrs, instrTokenA)
	}

	bytes, err := swapData.GetBytes()
	if err != nil {
		return solana.Signature{}, err
	}

	swap := instructions.NewSwap(programId).
		SetSwapAccount(swapAccount).
		SetAuthority(swapAuthority).
		SetUserAuthority(wallet.PublicKey()).
		SetUserSource(userTokenA).
		SetPoolSource(swapTokenA.TokenReserve).
		SetPoolDestination(swapTokenB.TokenReserve).
		SetUserDestination(userTokenB).
		SetAdminDestination(swapTokenB.TokenFee).
		SetData(bytes)

	if showAccounts {
		swap.ShowAccounts()
		return solana.Signature{}, nil
	}

	swapInstr, err := swap.Build()
	if err != nil {
		return solana.Signature{}, err
	}

	instrs = append(instrs, swapInstr)
	sig, err := c.SendInstructions(ctx, instrs, wallet)
	if err != nil {
		return solana.Signature{}, err
	}

	return sig, nil

}
