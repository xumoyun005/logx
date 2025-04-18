package logx

var defaultLogger = New()

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) WithField(key string, value interface{}) *Entry {
	return NewEntry(l).WithField(key, value)
}

func (l *Logger) WithFields(fields Fields) *Entry {
	return NewEntry(l).WithFields(fields)
}

func WithField(key string, value interface{}) *Entry {
	return defaultLogger.WithField(key, value)
}

func WithFields(fields Fields) *Entry {
	return defaultLogger.WithFields(fields)
}
func Info(args ...interface{}) {
	NewEntry(defaultLogger).Info(args...)
}

func Warn(args ...interface{}) {
	NewEntry(defaultLogger).Warn(args...)
}

func Error(args ...interface{}) {
	NewEntry(defaultLogger).Error(args...)
}
func Panic(args ...interface{}) {
	NewEntry(defaultLogger).Panic(args...)
}