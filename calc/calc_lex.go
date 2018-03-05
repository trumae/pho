package calc

import (
	"log"
	"strconv"

	"github.com/trumae/pho"
)

///////////////////////////////////
///   Identifier
///////////////////////////////////
type Identifier struct {
	Value string
}

//Run for Identifier
func (ident *Identifier) Run(input []interface{}) ([]interface{}, error) {
	pho.DebugOnInit("Identifier", input)

	grammar := pho.Seq{pho.Some(pho.IsLetter),
		pho.Many{
			Value: pho.Or{pho.Some(pho.IsLetter),
				pho.Some(pho.IsDigit)}}}

	out, err := grammar.Run(input)

	first := out[0]
	rest := out[1:][0]

	val := []rune{}
	val = append(val, first.([]interface{})[0].(rune))
	for _, c := range rest.([]interface{}) {
		val = append(val, c.(rune))
	}

	ident.Value = string(val)

	return []interface{}{ident.Value}, err
}

///////////////////////////////////
///   Integer
///////////////////////////////////
type Integer struct {
	Value int64
}

//Run for Integer
func (integer *Integer) Run(input []interface{}) ([]interface{}, error) {
	pho.DebugOnInit("Integer", input)

	grammar := pho.Seq{
		pho.Many{Value: pho.Some(pho.IsSpace)},
		pho.Some(pho.IsDigit),
		pho.Many{Value: pho.Some(pho.IsDigit)}}

	out, err := grammar.Run(input)

	idx := 0
	for _, r := range out {
		if len(r.([]interface{})) > 0 {
			break
		}
		idx++
	}
	first := out[idx]
	rest := out[idx+1:][0]

	val := []rune{}
	val = append(val, first.([]interface{})[0].(rune))
	for _, c := range rest.([]interface{}) {
		val = append(val, c.(rune))
	}

	s := string(val)

	log.Println("OUT Integer", out, first, rest, val, s)

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err

	}

	integer.Value = i

	return []interface{}{integer.Value}, err
}
