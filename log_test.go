package logging

import (
	"os"
	"testing"
)

func TestBaseFormatter(t *testing.T) {
	format := NewBaseFormatter(os.Stdout)
	format.SetFormat("[{{.Level}}]:{{.Time}}:{{.ShortFilename}}:{{.Filename}}:{{.Lineno}} {{.Message}} ({{.Name}})")
	format.Parse(1, "__main__", DEBUG, "hello")
}
