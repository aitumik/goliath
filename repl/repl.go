package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/aitumik/interpreter/lexer"
	"github.com/aitumik/interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		p := parser.New(l)

		prog := p.ParseProgram()

		fmt.Println(len(prog.Statements))
	}
}
