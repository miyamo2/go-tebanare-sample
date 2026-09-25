package collection

// Stack is a generic LIFO container.
type Stack[T any] struct {
	items []T
	n     int
}

// NewStack creates an empty Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Len returns the number of items on the stack. Matches getter: a generic
// receiver counts, same as a concrete one.
func (s *Stack[T]) Len() int { return s.n }

// Push adds v to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
	s.n++
}

// Peek returns the top item without removing it. Not a getter: it returns
// two values and indexes a field.
func (s *Stack[T]) Peek() (T, bool) {
	if s.n == 0 {
		var zero T
		return zero, false
	}
	return s.items[s.n-1], true
}
