package parser

import (
	"fmt"
	"testing"
)

func TestReading(t *testing.T) {
	path := "../../compta2015/comptes2014.gnucash"
	gnc, err := ReadFile(path)
	fmt.Println(err)
	fmt.Printf("\n\n%#v\n", gnc.Book.Accounts[50])
	fmt.Printf("\n\n%#v\n", gnc.Book.Transactions)

}
