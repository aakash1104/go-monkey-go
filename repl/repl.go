package repl

import (
    "fmt"
    "io"
    "bufio"
    "go-monkey-go/token"
    "go-monkey-go/lexer"
)

const PROMPT = ">> "

/**
 * @func Start: Start REPL and repeatedly print out tokens till EOF
 **/
func Start (in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()
        if (!scanned) {
            return
        }
        line := scanner.Text()
        l := lexer.New(line)

        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Printf("%+v\n", tok)
        }
    }
}
