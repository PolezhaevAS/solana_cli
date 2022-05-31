package instructions

import (
	"errors"
	"log"

	"github.com/gagliardetto/solana-go"
)

/// Swap the tokens in the pool.
///
/// 0. `[]`StableSwap
/// 1. `[]` $authority
/// 2. `[signer]` User authority.
/// 3. `[writable]` token_(A|B) SOURCE Account, amount is transferable by $authority,
/// 4. `[writable]` token_(A|B) Base Account to swap INTO.  Must be the SOURCE token.
/// 5. `[writable]` token_(A|B) Base Account to swap FROM.  Must be the DESTINATION token.
/// 6. `[writable]` token_(A|B) DESTINATION Account assigned to USER as the owner.
/// 7. `[writable]` token_(A|B) admin fee Account. Must have same mint as DESTINATION token.
/// 8. `[]` Token program id

type Swap struct {
	prog     solana.PublicKey
	accounts []*solana.AccountMeta
	data     []byte
}

func NewSwap(prog solana.PublicKey) *Swap {
	return &Swap{prog: prog, accounts: make([]*solana.AccountMeta, 8)}
}

func (i *Swap) Build() (*solana.GenericInstruction, error) {
	if len(i.data) == 0 {
		return nil, errors.New("add data bytes to instruction")
	}

	i.accounts = append(i.accounts, solana.NewAccountMeta(solana.TokenProgramID, false, false))

	return solana.NewInstruction(i.prog, i.accounts, i.data), nil
}

func (i *Swap) SetData(data []byte) *Swap {
	i.data = data
	return i
}

func (i *Swap) SetSwapAccount(key solana.PublicKey) *Swap {
	i.accounts[0] = solana.NewAccountMeta(key, false, false)
	return i
}

func (i *Swap) SetAuthority(key solana.PublicKey) *Swap {
	i.accounts[1] = solana.NewAccountMeta(key, false, false)
	return i
}

func (i *Swap) SetUserAuthority(key solana.PublicKey) *Swap {
	i.accounts[2] = solana.NewAccountMeta(key, true, true)
	return i
}

func (i *Swap) SetUserSource(key solana.PublicKey) *Swap {
	i.accounts[3] = solana.NewAccountMeta(key, true, false)
	return i
}

func (i *Swap) SetPoolSource(key solana.PublicKey) *Swap {
	i.accounts[4] = solana.NewAccountMeta(key, true, false)
	return i
}

func (i *Swap) SetPoolDestination(key solana.PublicKey) *Swap {
	i.accounts[5] = solana.NewAccountMeta(key, true, false)
	return i
}

func (i *Swap) SetUserDestination(key solana.PublicKey) *Swap {
	i.accounts[6] = solana.NewAccountMeta(key, true, false)
	return i
}

func (i *Swap) SetAdminDestination(key solana.PublicKey) *Swap {
	i.accounts[7] = solana.NewAccountMeta(key, true, false)
	return i
}

func (i *Swap) ShowAccounts() {
	log.Println("Swap account:\t", i.accounts[0].PublicKey.String())
	log.Println("Authority:\t", i.accounts[1].PublicKey.String())
	log.Println("User Authority:\t", i.accounts[2].PublicKey.String())
	log.Println("User Source:\t", i.accounts[3].PublicKey.String())
	log.Println("Pool Source:\t", i.accounts[4].PublicKey.String())
	log.Println("Pool Destination:\t", i.accounts[5].PublicKey.String())
	log.Println("User Destination:\t", i.accounts[6].PublicKey.String())
	log.Println("Admin Destination:\t", i.accounts[7].PublicKey.String())
}
