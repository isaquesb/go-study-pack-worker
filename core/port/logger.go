package port

type Logger interface {
	Log(Message string, Context ...string)
	Debug(Message string, Context ...string)
	Notice(Message string, Context ...string)
	Info(Message string, Context ...string)
	Warning(Message string, Context ...string)
	Error(Error error, Context ...string)
	Critical(Error error, Context ...string)
	Alert(Error error, Context ...string)
	Emergency(Error error, Context ...string)
}
