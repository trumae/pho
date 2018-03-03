package pho

import (
	"fmt"
	"reflect"
)

type FPredicate func(interface{}) bool
type FProc func(interface{}) interface{}
type FApply func(interface{}) interface{}

type Operator interface {
	Run([]interface{}) ([]interface{}, error)
}

type Some FPredicate
type One struct {
	Value interface{}
}
type Seq []Operator
type Or []Operator
type Many FPredicate

type Then FProc
type Apply FApply

//Run for Some
func (some Some) Run(input []interface{}) ([]interface{}, error) {
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
	pred := func(i interface{}) bool {
		return reflect.DeepEqual(i, one.Value)
	}

	return Some(pred).Run(input)
}

//Run for And
func (seq Seq) Run(input []interface{}) ([]interface{}, error) {
	ret := []interface{}{}
	for _, term := range seq {
		o, err := term.Run(input)
		if err != nil {
			return nil, err
		}

		ret = append(ret, o)
		input = input[1:]
	}
	return ret, nil
}

//Run for Or
func (o Or) Run(input []interface{}) ([]interface{}, error) {
	for _, term := range o {
		o, err := term.Run(input)
		if err == nil {
			return o, err
		}
	}
	return nil, fmt.Errorf("No match OR")
}
