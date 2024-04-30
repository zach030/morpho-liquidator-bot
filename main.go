package main

import (
	"flag"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/zach030/morpho-liquidator-bot/strategy"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/zach030/morpho-liquidator-bot/pkg"

	"github.com/zach030/morpho-liquidator-bot/config"

	log "github.com/sirupsen/logrus"
)

type HighlightFormatter struct {
}

func (f *HighlightFormatter) Format(entry *log.Entry) ([]byte, error) {
	var highlightMsg string
	switch entry.Level {
	case log.InfoLevel:
		highlightMsg = ansi.Color(strings.ToUpper(entry.Level.String()), "cyan+b")
	case log.WarnLevel:
		highlightMsg = ansi.Color(strings.ToUpper(entry.Level.String()), "yellow+b")
	case log.ErrorLevel:
		highlightMsg = ansi.Color(strings.ToUpper(entry.Level.String()), "red+b")
	}
	msg := fmt.Sprintf("[%s] %v %s\n", highlightMsg, time.Now().Format(time.DateTime), entry.Message)
	return []byte(msg), nil
}

var (
	flagconf string
)

func init() {
	log.SetFormatter(&HighlightFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	flag.StringVar(&flagconf, "conf", "config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	log.Info("morpho liquidator start")
	cfg := config.LoanConfig(flagconf)
	s := pkg.NewSubscriber(cfg)
	s.Subscribe()
	bot := strategy.NewMorpho(cfg)
	go bot.Start(s.Block(), s.PendingTx(), s.Log())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-c
	s.UnSubscribe()
	log.Info("morpho liquidator exit")
}
