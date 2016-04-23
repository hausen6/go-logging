package logging

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"
)

type LogFormat struct {
	Name          string
	Time          string
	Filename      string
	ShortFilename string
	Lineno        int
	Level         LogLevel
	Message       interface{}
}

type Formatter interface {
	SetFormat(string)
	SetTimeFormat(string)
	Parse(io.Writer, int, string, LogLevel, ...interface{})
}

type BaseFormatter struct {
	format     string
	timeFormat string
}

func NewBaseFormatter() *BaseFormatter {
	formatter := new(BaseFormatter)
	formatter.format = "[{{.Level}}] {{.Message}}"
	formatter.timeFormat = "2006/01/02 03:04:05"

	return formatter
}

func (self *BaseFormatter) SetFormat(format string) {
	self.format = format
}

func (self *BaseFormatter) SetTimeFormat(format string) {
	self.timeFormat = format
}

func (self *BaseFormatter) Parse(out io.Writer, skip int, name string, level LogLevel, messages ...interface{}) {
	var err error
	tmpl, err := template.New("logformat").Parse(self.format + "\n")
	if err != nil {
		log.Printf("incorrect Log Format has been set. => %v\n", self.format)
	}

	filename, shortfilename, lineno := getFileInfo(skip + 1)
	msg := make([]string, len(messages))
	for i, v := range messages {
		msg[i] = v.(string)
	}
	format := LogFormat{
		Name:          name,
		Time:          time.Now().Format(self.timeFormat),
		Filename:      filename,
		ShortFilename: shortfilename,
		Lineno:        lineno,
		Level:         level,
		Message:       strings.Join(msg, ", "),
	}

	err = tmpl.Execute(out, format)
	if err != nil {
		log.Println(err)
	}
}

func getFileInfo(skip int) (filename string, shortfilename string, lineno int) {
	var err error

	_, filename, lineno, ok := runtime.Caller(skip)
	if !ok {
		filename = "???"
		shortfilename = "???"
		lineno = -1
	}
	pwd, err := os.Getwd()
	if err != nil {
		shortfilename = "???"
		return
	}
	shortfilename, err = filepath.Rel(pwd, filename)
	if err != nil {
		shortfilename = "???"
	}
	return
}
