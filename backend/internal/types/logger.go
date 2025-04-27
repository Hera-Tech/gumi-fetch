package types

type Logger interface {
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Debugf(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
}
