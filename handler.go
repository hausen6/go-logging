package logging

type Handler interface {
	SetLevel(LogLevel)
	SetFormeter(Formatter)
	Output(...interface{})
}
