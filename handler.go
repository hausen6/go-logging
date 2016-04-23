package logging

import (
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
)

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
	if level >= b.level {
		b.format.Parse(b.out, 2, name, level, messages...)
	}
}

type StreamHandler struct {
	BaseHandler
}

func NewStreamHandler() *StreamHandler {
	handler := new(StreamHandler)
	handler.out = os.Stderr
	handler.level = WARN
	handler.format = NewBaseFormatter()

	return handler
}

type ColorStreamHandler struct {
	BaseHandler
	debug_color *color.Color
	info_color  *color.Color
	warn_color  *color.Color
	error_color *color.Color
	fatal_color *color.Color
}

func NewColorStreamHandler() *ColorStreamHandler {
	handler := new(ColorStreamHandler)
	handler.out = colorable.NewColorableStderr()
	handler.level = WARN
	handler.format = NewBaseFormatter()
	handler.debug_color = color.New(color.FgGreen)
	handler.info_color = color.New(color.FgCyan)
	handler.warn_color = color.New(color.FgYellow)
	handler.error_color = color.New(color.FgRed)
	handler.fatal_color = color.New(color.FgBlack).Add(color.BgRed)

	return handler
}

func (h *ColorStreamHandler) Output(name string, level LogLevel, messages ...interface{}) {
	if level >= h.level {
		if DEBUG <= level && level < INFO {
			h.debug_color.Set()
			defer color.Unset()
			h.format.Parse(h.out, 2, name, level, messages...)
		} else if INFO <= level && level < WARN {
			h.info_color.Set()
			defer color.Unset()
			h.format.Parse(h.out, 2, name, level, messages...)
		} else if WARN <= level && level < ERROR {
			h.warn_color.Set()
			defer color.Unset()
			h.format.Parse(h.out, 2, name, level, messages...)
		} else if ERROR <= level && level < FATAL {
			h.error_color.Set()
			defer color.Unset()
			h.format.Parse(h.out, 2, name, level, messages...)
		} else if FATAL <= level {
			h.fatal_color.Set()
			defer color.Unset()
			h.format.Parse(h.out, 2, name, level, messages...)
		}
	}
}
