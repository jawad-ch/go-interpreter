package repl

import (
	"bufio"
	"fmt"
	"github.com/jawad-ch/go-interpreter/lexer"
	"github.com/jawad-ch/go-interpreter/parser"
	"io"
)

const PROMPT = ">> "

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		_, _ = io.WriteString(out, program.String())
		_, _ = io.WriteString(out, "\n")

		//for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		//	fmt.Printf("%+v\n", tok)
		//}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	_, _ = io.WriteString(out, MONKEY_FACE)
	_, _ = io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	_, _ = io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
