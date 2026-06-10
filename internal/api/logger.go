package api

type Logger interface {
	Info(message string, args ...any)
	Warn(message string, args ...any)
	Error(message string, args ...any)

	With(args ...any) Logger
}
