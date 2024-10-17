package log

type Logger interface {
	Info(string)
	Debug(string)
	Error(string)
	Fatal(string)
	Infof(string, ...interface{})
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
}
