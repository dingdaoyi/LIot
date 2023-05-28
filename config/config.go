package config

import (
	"github.com/rs/zerolog"
	"os"
	"sync"
)

var (
	logger     zerolog.Logger
	loggerLock = sync.Once{}
)

func GetLogger() *zerolog.Logger {
	loggerLock.Do(func() {
		logger = zerolog.New(os.Stdout)
	})
	return &logger
}
