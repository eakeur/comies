package telemetry

import (
	"os"
	"sync"

	"go.uber.org/zap"
)

type Telemetry struct {
	Logger *zap.Logger
}

var (
	instrumentation *Telemetry
	once            = sync.Once{}
)

func New(logger *zap.Logger) *Telemetry {
	return &Telemetry{
		Logger: logger,
	}
}

func FakeInstrumentation() *Telemetry {
	return &Telemetry{
		Logger: NewLogger(os.Stdout),
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
