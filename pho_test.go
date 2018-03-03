package pho

import (
	"testing"
)

func TestIsLetter(t *testing.T) {
	const mixedTrue = "abcd"
	const mixedFalse = "\b5̀9!℃"

	for _, c := range mixedTrue {
		b := IsLetter(c)
		if !b {
			t.Error("Expected true for value", c)
		}
	}

	for _, c := range mixedFalse {
		b := IsLetter(c)
		if b {
			t.Errorf("Expected false for value %c", c)
		}
	}
}

func TestSome(t *testing.T) {
	const sinput = "ab"

	input := []interface{}{}
	for _, c := range sinput {
		input = append(input, c)
	}

	grammar := Some(IsLetter)

	_, err := grammar.Run(input)
	if err != nil {
		t.Error(err)
	}

	_, err = grammar.Run(input[1:])
	if err != nil {
		t.Error(err)
	}

	_, err = grammar.Run(input[2:])
	if err == nil {
		t.Error("Empty input - error expected")
	}

}

func TestOne(t *testing.T) {
	const sinput = "ab"

	input := []interface{}{}
	for _, c := range sinput {
		input = append(input, c)
	}

	grammar := One{Value: input[0]}

	_, err := grammar.Run(input)
	if err != nil {
		t.Error(err)
	}

	_, err = grammar.Run(input[1:])
	if err == nil {
		t.Error("error expected")
	}

}

func TestSeq(t *testing.T) {
	const sinput = "ab"

	input := []interface{}{}
	for _, c := range sinput {
		input = append(input, c)
	}

	grammar := Seq{Some(IsLetter), Some(IsLetter)}

	_, err := grammar.Run(input)
	if err != nil {
		t.Error(err)
	}

	grammar = Seq{One{Value: input[0]}, Some(IsLetter)}
	_, err = grammar.Run(input)
	if err != nil {
		t.Error(err)
	}

	grammar = Seq{One{Value: input[1]}, Some(IsLetter)}
	_, err = grammar.Run(input)
	if err == nil {
		t.Error("Error expected")
	}
}

func TestOr(t *testing.T) {
	const sinput = "ab"

	input := []interface{}{}
	for _, c := range sinput {
		input = append(input, c)
	}

	grammar := Or{One{Value: input[0]},
		One{Value: input[1]}}

	_, err := grammar.Run(input)
	if err != nil {
		t.Error(err)
	}
}

func TestMany(t *testing.T) {
	const sinput = "Trumae"
	const sinput2 = "rr2"

	input := []interface{}{}
	for _, c := range sinput {
		input = append(input, c)
	}

	grammar := Many(IsLetter)

	_, err := grammar.Run(input)
	if err != nil {
		t.Error(err)
	}

	input = []interface{}{}
	for _, c := range sinput2 {
		input = append(input, c)
	}

	_, err = grammar.Run(input)
	if err != nil {
		t.Error(err)
	}
}
