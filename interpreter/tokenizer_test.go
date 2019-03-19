package interpreter

import (
	"testing"
)

func TestInvalidCharactersAreIgnored(t *testing.T) {
	input := ">>\n+,,..TESTIN[[[]]]VALID+\t\n\t<-CHARACTERS<-"
	got := string(Tokenize(input))
	expected := ">>+,,..[[[]]]+<-<-"
	if got != expected {
		t.Errorf("Got %s, expected %s", got, expected)
	}
}
