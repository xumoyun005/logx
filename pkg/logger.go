package logrus

var defaultLogger = New()

type Logger struct{}

// New creates a new logger instance
func New() *Logger {
	return &Logger{}
}

// Instance methods
func (l *Logger) WithField(key string, value interface{}) *Entry {
	return NewEntry(l).WithField(key, value)
}

func (l *Logger) WithFields(fields Fields) *Entry {
	return NewEntry(l).WithFields(fields)
}

// Package-level functions
func WithField(key string, value interface{}) *Entry {
	return defaultLogger.WithField(key, value)
}

func WithFields(fields Fields) *Entry {
	return defaultLogger.WithFields(fields)
}
