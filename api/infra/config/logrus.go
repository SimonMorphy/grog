package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type HighlightFormatter struct{}

func (f *HighlightFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := map[logrus.Level]string{
		logrus.DebugLevel: "\033[36m", // blue
		logrus.InfoLevel:  "\033[32m", // green
		logrus.WarnLevel:  "\033[33m", // yellow
		logrus.ErrorLevel: "\033[31m", // red
		logrus.FatalLevel: "\033[35m", // purple
		logrus.PanicLevel: "\033[41m", // red
	}
	levelEmoji := map[logrus.Level]string{
		logrus.DebugLevel: "🐛",
		logrus.InfoLevel:  "ℹ️ ",
		logrus.WarnLevel:  "⚠️ ",
		logrus.ErrorLevel: "❌",
		logrus.FatalLevel: "💀",
		logrus.PanicLevel: "🔥",
	}
	color := levelColor[entry.Level]
	emoji := levelEmoji[entry.Level]
	reset := "\033[0m"
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	file := ""
	if entry.HasCaller() {
		file = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
	}
	msg := fmt.Sprintf("%s%s [%s] [%s] [%s] %s%s\n",
		color, emoji, entry.Level.String(), timestamp, file, entry.Message, reset)
	return []byte(msg), nil
}

// NewLogrus initializes the logrus logger with custom settings.
func NewLogrus() {
	// standard output
	logrus.SetOutput(os.Stdout)

	// level setting
	logrus.SetLevel(logrus.InfoLevel)

	// style setting
	logrus.SetFormatter(&HighlightFormatter{})

	// caller setting
	logrus.SetReportCaller(true)

	// json
	// logrus.SetFormatter(&logrus.JSONFormatter{})
}
