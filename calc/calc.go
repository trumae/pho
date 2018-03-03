package main

import (
	"log"

	"github.com/trumae/pho"
)

func main() {
	grammar := pho.Seq{pho.Some(pho.IsLetter),
		pho.Or{pho.Some(pho.IsLetter),
			pho.Some(pho.IsDigit)}}

	sinput := "ident"
	input := []interface{}{}
	for _, c := range sinput {
		input = append(input, c)
	}

	out, err := grammar.Run(input)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(out)
}
