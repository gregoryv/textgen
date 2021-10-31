// Command english provides ways to query english words.
package main

import (
	"fmt"
	"strings"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/english"
)

func main() {
	var (
		cli = cmdline.NewBasicParser()
		p   = cli.Option("-p, --print", "Print all words").Bool()
		r   = cli.Option("-r, --random", "Print random words").Int(0)
		rs  = cli.Option("-rs, --random-statement", "Print random statement").Int(0)
		rq  = cli.Option("-rq, --random-question", "Print random question").Int(0)
	)
	cli.Parse()

	switch {
	case p:
		fmt.Println(strings.Join(english.Words(), " "))

	case r > 0:
		fmt.Println(strings.Join(english.RandomWords(r), " "))

	case rs > 0:
		words := english.RandomStatement(rs, rs)
		fmt.Println(english.Sentence(words, '.'))

	case rq > 0:
		words := english.RandomQuestion(rq, rq)
		fmt.Println(english.Sentence(words, '?'))

	default:
		fmt.Println(len(english.Words()), "words")
	}
}
