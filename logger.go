package logging

import (
	"io"
	"sync"
)

type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})

	AddHandler(Handler)
	SetLevel(int)
}

type BaseLogger struct {
	mu  sync.Mutex
	out io.Writer
	buf []byte

	level int
}

func (self *BaseLogger) output(...interface{}) {
	self.mu.Lock()
	defer self.mu.Unlock()
}
