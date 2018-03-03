package pho

import (
	"log"
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

func TestAnd(t *testing.T) {
	const sinput = "ab"

	input := []interface{}{}
	for _, c := range sinput {
		input = append(input, c)
	}

	grammar := And{Some(IsLetter), Some(IsLetter)}

	_, err := grammar.Run(input)
	if err != nil {
		t.Error(err)
	}

	grammar = And{One{Value: input[0]}, Some(IsLetter)}
	_, err = grammar.Run(input)
	if err != nil {
		t.Error(err)
	}

	grammar = And{One{Value: input[1]}, Some(IsLetter)}
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

	log.Println(grammar)

	out, err := grammar.Run(input)
	if err != nil {
		t.Error(err)
	}

}