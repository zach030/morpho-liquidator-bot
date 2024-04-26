package pkg

import (
	"context"

	"github.com/zach030/morpho-liquidator-bot/config"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	log "github.com/sirupsen/logrus"
)

var (
	MorphoAddress = common.HexToAddress("0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb")
)

type Subscriber struct {
	ethClient *ethclient.Client
	wsClient  *ethclient.Client
	headers   chan *types.Header
	txs       chan *types.Transaction
	logs      chan *types.Log
	cancel    context.CancelFunc
}

func NewSubscriber(cfg *config.Config) *Subscriber {
	httpClient, err := ethclient.Dial(cfg.HttpEndpoint)
	if err != nil {
		panic(err)
	}
	wsClient, err := ethclient.Dial(cfg.WsEndpoint)
	if err != nil {
		panic(err)
	}
	return &Subscriber{
		ethClient: httpClient,
		wsClient:  wsClient,
		headers:   make(chan *types.Header),
		txs:       make(chan *types.Transaction),
		logs:      make(chan *types.Log),
	}
}

func (s *Subscriber) Subscribe() {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	go s.subscribeNewBlock(ctx)
	go s.subscribeEvent(ctx)
	go s.subscribePendingTx(ctx)
}

func (s *Subscriber) UnSubscribe() {
	s.cancel()
	s.wsClient.Client().Close()
	s.wsClient.Close()
}

func (s *Subscriber) Block() <-chan *types.Header {
	return s.headers
}

func (s *Subscriber) PendingTx() <-chan *types.Transaction {
	return s.txs
}

func (s *Subscriber) Log() <-chan *types.Log {
	return s.logs
}

func (s *Subscriber) subscribeNewBlock(ctx context.Context) {
	headers := make(chan *types.Header)
	sub, err := s.wsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Errorf("[Subscriber] headers subscription error: %v", err)
		return
	}
	defer sub.Unsubscribe()
	for {
		select {
		case <-ctx.Done():
			return
		case err := <-sub.Err():
			log.Errorf("[Subscriber] headers subscription error: %v", err)
			return
		case header := <-s.headers:
			log.Infof("[Subscriber] new headers timestamp=%d number=%v", header.Time, header.Number.Uint64())
			go func() {
				s.headers <- header
			}()
		}
	}
}

func (s *Subscriber) subscribeEvent(ctx context.Context) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			MorphoAddress,
		},
	}
	logs := make(chan types.Log)
	sub, err := s.wsClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Errorf("[Subscriber] event subscription error: %v", err)
		return
	}
	defer sub.Unsubscribe()
	for {
		select {
		case <-ctx.Done():
			log.Warn("[Subscriber] stop subscribe event")
			return
		case err := <-sub.Err():
			log.Errorf("[Subscriber] event subscription error: %v", err)
			return
		case vLog := <-logs:
			log.Debugf("[Subscriber] vLog address=%v txHash=%v block=%v", vLog.Address.Hex(), vLog.TxHash.Hex(), vLog.BlockNumber)
			go func() {
				s.logs <- &vLog
			}()
		}
	}
}

func (s *Subscriber) subscribePendingTx(ctx context.Context) {
	cli := gethclient.New(s.wsClient.Client())
	pending := make(chan *types.Transaction)
	sub, err := cli.SubscribeFullPendingTransactions(context.Background(), pending)
	if err != nil {
		log.Errorf("[Subscriber] pending tx subscription error: %v", err)
		return
	}
	for {
		select {
		case <-ctx.Done():
			log.Warn("[Subscriber] stop subscribe pending tx")
			return
		case err := <-sub.Err():
			log.Errorf("[Subscriber] pending tx subscription error: %v", err)
			return
		case tx := <-pending:
			go func() {
				s.txs <- tx
			}()
		}
	}
}
