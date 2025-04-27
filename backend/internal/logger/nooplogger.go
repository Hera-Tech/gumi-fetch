package logger

type NoOpLogger struct{}

func (NoOpLogger) Infof(string, ...interface{})  {}
func (NoOpLogger) Errorf(string, ...interface{}) {}
func (NoOpLogger) Debugf(string, ...interface{}) {}
func (NoOpLogger) Warnf(string, ...interface{})  {}
func (NoOpLogger) Infow(string, ...interface{})  {}
func (NoOpLogger) Errorw(string, ...interface{}) {}
func (NoOpLogger) Debugw(string, ...interface{}) {}
func (NoOpLogger) Warnw(string, ...interface{})  {}
