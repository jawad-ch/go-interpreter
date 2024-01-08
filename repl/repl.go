package repl

import (
	"github.com/jawad-ch/go-interpreter/compiler"
	"github.com/jawad-ch/go-interpreter/object"
	"github.com/jawad-ch/go-interpreter/vm"
	"io"
)

// Read-Parse-Print-Loop
// -----Eval------------
import (
	"bufio"
	"fmt"
	"github.com/jawad-ch/go-interpreter/lexer"
	"github.com/jawad-ch/go-interpreter/parser"
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
	//env := object.NewEnvironment()
	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)
	symbolTable := compiler.NewSymbolTable()

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

		comp := compiler.NewWithState(symbolTable, constants)
		err := comp.Compile(program)

		if err != nil {
			_, _ = fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}

		code := comp.Bytecode()
		constants = code.Constants

		machine := vm.NewWithGlobalsStore(code, globals)
		err = machine.Run()
		if err != nil {
			_, _ = fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
		}

		stackTop := machine.LastPoppedStackElem()
		_, _ = io.WriteString(out, stackTop.Inspect())
		_, _ = io.WriteString(out, "\n")

		//evaluated := evaluator.Eval(program, env)
		//if evaluated != nil {
		//	_, _ = io.WriteString(out, evaluated.Inspect())
		//	_, _ = io.WriteString(out, "\n")
		//}

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
