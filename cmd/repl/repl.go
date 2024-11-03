package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/cyf0rk/interpreter/pkg/lexer"
	"github.com/cyf0rk/interpreter/pkg/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%s\n", tok.Literal)
		}
	}
}
