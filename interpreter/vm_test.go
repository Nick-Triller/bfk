package interpreter

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestHelloWorldProgram(t *testing.T) {
	in, err := ioutil.ReadFile("../programs/helloworld.bfk")
	if err != nil {
		t.Errorf("Error reading file: %s", err)
		return
	}
	program := string(in)
	var b bytes.Buffer

	if err := Execute(program, os.Stdin, &b); err != nil {
		t.Errorf("Error %s", err)
	}
	got := b.String()

	expected := "Hello World!\n"
	if got != expected {
		t.Errorf("Got %s, expected %s", got, expected)
	}
}

func TestSquaresProgram(t *testing.T) {
	in, err := ioutil.ReadFile("../programs/squares.bfk")
	if err != nil {
		t.Errorf("Error reading file: %s", err)
		return
	}
	program := string(in)
	var b bytes.Buffer

	if err := Execute(program, os.Stdin, &b); err != nil {
		t.Errorf("Error %s", err)
	}
	got := b.String()

	// Generate expected
	var sb strings.Builder
	for i := 0; ; i++ {
		v := i * i
		if v > 10000 {
			break
		}
		sb.WriteString(strconv.Itoa(v) + "\n")
	}
	expected := sb.String()

	// Assert
	if got != expected {
		t.Errorf("Got %s, expected %s", got, expected)
	}
}

func TestUnmatchedBracketsFail(t *testing.T) {
	program := "[>][<]]"
	if err := Execute(program, os.Stdin, os.Stdout); err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestInputInstruction(t *testing.T) {
	program := ",."
	in := strings.NewReader("g")
	var b bytes.Buffer
	if err := Execute(program, in, &b); err != nil {
		t.Errorf("Error: %s", err)
	}
	got := b.String()
	expected := "Enter character: g"
	if got != expected {
		t.Errorf("Got %s, expected %s", got, expected)
	}
}
