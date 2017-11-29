package imagespy

// Logger provides logging.
type Logger interface {
	Debug(args ...interface{})
}

// NullLogger is noop.
type NullLogger struct{}

// Debug implements Logger.
func (l *NullLogger) Debug(args ...interface{}) {}

// Log is used by this library to log messages.
var Log Logger = &NullLogger{}
