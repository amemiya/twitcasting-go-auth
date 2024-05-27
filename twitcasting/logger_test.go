package twitcasting_test

import (
	"bytes"
	"github.com/amemiya/twitcasting-go-auth/twitcasting"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLoggerHappyPath(t *testing.T) {
	logger := &twitcasting.BasicLogger{
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
	var buf bytes.Buffer
	logger.Logger.SetOutput(&buf)
	logger.Logger.SetFlags(0)
	defer func() {
		logger.Logger.SetOutput(os.Stderr)
		logger.Logger.SetFlags(logger.Logger.Flags())
	}()
	logger.Debug("debug", 1, twitcasting.ErrorResponse{Error: twitcasting.Error{Code: 1000, Message: "test_message"}})
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	assert.Equal(t, strings.Join(
		[]string{
			"debug 1 {{1000 test_message}}",
			"info",
			"warn",
			"error",
		}, "\n")+"\n",
		buf.String(),
	)
}
