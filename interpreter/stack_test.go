package interpreter

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := newStack()
	stack.push(10)
	stack.push(20)
	second, _ := stack.pop()
	first, _ := stack.pop()
	if first != 10 || second != 20 {
		t.Errorf("got '%v, %v' want '%v, %v'", first, second, 10, 20)
	}
}

func TestPopEmptyStackFails(t *testing.T) {
	stack := newStack()
	_, err := stack.pop()
	if err == nil {
		t.Error("Expected err was not created.")
	}
}
