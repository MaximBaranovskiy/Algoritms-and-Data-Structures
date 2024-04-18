package dataStructures

import "errors"

var ErrStackEmpty = errors.New("stack is empty")

type Stack[T any] struct {
	keys []T
}

// NewStack creates a new stack and returns a pointer to it.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

// Push adds an element to the top of the stack.
func (stack *Stack[T]) Push(key T) {
	stack.keys = append(stack.keys, key)
}

// Empty checks if the stack is empty. It returns true if the stack is empty, and false otherwise.
func (stack *Stack[T]) Empty() bool {
	return len(stack.keys) == 0
}

// Pop removes the top element from the stack and returns it.
// If the stack is empty, it returns a zero value for type T and an error.
func (stack *Stack[T]) Pop() (T, error) {
	if stack.Empty() {
		return *new(T), ErrStackEmpty
	}

	val := stack.keys[len(stack.keys)-1]
	stack.keys = stack.keys[:len(stack.keys)-1]
	return val, nil
}

// Top returns the top element of the stack without removing it.
// If the stack is empty, it returns a zero value for type T and an error.
func (stack *Stack[T]) Top() (T, error) {

	if stack.Empty() {
		return *new(T), ErrStackEmpty
	}

	return stack.keys[len(stack.keys)-1], nil
}
