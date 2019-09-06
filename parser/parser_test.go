package parser

import (
	"fmt"
	"testing"
)

func TestParsing(t *testing.T) {
	path := "sample.xml.gz"
	g := NewGNC()
	g.ParseFile(path)

	fmt.Printf("\nThere are %d accounts and %d transactions\n", len(g.Accounts), len(g.Transactions))
	fmt.Printf("Account dump :\n%#v\n", g.Accounts)
	fmt.Printf("Transaction dump :\n%#v\n", g.Transactions)

}
