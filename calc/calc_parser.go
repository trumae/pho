package calc

import "github.com/trumae/pho"

type Expression struct{}
type Term struct{}
type Factor struct{}

///////////////////////////////////
///      Factor
///////////////////////////////////

//Run for Factor
func (factor *Factor) Run(input []interface{}) ([]interface{}, error) {
	pho.DebugOnInit("Factor", input)

	s := []rune("()")
	lparen := s[0]
	rparen := s[1]

	grammar := pho.Or{pho.Seq{pho.One{Value: lparen}, &Expression{}, pho.One{Value: rparen}},
		&Integer{}}

	out, err := grammar.Run(input)

	return out, err
}

///////////////////////////////////
///         Term
///////////////////////////////////

//Run for Term
func (term *Term) Run(input []interface{}) ([]interface{}, error) {
	pho.DebugOnInit("Term", input)

	times := []rune("*")[0]
	grammar := pho.Or{pho.Seq{&Factor{}, pho.One{Value: times}, &Term{}},
		&Factor{}}

	out, err := grammar.Run(input)

	return out, err
}

///////////////////////////////////
///   Expression
///////////////////////////////////

//Run for Expression
func (expr *Expression) Run(input []interface{}) ([]interface{}, error) {
	pho.DebugOnInit("Expression", input)

	plus := []rune("+")[0]
	grammar := pho.Or{pho.Seq{&Term{}, pho.One{Value: plus}, &Expression{}},
		&Term{}}

	out, err := grammar.Run(input)

	return out, err
}
