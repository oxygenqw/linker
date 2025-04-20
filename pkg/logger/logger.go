package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return "", fmt.Sprintf("%s:%d", filename, frame.Line) // Убрали вывод функции
		},
		FullTimestamp: false,
	}

	l.SetOutput(io.MultiWriter(os.Stdout))

	e = logrus.NewEntry(l)
}
