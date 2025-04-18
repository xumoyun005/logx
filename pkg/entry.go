package logrus

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"
)

type Entry struct {
	Logger  *Logger
	Data    Fields
	Time    time.Time
	Level   Level
	Message string
	Context context.Context
}

func NewEntry(logger *Logger) *Entry {
	return &Entry{
		Logger: logger,
		Data:   make(Fields, 6),
	}
}

func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return entry.WithFields(Fields{key: value})
}

func (entry *Entry) WithFields(fields Fields) *Entry {
	data := make(Fields, len(entry.Data)+len(fields))
	for k, v := range entry.Data {
		data[k] = v
	}
	for k, v := range fields {
		if t := reflect.TypeOf(v); t != nil && (t.Kind() == reflect.Func || (t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Func)) {
			continue
		}
		data[k] = v
	}
	return &Entry{
		Logger:  entry.Logger,
		Data:    data,
		Time:    entry.Time,
		Context: entry.Context,
	}
}

func (entry *Entry) Log(level Level, msg string) {
	entry.Level = level
	entry.Time = time.Now()
	entry.Message = msg
	entry.write()
}

func (entry *Entry) write() {
	fieldStr := ""
	for k, v := range entry.Data {
		fieldStr += fmt.Sprintf("%s=%v ", k, v)
	}

	logLine := fmt.Sprintf("[%s] [%s] %s%s",
		entry.Time.Format(time.RFC3339),
		entry.Level.String(),
		fieldStr,
		entry.Message,
	)

	fmt.Fprintln(os.Stdout, logLine)
}
