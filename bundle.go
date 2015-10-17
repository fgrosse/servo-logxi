package logxi

import (
	"github.com/fgrosse/servo"
"github.com/mgutz/logxi/v1"
)

type Bundle struct{}

func (b *Bundle) Boot(kernel *servo.Kernel) {
	level := servo.LevelInfo
	if kernel.Log != nil {
		level = kernel.Log.Level()
	}

	kernel.Log = &loggerAdapter{Logger: log.New("kernel")}
	kernel.Log.SetLevel(level)

	RegisterTypes(kernel.TypeRegistry)
}
