package logxi

import (
	"github.com/fgrosse/servo"
	"github.com/mgutz/logxi/v1"
)

type loggerAdapter struct {
	log.Logger
	level int
}

func (l *loggerAdapter) SetLevel(level int) {
	l.level = level
	l.Logger.SetLevel(level)
}

func (l *loggerAdapter) Level() int {
	return l.level
}

func (l *loggerAdapter) IsError() bool {
	return l.level <= servo.LevelError
}
