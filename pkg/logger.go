package logrus

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Info(args ...interface{}) {
	NewEntry(l).Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	NewEntry(l).Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	NewEntry(l).Error(args...)
}

func (l *Logger) WithField(key string, value interface{}) *Entry {
	return NewEntry(l).WithField(key, value)
}

func (l *Logger) WithFields(fields Fields) *Entry {
	return NewEntry(l).WithFields(fields)
}
