package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonSwapInfo struct {
	Addresses struct {
		Landlord     string `json:"landlord"`
		LandlordBase string `json:"landlordBase"`
		Rewarder     string `json:"rewarder"`
		MintWrapper  string `json:"mintWrapper"`
		IouMint      string `json:"iouMint"`
		Redeemer     string `json:"redeemer"`
		Sbr          string `json:"sbr"`
	} `json:"addresses"`
	Pools []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Tokens []struct {
			Name       string   `json:"name"`
			Address    string   `json:"address"`
			Decimals   int      `json:"decimals"`
			ChainID    int      `json:"chainId"`
			Symbol     string   `json:"symbol"`
			LogoURI    string   `json:"logoURI"`
			Tags       []string `json:"tags"`
			Extensions struct {
				Currency string `json:"currency"`
			} `json:"extensions,omitempty"`
			Extensions2 struct {
				Currency         string   `json:"currency"`
				Website          string   `json:"website"`
				AssetContract    string   `json:"assetContract"`
				UnderlyingTokens []string `json:"underlyingTokens"`
			} `json:"extensions,omitempty"`
		} `json:"tokens"`
		TokenIcons []struct {
			Name       string   `json:"name"`
			Address    string   `json:"address"`
			Decimals   int      `json:"decimals"`
			ChainID    int      `json:"chainId"`
			Symbol     string   `json:"symbol"`
			LogoURI    string   `json:"logoURI"`
			Tags       []string `json:"tags"`
			Extensions struct {
				Currency string `json:"currency"`
			} `json:"extensions,omitempty"`
			Extensions2 struct {
				Currency         string   `json:"currency"`
				Website          string   `json:"website"`
				AssetContract    string   `json:"assetContract"`
				UnderlyingTokens []string `json:"underlyingTokens"`
			} `json:"extensions,omitempty"`
		} `json:"tokenIcons"`
		UnderlyingIcons []struct {
			Name       string   `json:"name"`
			Address    string   `json:"address"`
			Decimals   int      `json:"decimals"`
			ChainID    int      `json:"chainId"`
			Symbol     string   `json:"symbol"`
			LogoURI    string   `json:"logoURI"`
			Tags       []string `json:"tags"`
			Extensions struct {
				Currency string `json:"currency"`
			} `json:"extensions"`
		} `json:"underlyingIcons"`
		Currency string `json:"currency"`
		LpToken  struct {
			Symbol     string   `json:"symbol"`
			Name       string   `json:"name"`
			LogoURI    string   `json:"logoURI"`
			Decimals   int      `json:"decimals"`
			Address    string   `json:"address"`
			ChainID    int      `json:"chainId"`
			Tags       []string `json:"tags"`
			Extensions struct {
				Website          string   `json:"website"`
				UnderlyingTokens []string `json:"underlyingTokens"`
				Source           string   `json:"source"`
			} `json:"extensions"`
		} `json:"lpToken"`
		PlotKey string `json:"plotKey"`
		Swap    struct {
			Config struct {
				SwapAccount    string `json:"swapAccount"`
				SwapProgramID  string `json:"swapProgramID"`
				TokenProgramID string `json:"tokenProgramID"`
				Authority      string `json:"authority"`
			} `json:"config"`
			State struct {
				IsInitialized       bool   `json:"isInitialized"`
				IsPaused            bool   `json:"isPaused"`
				Nonce               int    `json:"nonce"`
				FutureAdminDeadline int    `json:"futureAdminDeadline"`
				FutureAdminAccount  string `json:"futureAdminAccount"`
				AdminAccount        string `json:"adminAccount"`
				TokenA              struct {
					AdminFeeAccount string `json:"adminFeeAccount"`
					Reserve         string `json:"reserve"`
					Mint            string `json:"mint"`
				} `json:"tokenA"`
				TokenB struct {
					AdminFeeAccount string `json:"adminFeeAccount"`
					Reserve         string `json:"reserve"`
					Mint            string `json:"mint"`
				} `json:"tokenB"`
				PoolTokenMint      string `json:"poolTokenMint"`
				InitialAmpFactor   string `json:"initialAmpFactor"`
				TargetAmpFactor    string `json:"targetAmpFactor"`
				StartRampTimestamp int    `json:"startRampTimestamp"`
				StopRampTimestamp  int    `json:"stopRampTimestamp"`
				Fees               struct {
					AdminTrade struct {
						Formatted   string `json:"formatted"`
						Numerator   string `json:"numerator"`
						Denominator string `json:"denominator"`
					} `json:"adminTrade"`
					AdminWithdraw struct {
						Formatted   string `json:"formatted"`
						Numerator   string `json:"numerator"`
						Denominator string `json:"denominator"`
					} `json:"adminWithdraw"`
					Trade struct {
						Formatted   string `json:"formatted"`
						Numerator   string `json:"numerator"`
						Denominator string `json:"denominator"`
					} `json:"trade"`
					Withdraw struct {
						Formatted   string `json:"formatted"`
						Numerator   string `json:"numerator"`
						Denominator string `json:"denominator"`
					} `json:"withdraw"`
				} `json:"fees"`
			} `json:"state"`
		} `json:"swap"`
		Quarry string `json:"quarry"`
	} `json:"pools"`
}

func NewJsonSwapInfo(url string) (*JsonSwapInfo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jsonSwapInfo JsonSwapInfo
	if err := json.Unmarshal(body, &jsonSwapInfo); err != nil {
		return nil, err
	}

	return &jsonSwapInfo, nil
}

func (j *JsonSwapInfo) ListPools() {
	for _, pool := range j.Pools {
		log.Print("\n----------\n")

		log.Printf("Pool id: %s\n", pool.ID)

		log.Print("Tokens:\n")
		for _, token := range pool.Tokens {
			log.Printf("\tToken %s (symbol %s) address: %s\n", token.Name, token.Symbol, token.Address)
		}

		log.Printf("LP token %s (symbol %s) address: %s\n", pool.LpToken.Name, pool.LpToken.Symbol, pool.LpToken.Address)

		log.Printf("Swap account: %s", pool.Swap.Config.SwapAccount)

		log.Print("\n----------\n")
	}
}
