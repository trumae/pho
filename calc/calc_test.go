package calc

import (
	"testing"

	"github.com/trumae/pho"
)

func TestIndentifier(t *testing.T) {
	pho.Debug = false

	sinput := "x16 = 8"
	input := []interface{}{}
	for _, c := range sinput {
		input = append(input, c)
	}

	ident := Identifier{}

	_, err := ident.Run(input)
	if err != nil {
		t.Error(err)
	}

	if ident.Value != "x16" {
		t.Error("Identifier not expected", ident.Value)
	}
}

func TestInteger(t *testing.T) {
	pho.Debug = false

	sinput := "878"
	input := []interface{}{}
	for _, c := range sinput {
		input = append(input, c)
	}

	ident := Integer{}

	_, err := ident.Run(input)
	if err != nil {
		t.Error(err)
	}

	if ident.Value != 878 {
		t.Error("Value not expected", ident.Value)
	}
}
