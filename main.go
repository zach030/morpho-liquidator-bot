package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/zach030/morpho-liquidator-bot/pkg"

	"github.com/zach030/morpho-liquidator-bot/config"

	log "github.com/sirupsen/logrus"
)

const (
	defaultConfigPath = "config.yaml"
)

func main() {
	log.Info("morpho liquidator start")
	cfg := config.LoanConfig(defaultConfigPath)
	s := pkg.NewSubscriber(cfg)
	s.Subscribe()
	bot := NewMorpho(cfg)
	go bot.Start(s.Block(), s.PendingTx(), s.Log())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-c
	s.UnSubscribe()
	log.Info("morpho liquidator exit")
}
