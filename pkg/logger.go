package logrus

var defaultLogger = New()

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) WithField(key string, value interface{}) *Entry {
	return NewEntry(l).WithField(key, value)
}

// WithFields adds multiple fields to the entry
func (l *Logger) WithFields(fields Fields) *Entry {
	return NewEntry(l).WithFields(fields)
}

func WithField(key string, value interface{}) *Entry {
	return defaultLogger.WithField(key, value)
}

func WithFields(fields Fields) *Entry {
	return defaultLogger.WithFields(fields)
}