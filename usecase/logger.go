package usecase

// Logger is the logging port the usecase layer depends on. Adapters in
// infra/logger implement it.
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Error(msg string)
	Flush()
}
