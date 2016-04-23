package logging

import "fmt"

type LogLevel int

const (
	DEBUG LogLevel = 10 * (1 << (iota + 1))
	INFO
	WARN
	ERROR
	FATAL
)

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBU"
	case INFO:
		return "INFO"
	case WARN:
		return "WORN"
	case ERROR:
		return "ERRO"
	case FATAL:
		return "FATA"
	}
	return fmt.Sprintf("%v", l)
}
