package logging

import "sync"

type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})

	SetLevel(int)
	AddHandler(...Handler)
	RemeveHandlers()
}

type BaseLogger struct {
	mu sync.Mutex

	name     string
	level    LogLevel
	handlers []Handler
}

func NewLogger(name string) *BaseLogger {
	logger := new(BaseLogger)
	logger.name = name
	logger.level = WARN

	return logger
}

func (l *BaseLogger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
	for i, _ := range l.handlers {
		l.handlers[i].SetLevel(level)
	}
}

func (l *BaseLogger) AddHandler(handlers ...Handler) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, h := range handlers {
		l.handlers = append(l.handlers, h)
	}
}

func (l *BaseLogger) RemeveHandlers() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.handlers = []Handler{}
}

func (l *BaseLogger) Debug(messages ...interface{}) {
	if l.level <= DEBUG {
		l.mu.Lock()
		defer l.mu.Unlock()
		for _, h := range l.handlers {
			h.Output(l.name, DEBUG, messages...)
		}
	}
}

func (l *BaseLogger) Info(messages ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= INFO {
		for _, h := range l.handlers {
			h.Output(l.name, INFO, messages...)
		}
	}
}

func (l *BaseLogger) Warn(messages ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= WARN {
		for _, h := range l.handlers {
			h.Output(l.name, WARN, messages...)
		}
	}
}

func (l *BaseLogger) Error(messages ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= ERROR {
		for _, h := range l.handlers {
			h.Output(l.name, ERROR, messages...)
		}
	}
}

func (l *BaseLogger) Fatal(messages ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= FATAL {
		for _, h := range l.handlers {
			h.Output(l.name, FATAL, messages...)
		}
	}
}
