package logging

import (
	"os"
	"sync"
	"testing"
)

func TestLogLevelDefine(t *testing.T) {
	if DEBUG != 10 {
		t.Errorf("DEBUG: %d", DEBUG)
	}
	if INFO != 20 {
		t.Errorf("INFO: %d", INFO)
	}
	if WARN != 30 {
		t.Errorf("WARN: %d", WARN)
	}
	if ERROR != 40 {
		t.Errorf("ERROR: %d", ERROR)
	}
	if FATAL != 50 {
		t.Errorf("FATAL: %d", FATAL)
	}
}

func TestBaseFormatter(t *testing.T) {
	format := NewBaseFormatter()
	format.SetFormat("[{{.Level}}]:{{.Time}}:{{.ShortFilename}}:{{.Filename}}:{{.Lineno}} {{.Message}} ({{.Name}})")
	format.Parse(os.Stderr, 1, "__main__", DEBUG, "hello")

	// interfaceの実装を確認
	if _, ok := interface{}(format).(Formatter); !ok {
		t.Errorf("BaseFormatterはFormatterとしてのinterfaceを満たしていません．")
	}
}

func TestBaseHandler(t *testing.T) {
	handler := &BaseHandler{
		out:    os.Stderr,
		level:  WARN,
		format: NewBaseFormatter(),
	}
	handler.Output("__main__", INFO, "hello", "world")
	handler.Output("__main__", WARN, "hello", "world")

	// interfaceの実装確認
	if _, ok := interface{}(handler).(Handler); !ok {
		t.Errorf("BaseHandlerはHandlerとしてのinterfaceを満たしていません．")
	}
}

func TestStreamHandler(t *testing.T) {
	handler := NewStreamHandler()
	handler.Output("__main__", INFO, "hello", "world")
	handler.Output("__main__", WARN, "hello", "world")
}

func TestColorStreamHandler(t *testing.T) {
	handler := NewColorStreamHandler()
	handler.Output("__main__", INFO, "hello", "world")
	handler.Output("__main__", WARN, "hello", "world")

	handler.SetLevel(DEBUG)
	handler.Output("__main__", DEBUG, "hello", "world")
	handler.Output("__main__", INFO, "hello", "world")
	handler.Output("__main__", WARN, "hello", "world")
	handler.Output("__main__", ERROR, "hello", "world")
	handler.Output("__main__", FATAL, "hello", "world")
}

func TestLoggerTest(t *testing.T) {
	logger := NewLogger("test")
	logger.AddHandler(NewStreamHandler(), NewColorStreamHandler())
	logger.SetLevel(DEBUG)
	group := new(sync.WaitGroup)
	group.Add(2)
	go func() {
		logger.Debug("this is", "debug")
		logger.Info("this is", "info")
		group.Done()
	}()
	go func() {
		logger.Warn("this is", "warn")
		logger.Error("this is", "error")
		logger.Fatal("this is", "fatal")
		group.Done()
	}()
	group.Wait()
}
