package interpreter

import (
	"bytes"
	"io"
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
	input := string(in)
	output := captureStdout(func() {
		Execute(input)
	})
	expected := "Hello World!\n"
	if output != expected {
		t.Errorf("Got %s, expected %s", output, expected)
	}
}

func TestSquaresProgram(t *testing.T) {
	in, err := ioutil.ReadFile("../programs/squares.bfk")
	if err != nil {
		t.Errorf("Error reading file: %s", err)
		return
	}
	input := string(in)
	output := captureStdout(func() {
		Execute(input)
	})

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
	if output != expected {
		t.Errorf("Got %s, expected %s", output, expected)
	}
}

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
