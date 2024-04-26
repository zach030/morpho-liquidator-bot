package pkg

import (
	"fmt"
	"math/big"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type OneInch struct {
	swapCli *resty.Client
	ApiKey  string
}

func NewOneInch(key string, chainID *big.Int) *OneInch {
	return &OneInch{
		swapCli: resty.New().SetRetryCount(3).SetBaseURL(fmt.Sprintf("https://api.1inch.dev/swap/v6.0/%s", chainID.String())),
		ApiKey:  key,
	}
}

type QuoteResp struct {
	DstAmount string `json:"dstAmount"`
	Gas       int    `json:"gas"`
}

type SwapCalldataResp struct {
	DstAmount string `json:"dstAmount"`
	Tx        struct {
		From     string `json:"from"`
		To       string `json:"to"`
		Data     string `json:"data"`
		Value    string `json:"value"`
		Gas      int    `json:"gas"`
		GasPrice string `json:"gasPrice"`
	} `json:"tx"`
}

func (i *OneInch) QuoteToken(tokenIn, tokenOut string, inAmt *big.Int) (*QuoteResp, error) {
	var resp QuoteResp
	_, err := i.swapCli.R().
		SetResult(&resp).
		SetHeader("Authorization", i.ApiKey).
		SetQueryParams(map[string]string{
			"src":        tokenIn,
			"dst":        tokenOut,
			"amount":     inAmt.String(),
			"includeGas": "true",
		}).Get("/quote")
	if err != nil {
		log.Errorf("get token quote error=%v", err)
		return nil, err
	}
	return &resp, nil
}

func (i *OneInch) SwapCalldata(tokenIn, tokenOut, from string, inAmt *big.Int) (*SwapCalldataResp, error) {
	var resp SwapCalldataResp
	_, err := i.swapCli.R().
		SetResult(&resp).
		SetHeader("Authorization", i.ApiKey).
		SetQueryParams(map[string]string{
			"src":             tokenIn,
			"dst":             tokenOut,
			"amount":          inAmt.String(),
			"from":            from,
			"slippage":        "10",
			"includeGas":      "true",
			"disableEstimate": "true",
		}).Get("/swap")
	if err != nil {
		log.Errorf("get token quote error=%v", err)
		return nil, err
	}
	return &resp, nil
}
