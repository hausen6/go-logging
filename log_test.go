package logging

import (
	"os"
	"testing"
)

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
