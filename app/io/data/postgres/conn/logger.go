package conn

import "log"

type logger struct{}

func (l logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l logger) Verbose() bool {
	return true
}
