package logx

import "strings"

type Level uint8

const (
	InfoLevel Level = iota
	WarnLevel
	ErrorLevel
	PanicLevel
)

func (l Level) String() string {
	switch l {
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	default:
		return "UNKNOWN"
	}
}

func ParseLevel(level string) Level {
	switch strings.ToLower(level) {
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	default:
		return InfoLevel
	}
}
