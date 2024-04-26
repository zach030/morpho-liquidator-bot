package main

import (
	"fmt"
	"github.com/mgutz/ansi"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

// HighlightFormatter is a custom formatter that adds highlighting to log messages
type HighlightFormatter struct {
}

// Format formats the log entry and adds highlighting to the message
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

func init() {
	// Receipt as JSON instead of the default ASCII formatter.
	log.SetFormatter(&HighlightFormatter{})

	// Output to stdout instead of the default stderr.
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}
