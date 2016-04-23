package logging

import "io"

type Handler interface {
	SetLevel(LogLevel)
	SetFormeter(Formatter)
	Output(string, LogLevel, ...interface{})
}

type BaseHandler struct {
	out    io.Writer
	level  LogLevel
	format Formatter
}

func (b *BaseHandler) SetLevel(level LogLevel) {
	b.level = level
}

func (b *BaseHandler) SetFormeter(formatter Formatter) {
	b.format = formatter
}

func (b *BaseHandler) Output(name string, level LogLevel, messages ...interface{}) {
	b.format.Parse(b.out, 2, name, level, messages...)
}
