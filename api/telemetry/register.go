package telemetry

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Telemetry struct {
	Logger *zap.Logger
	SQL    *zap.Logger
}

var (
	instrumentation *Telemetry = &Telemetry{}
	once                       = sync.Once{}
)

func New(logger *zap.Logger) *Telemetry {
	return &Telemetry{
		Logger: logger,
	}
}

func FakeInstrumentation() *Telemetry {
	return &Telemetry{
		Logger: NewLogger(os.Stdout, zapcore.DebugLevel),
		SQL:    NewLogger(os.Stdout, zapcore.DebugLevel),
	}
}

func Register(instr *Telemetry) {
	if instr == nil {
		panic("tried to register nil instrumentation")
	}
	instrumentation = instr
}

func RegisterOnce(instr *Telemetry) {
	once.Do(func() {
		Register(instr)
	})
}
