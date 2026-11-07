package logger

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"app/pkg/requestid"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// TextFormatter prints colored, human-readable lines for local development.
// Structured fields (request_id, status, ...) are appended as key=value pairs.
type TextFormatter struct{}

var levelList = []string{
	"PANIC",
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
	"TRACE",
}

func (TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	level := levelList[int(entry.Level)]
	switch level {
	case "ERROR":
		level = color.RedString("%s ", " ERROR")
	case "WARN":
		level = color.YellowString("%s ", " WARN")
	case "INFO":
		level = color.GreenString("%s ", " INFO")
	case "DEBUG":
		level = color.HiBlueString("%s ", " DEBUG")
	default:
		level = color.HiMagentaString("%s ", level)
	}

	caller := ""
	if entry.Caller != nil {
		caller = fmt.Sprintf("%s - [line:%d] - ", entry.Caller.File, entry.Caller.Line)
	}
	fmt.Fprintf(b, "%s - %s%s msg ▶ %s", entry.Time.Format("2006-01-02 15:04:05.000"), caller, level, entry.Message)

	keys := make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(b, " %s=%v", k, entry.Data[k])
	}
	b.WriteByte('\n')
	return b.Bytes(), nil
}

// requestIDHook copies the request ID from the entry's context (set with
// logger.WithContext(ctx)) into the "request_id" field, so every line written
// on behalf of a request carries the same ID.
type requestIDHook struct{}

func (requestIDHook) Levels() []logrus.Level { return logrus.AllLevels }

func (requestIDHook) Fire(entry *logrus.Entry) error {
	if entry.Context == nil {
		return nil
	}
	if _, exists := entry.Data["request_id"]; exists {
		return nil
	}
	if id := requestid.FromContext(entry.Context); id != "" {
		entry.Data["request_id"] = id
	}
	return nil
}

// ParseLevel converts a LOG_LEVEL value into a logrus level. An empty or
// unknown value yields info.
func ParseLevel(level string) logrus.Level {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

// NewLogger builds the application logger. With json=true (production) it emits
// one JSON object per line, Cloud Logging friendly; otherwise colored text.
func NewLogger(level string, json bool) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(ParseLevel(level))
	logger.AddHook(requestIDHook{})

	if json {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyLevel: "severity",
				logrus.FieldKeyMsg:   "message",
			},
		})
		return logger
	}

	logger.SetReportCaller(true)
	logger.SetFormatter(TextFormatter{})
	return logger
}
