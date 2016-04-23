package logging

import (
	"os"
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
		level:  DEBUG,
		format: NewBaseFormatter(),
	}
	handler.Output("__main__", INFO, "hello", "world")
}
