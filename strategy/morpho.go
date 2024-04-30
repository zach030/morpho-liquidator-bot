package strategy

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/zach030/morpho-liquidator-bot/config"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/zach030/morpho-liquidator-bot/abi/bot"
	"github.com/zach030/morpho-liquidator-bot/pkg"

	log "github.com/sirupsen/logrus"
	"github.com/zach030/morpho-liquidator-bot/abi/irm"
	"github.com/zach030/morpho-liquidator-bot/abi/morpho"
	"github.com/zach030/morpho-liquidator-bot/abi/oracle"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
)

var (
	WETH                    = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	supplyTopic             = EventTopic("Supply(bytes32,address,address,uint256,uint256)")
	withdrawTopic           = EventTopic("Withdraw(bytes32,address,address,address,uint256,uint256)")
	borrowTopic             = EventTopic("Borrow(bytes32,address,address,address,uint256,uint256)")
	repayTopic              = EventTopic("Repay(bytes32,address,address,uint256,uint256)")
	supplyCollateralTopic   = EventTopic("SupplyCollateral(bytes32,address,address,uint256)")
	withdrawCollateralTopic = EventTopic("WithdrawCollateral(bytes32,address,address,address,uint256)")
	liquidateTopic          = EventTopic("Liquidate(bytes32,address,address,uint256,uint256,uint256,uint256,uint256)")
)

type Morpho struct {
	botAddress    common.Address
	httpClient    *ethclient.Client
	blueClient    *resty.Client
	oneinchClient *pkg.OneInch
	bot           *bot.Bot
	privateKey    *ecdsa.PrivateKey
	activeMarkets map[string]*MarketInfo
	pendingTx     map[*types.Transaction]struct{}
}

func NewMorpho(cfg *config.Config) *Morpho {
	httpClient, err := ethclient.Dial(cfg.HttpEndpoint)
	if err != nil {
		panic(err)
	}
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		panic(err)
	}
	blueCli := resty.New().SetBaseURL("https://blue-api.morpho.org/graphql").SetRetryCount(3)
	oneinchCli := pkg.NewOneInch(cfg.OneInchApiKey, big.NewInt(1))
	b, err := bot.NewBot(common.HexToAddress(cfg.BotAddress), httpClient)
	m := &Morpho{
		botAddress:    common.HexToAddress(cfg.BotAddress),
		httpClient:    httpClient,
		blueClient:    blueCli,
		bot:           b,
		privateKey:    privateKey,
		oneinchClient: oneinchCli,
		activeMarkets: make(map[string]*MarketInfo),
		pendingTx:     make(map[*types.Transaction]struct{}),
	}
	m.loadMarkets(cfg.Markets)
	go m.checkResult()
	return m
}

func (m *Morpho) checkResult() {
	ticker := time.NewTicker(5 * time.Second)
	for t := range ticker.C {
		log.Infof("liquidate check receipt ticker: %s", t.Format(time.DateTime))
		for tx := range m.pendingTx {
			receipt, err := m.httpClient.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Errorf("get tx receipt error: %s", err)
				continue
			}
			if receipt.Status == types.ReceiptStatusSuccessful {
				log.Warnf("liquidate tx %s success!", tx.Hash().Hex())
			} else {
				log.Errorf("liquidate tx %s failed!", tx.Hash().Hex())
			}
			delete(m.pendingTx, tx)
		}
	}
}

func (m *Morpho) Start(header <-chan *types.Header, tx <-chan *types.Transaction, log <-chan *types.Log) {
	for {
		select {
		case h := <-header:
			m.processNewBlock(h)
		case l := <-log:
			m.processNewLog(l)
		case <-tx:

		}
	}
}

func (m *Morpho) processNewBlock(header *types.Header) {
	var wg sync.WaitGroup
	for _, market := range m.activeMarkets {
		wg.Add(1)
		go func(market *MarketInfo) {
			defer wg.Done()
			m.updateMarket(market)
			m.checkPosition(market)
		}(market)
	}
	wg.Wait()
	log.Infof("Morpho success processNewBlock height=%v", header.Number)
}

func (m *Morpho) processNewLog(vLog *types.Log) {
	if vLog.Address != pkg.MorphoAddress {
		log.Debugf("morpho: unknown contract: %s", vLog.Address.Hex())
		return
	}
	log.Warnf("[Morpho] morpho handle log: %s", vLog.Topics[0].Hex())
	switch vLog.Topics[0] {
	case supplyTopic, withdrawTopic, borrowTopic, repayTopic, liquidateTopic, withdrawCollateralTopic: // 更新market
		m.handleEventUpdateMarket(vLog)
	default:
		return
	}
}

func (m *Morpho) handleEventUpdateMarket(vLog *types.Log) {
	var marketID = vLog.Topics[1]
	market, ok := m.activeMarkets[marketID.String()]
	if !ok {
		log.Warnf("handleEventUpdateMarket new market id=%s not exist", marketID.String())
		return
	}
	newMarket, err := m.getMarketInfo(marketID.String())
	if err != nil {
		log.Errorf("getMarketInfo marketID=%s err=%v", marketID.String(), err)
		return
	}
	newMarketInfo, err := m.buildMarket(newMarket)
	if newMarketInfo.MarketState.LastUpdate.Uint64() < vLog.BlockNumber {
		log.Warnf("handleEventUpdateMarket ignore old event id=%s", marketID.String())
		return
	}
	m.activeMarkets[marketID.String()] = newMarketInfo
	log.Infof("handleEventUpdateMarket update market id=%s name=%s", marketID.String(), market.Name)
}

func (m *Morpho) updateMarket(market *MarketInfo) {
	price, err := market.Oracle.Price(nil)
	if err != nil {
		log.Errorf("updateMarket get price error=%v", err)
		return
	}
	market.CollateralPrice = price
	market.BorrowRate = m.getMarketBorrowRate(market)
}

func (m *Morpho) checkPosition(market *MarketInfo) {
	positions, err := m.loadMarketPositions(market.ID)
	if err != nil {
		log.Errorf("checkPosition loadMarketPositions error=%v", err)
		return
	}
	for user, position := range positions {
		tx, err := m.startLiquidate(user, position, market)
		if err != nil {
			log.Errorf("startLiquidate error=%v", err)
			continue
		}
		m.pendingTx[tx] = struct{}{}
	}
}

func (m *Morpho) startLiquidate(borrower string, position *Position, market *MarketInfo) (*types.Transaction, error) {
	seizedAssets := GetSeizedAmtByShares(position.BorrowShares, market)
	repaid := toAssetsUp(position.BorrowShares, market.MarketState.TotalBorrowAssets, market.MarketState.TotalBorrowShares)
	swapResult, err := m.generateSwapData(market.CollateralToken.Address, market.LoanToken.Address, seizedAssets)
	if err != nil {
		return nil, err
	}
	loanTokenOutAmt, ok := new(big.Int).SetString(swapResult.DstAmt, 10)
	if !ok {
		return nil, errors.New("parse loan token out amount error")
	}
	if loanTokenOutAmt.Cmp(repaid) <= 0 {
		return nil, errors.New("loan token out amount less than repaid")
	}
	loanTokenRevenue := new(big.Int).Sub(loanTokenOutAmt, repaid)
	quoteResp, err := m.oneinchClient.QuoteToken(market.LoanToken.Address, WETH.Hex(), loanTokenRevenue)
	if err != nil {
		return nil, err
	}
	wethProfit, _ := new(big.Int).SetString(quoteResp.DstAmount, 10)
	opt, err := bind.NewKeyedTransactorWithChainID(m.privateKey, big.NewInt(1))
	if err != nil {
		return nil, err
	}
	opt.NoSend = true
	swapData, _ := hexutil.Decode(swapResult.Data)
	tx, err := m.bot.MorphoLiquidate(opt, common.HexToHash(market.ID), common.HexToAddress(borrower), seizedAssets, common.HexToAddress(swapResult.To), swapData)
	if err != nil {
		log.Errorf("startLiquidate error=%v", err)
		return nil, err
	}
	if tx.Cost().Cmp(wethProfit) >= 0 {
		return nil, errors.New("no profit")
	}
	return tx, nil
}

type SwapResult struct {
	To     string
	Data   string
	DstAmt string
}

func (m *Morpho) generateSwapData(collateral, loan string, in *big.Int) (*SwapResult, error) {
	resp, err := m.oneinchClient.SwapCalldata(collateral, loan, m.botAddress.Hex(), in)
	if err != nil {
		log.Errorf("generateSwapData failed, err=%v", err)
		return nil, err
	}
	return &SwapResult{
		To:     resp.Tx.To,
		Data:   resp.Tx.Data,
		DstAmt: resp.DstAmount,
	}, nil
}

type MorphoMarket struct {
	Id            string      `json:"id"`
	UniqueKey     string      `json:"uniqueKey"`
	Lltv          interface{} `json:"lltv"`
	OracleAddress string      `json:"oracleAddress"`
	IrmAddress    string      `json:"irmAddress"`
	LoanAsset     struct {
		Address  string `json:"address"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
	} `json:"loanAsset"`
	CollateralAsset struct {
		Address  string `json:"address"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
	} `json:"collateralAsset"`
	State struct {
		BorrowShares interface{} `json:"borrowShares"`
		BorrowAssets interface{} `json:"borrowAssets"`
		SupplyShares interface{} `json:"supplyShares"`
		SupplyAssets interface{} `json:"supplyAssets"`
		Fee          interface{} `json:"fee"`
		TimeStamp    int64       `json:"timestamp"`
	} `json:"state"`
}

type MarketItem struct {
	Items []MorphoMarket `json:"items"`
}

type MorphoMarketsResp struct {
	Data struct {
		Markets MarketItem `json:"markets"`
	} `json:"data"`
}

func (m *Morpho) loadMarkets(markets []string) {
	var res MorphoMarketsResp
	marketsArray := fmt.Sprintf("[\"%s\"]", strings.Join(markets, "\" \""))
	query := `{markets(where: {uniqueKey_in: %s}) {items {id uniqueKey lltv oracleAddress irmAddress loanAsset {address symbol decimals} collateralAsset {address symbol decimals} state {borrowShares borrowAssets supplyShares supplyAssets fee timestamp utilization}}}}`
	query = fmt.Sprintf(query, marketsArray)
	payload := map[string]string{
		"query": query,
	}
	payloadBytes, err := json.Marshal(payload)
	_, err = m.blueClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payloadBytes).
		SetResult(&res).
		Post("")
	if err != nil {
		log.Errorf("loadMarkets failed, err=%v", err)
		return
	}
	for _, market := range res.Data.Markets.Items {
		zeroAddress := common.Address{}.Hex()
		if market.UniqueKey == zeroAddress {
			continue
		}
		marketInfo, err := m.buildMarket(&market)
		if err != nil {
			log.Errorf("buildMarketInfo failed, err=%v", err)
			continue
		}
		m.activeMarkets[market.UniqueKey] = marketInfo
	}
	log.Infof("loadMarkets success, count=%d", len(m.activeMarkets))
}

type Token struct {
	Address  string
	Decimals string
	Symbol   string
}

type MarketInfo struct {
	ID              string
	Name            string
	LoanToken       Token
	CollateralToken Token
	Params          morpho.IMorphoMarketParams
	MarketState     morpho.IMorphoMarket
	CollateralPrice *big.Int
	BorrowRate      *big.Int
	Oracle          *oracle.Oracle
	Irm             *irm.Irm
}

func (m *Morpho) buildMarket(market *MorphoMarket) (*MarketInfo, error) {
	zeroAddress := common.Address{}.Hex()
	if market.CollateralAsset.Address == zeroAddress {
		return nil, errors.New(fmt.Sprintf("market id=%s collateral address is zero", market.UniqueKey))
	}
	if market.LoanAsset.Address == zeroAddress {
		return nil, errors.New(fmt.Sprintf("market id=%s loan address is zero", market.UniqueKey))
	}
	if market.OracleAddress == zeroAddress {
		return nil, errors.New(fmt.Sprintf("market id=%s oracle address is zero", market.UniqueKey))
	}
	if market.IrmAddress == zeroAddress {
		return nil, errors.New(fmt.Sprintf("market id=%s Irm address is zero", market.UniqueKey))
	}
	o, err := oracle.NewOracle(common.HexToAddress(market.OracleAddress), m.httpClient)
	if err != nil {
		return nil, err
	}
	lltv, _ := new(big.Int).SetString(fmt.Sprintf("%v", market.Lltv), 10)
	sa, ok := market.State.SupplyAssets.(string)
	if !ok {
		sa = strconv.FormatFloat(market.State.SupplyAssets.(float64), 'f', -1, 64)
	}
	tsa, _ := new(big.Int).SetString(fmt.Sprintf("%v", sa), 10)
	ss, ok := market.State.SupplyShares.(string)
	if !ok {
		ss = strconv.FormatFloat(market.State.SupplyShares.(float64), 'f', -1, 64)
	}
	tss, _ := new(big.Int).SetString(fmt.Sprintf("%v", ss), 10)
	ba, ok := market.State.BorrowAssets.(string)
	if !ok {
		ba = strconv.FormatFloat(market.State.BorrowAssets.(float64), 'f', -1, 64)
	}
	tba, _ := new(big.Int).SetString(fmt.Sprintf("%v", ba), 10)
	bs, ok := market.State.BorrowShares.(string)
	if !ok {
		bs = strconv.FormatFloat(market.State.BorrowShares.(float64), 'f', -1, 64)
	}
	tbs, _ := new(big.Int).SetString(fmt.Sprintf("%v", bs), 10)
	lu, _ := new(big.Int).SetString(fmt.Sprintf("%v", market.State.TimeStamp), 10)
	fee, _ := new(big.Int).SetString(fmt.Sprintf("%v", market.State.Fee), 10)
	price, err := o.Price(nil)
	if err != nil {
		log.Errorf("marketId=%s buildMarketInfo get price error=%v oracle=%s", market.UniqueKey, err, market.OracleAddress)
		return nil, err
	}
	irm, _ := irm.NewIrm(common.HexToAddress(market.IrmAddress), m.httpClient)
	mi := &MarketInfo{
		ID:   market.UniqueKey,
		Name: strings.Join([]string{market.CollateralAsset.Symbol, market.LoanAsset.Symbol}, "-"),
		LoanToken: Token{
			Address:  market.LoanAsset.Address,
			Decimals: strconv.Itoa(market.LoanAsset.Decimals),
			Symbol:   market.LoanAsset.Symbol,
		},
		CollateralToken: Token{
			Address:  market.CollateralAsset.Address,
			Decimals: strconv.Itoa(market.CollateralAsset.Decimals),
			Symbol:   market.CollateralAsset.Symbol,
		},
		Params: morpho.IMorphoMarketParams{
			LoanToken:       common.HexToAddress(market.LoanAsset.Address),
			CollateralToken: common.HexToAddress(market.CollateralAsset.Address),
			Oracle:          common.HexToAddress(market.OracleAddress),
			Irm:             common.HexToAddress(market.IrmAddress),
			Lltv:            lltv,
		},
		MarketState: morpho.IMorphoMarket{
			TotalSupplyAssets: tsa,
			TotalSupplyShares: tss,
			TotalBorrowAssets: tba,
			TotalBorrowShares: tbs,
			LastUpdate:        lu,
			Fee:               fee,
		},
		Irm:             irm,
		Oracle:          o,
		CollateralPrice: price,
	}
	mi.BorrowRate = m.getMarketBorrowRate(mi)
	log.Infof("buildMarketInfo success, market=%s", mi.Name)
	return mi, nil
}

func (m *Morpho) getMarketBorrowRate(market *MarketInfo) *big.Int {
	borrowRate, err := market.Irm.BorrowRateView(nil, irm.MarketParams{
		LoanToken:       market.Params.LoanToken,
		CollateralToken: market.Params.CollateralToken,
		Oracle:          market.Params.Oracle,
		Irm:             market.Params.Irm,
		Lltv:            market.Params.Lltv,
	}, irm.Market{
		TotalSupplyAssets: market.MarketState.TotalSupplyAssets,
		TotalSupplyShares: market.MarketState.TotalSupplyShares,
		TotalBorrowAssets: market.MarketState.TotalBorrowAssets,
		TotalBorrowShares: market.MarketState.TotalBorrowShares,
		LastUpdate:        market.MarketState.LastUpdate,
		Fee:               market.MarketState.Fee,
	})
	if err != nil {
		log.Errorf("get borrow rate error=%v", err)
		return nil
	}
	return borrowRate
}

type MarketPosition struct {
	Items []struct {
		Id              string      `json:"id"`
		BorrowShares    interface{} `json:"borrowShares"`
		BorrowAssets    interface{} `json:"borrowAssets"`
		BorrowAssetsUsd float64     `json:"borrowAssetsUsd"`
		Collateral      interface{} `json:"collateral"`
		CollateralUsd   float64     `json:"collateralUsd"`
		HealthFactor    float64     `json:"healthFactor"`
		User            struct {
			Address string `json:"address"`
		} `json:"user"`
	} `json:"items"`
}

type MorphoPositionsResp struct {
	Data struct {
		MarketPositions MarketPosition `json:"marketPositions"`
	} `json:"data"`
}

type Position struct {
	BorrowShares  *big.Int
	Collateral    *big.Int
	HealthyFactor float64
}

func (m *Morpho) loadMarketPositions(marketId string) (map[string]*Position, error) {
	var res MorphoPositionsResp
	data := map[string]interface{}{
		"operationName": "getMarketPositions",
		"variables": map[string]interface{}{
			"where": map[string]interface{}{
				"marketUniqueKey_in": []string{marketId},
				"chainId_in":         []int{1},
				"borrowShares_gte":   1,
				"collateral_gte":     1,
				"healthFactor_lte":   1,
			},
			"orderBy":        "BorrowShares",
			"orderDirection": "Desc",
			"first":          100,
			"skip":           0,
		},
		"query": `query getMarketPositions($where: MarketPositionFilters, $orderBy: MarketPositionOrderBy, $orderDirection: OrderDirection, $first: Int, $skip: Int) {marketPositions(first: $first skip: $skip where: $where orderBy: $orderBy orderDirection: $orderDirection) {items {id borrowShares borrowAssets borrowAssetsUsd collateral collateralUsd healthFactor user {address}}}}`,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	_, err = m.blueClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(bytes.NewBuffer(jsonData)).
		SetResult(&res).
		Post("")
	if err != nil {
		log.Errorf("get market info failed, err=%v", err)
		return nil, err
	}
	ret := make(map[string]*Position)
	for _, position := range res.Data.MarketPositions.Items {
		bs := fmt.Sprintf("%v", position.BorrowShares)
		b, ok := new(big.Int).SetString(bs, 10)
		if !ok {
			continue
		}
		if b.Cmp(big.NewInt(0)) == 0 {
			continue
		}
		c := fmt.Sprintf("%v", position.Collateral)
		cl, ok := new(big.Int).SetString(c, 10)
		if !ok {
			continue
		}
		if cl.Cmp(big.NewInt(0)) == 0 {
			continue
		}
		val := &Position{
			BorrowShares:  b,
			Collateral:    cl,
			HealthyFactor: position.HealthFactor,
		}
		ret[position.User.Address] = val
	}
	return ret, nil
}

func (m *Morpho) getMarketInfo(id string) (*MorphoMarket, error) {
	query := `{markets(where: {uniqueKey_in:["%s"]}) {items {id uniqueKey lltv oracleAddress irmAddress loanAsset {address symbol decimals} collateralAsset {address symbol decimals} state {borrowShares borrowAssets supplyShares supplyAssets fee timestamp utilization}}}}`
	query = fmt.Sprintf(query, id)
	payload := map[string]string{
		"query": query,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling the payload:", err)
		return nil, err
	}
	var res MorphoMarketsResp
	_, err = m.blueClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payloadBytes).
		SetResult(&res).
		Post("")
	if err != nil {
		log.Errorf("get market info failed, err=%v", err)
		return nil, err
	}
	if len(res.Data.Markets.Items) == 0 {
		return nil, errors.New(fmt.Sprintf("market id=%s not found", id))
	}
	return &res.Data.Markets.Items[0], nil
}
