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

func TestIsEmptyFalse(t *testing.T) {
	stack := newStack()
	stack.push(5)
	if stack.isEmpty() {
		t.Error("empty() should return false.")
	}
}

func TestIsEmptyTrue(t *testing.T) {
	stack := newStack()
	test := func() {
		if !stack.isEmpty() {
			t.Error("empty() should return true.")
		}
	}
	test()
	stack.push(5)
	stack.pop()
	test()
}
