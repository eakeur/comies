package conn

import (
	"fmt"

	"go.uber.org/zap"
)

type logger struct {
	log *zap.Logger
}

func (l logger) Printf(format string, v ...interface{}) {
	l.log.Info(fmt.Sprintf(format, v...))
}

func (l logger) Verbose() bool {
	return true
}
