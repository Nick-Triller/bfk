package interpreter

import (
	"testing"
)

func TestInvalidCharactersAreIgnored(t *testing.T) {
	input := ">>\n+,,..TESTIN[[[]]]VALID+\t\n\t<-CHARACTERS<-"
	cleaned := string(Tokenize(input))
	expected := ">>+,,..[[[]]]+<-<-"
	if cleaned != expected {
		t.Errorf("Got %s, expected %s", cleaned, expected)
	}
}
