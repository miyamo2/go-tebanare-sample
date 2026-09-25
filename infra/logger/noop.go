package logger

// Noop discards every message it receives. Useful for tests and for
// composing the application without a real logging backend.
type Noop struct{}

// Debug discards msg. Not a noop match: it takes a parameter.
func (Noop) Debug(msg string) {}

// Info discards msg. Not a noop match: it takes a parameter.
func (Noop) Info(msg string) {}

// Error discards msg. Not a noop match: it takes a parameter.
func (Noop) Error(msg string) {}

// Flush does nothing: Noop buffers nothing to flush.
func (Noop) Flush() {}
