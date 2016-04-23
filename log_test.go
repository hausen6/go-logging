package logging

import (
	"os"
	"testing"
)

func TestLogLevelDefine(t *testing.T) {
	if DEBUG != 10 {
		t.Errorf("DEBUG: %v", DEBUG)
	}
	if INFO != 10 {
		t.Errorf("INFO: %v", INFO)
	}
	if WARN != 10 {
		t.Errorf("WARN: %v", WARN)
	}
	if ERROR != 10 {
		t.Errorf("ERROR: %v", ERROR)
	}
	if FATAL != 10 {
		t.Errorf("FATAL: %v", FATAL)
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
