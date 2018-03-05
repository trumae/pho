package pho

import (
	"fmt"
	"log"
	"reflect"
)

var Debug = true

type FPredicate func(interface{}) bool
type FApply func(interface{}) (interface{}, error)

type Operator interface {
	Run([]interface{}) ([]interface{}, error)
}

type Some FPredicate
type One struct {
	Value interface{}
}
type Seq []Operator
type Or []Operator

type Many struct {
	Value Operator
}

func DebugOnInit(t string, input []interface{}) {
	if Debug {
		log.Println(t, input)
	}
}

func DebugOnFinish(t string, input []interface{}) {
	if Debug {
		log.Println("<--", t, input)
	}
}

//Run for Some
func (some Some) Run(input []interface{}) ([]interface{}, error) {
	DebugOnInit("Some", input)
	defer DebugOnFinish("Some", input)

	if len(input) == 0 {
		return nil, fmt.Errorf("Some value expected")
	}

	first := input[0]
	if some(first) {
		return []interface{}{first}, nil
	}

	return nil, fmt.Errorf("Value not expected")
}

//Run for One
func (one One) Run(input []interface{}) ([]interface{}, error) {
	DebugOnInit("One('"+string([]rune{one.Value.(rune)})+"')", input)
	defer DebugOnFinish("One", input)

	pred := func(i interface{}) bool {
		return reflect.DeepEqual(i, one.Value)
	}

	return Some(pred).Run(input)
}

//Run for And
func (seq Seq) Run(input []interface{}) ([]interface{}, error) {
	DebugOnInit("Seq", input)
	defer DebugOnFinish("Seq", input)

	ret := []interface{}{}
	for _, term := range seq {
		o, err := term.Run(input)
		if err != nil {
			return nil, err
		}

		ret = append(ret, o)
		if len(input) > 0 {
			input = input[1:]
		}
	}
	return ret, nil
}

//Run for Or
func (or Or) Run(input []interface{}) ([]interface{}, error) {
	DebugOnInit("Or", input)
	defer DebugOnFinish("Or", input)

	log.Println("--------------- Or content", or)
	for i, term := range or {
		log.Println("----Or content - term", term, i)
		o, err := term.Run(input)
		if err == nil {
			return o, err
		}
	}

	return nil, fmt.Errorf("No match OR")
}

//Run for Many
func (many Many) Run(input []interface{}) ([]interface{}, error) {
	DebugOnInit("Many", input)
	defer DebugOnFinish("Many", input)

	if len(input) == 0 {
		return input, nil
	}

	f, err := many.Value.Run(input[0:1])
	if err != nil {
		return nil, nil
	}

	rest, err := many.Run(input[1:])
	if err != nil {
		return f, nil
	}

	ret := append(f, rest...)
	return ret, nil
}
