package interpreter

import (
	"errors"
)

type stack struct {
	slice []int
}

func newStack() *stack {
	return &stack{
		slice: make([]int, 0, 50),
	}
}

// push pushes a value onto the stack.
// It might have to resize the underlying array.
func (st *stack) push(v int) {
	st.slice = append(st.slice, v)
}

// pop pops a value off the stack.
func (st *stack) pop() (int, error) {
	slice := st.slice
	if len(slice) == 0 {
		return 0, errors.New("Can't pop from empty stack.")
	}
	l := len(slice)
	// Read last value
	val := slice[l-1]
	// Remove last value
	st.slice = slice[:l-1]
	return val, nil
}
