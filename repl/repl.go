package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

// Start the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		// Scan a line from the input
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// Get the line from the input
		line := scanner.Text()

		// Create a lexer
		l := lexer.New(line)

		// Print the tokens
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
