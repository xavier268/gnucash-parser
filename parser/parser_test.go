package parser

import (
	"fmt"
	"testing"
)

func TestParsing(t *testing.T) {
	//path := "sample.xml.gz"
	path := "../../compta2015/comptes2014.gnucash"
	g := NewGNC()
	g.ParseFile(path)

	fmt.Printf("\nThere are %d accounts and %d transactions\n", len(g.Accounts), len(g.Transactions))
	//fmt.Printf("Account dump :\n%#v\n", g.Accounts)
	//fmt.Printf("Transaction dump :\n%#v\n", g.Transactions)

	guid := "d5789b335f3858a2df8f6be3148ee665"
	fmt.Printf("\n %s ==> %15.2f", g.AccountName(guid), g.Balance(guid))
	fmt.Println(g.Accounts[guid].Child)

	guid = "24bf3d6b788b27db1ac1047eae930eb0"
	fmt.Printf("\n %s ==> %15.2f", g.AccountName(guid), g.Balance(guid))
	fmt.Println(g.Accounts[guid].Child)

	guid = "01ca102dbb628766a4031103815d10fd"
	fmt.Printf("\n %s ==> %15.2f", g.AccountName(guid), g.Balance(guid))
	fmt.Println(g.Accounts[guid].Child)

	guid = "41b423a28ca11c434a4289281310c46f"
	fmt.Printf("\n %s ==> %15.2f", g.AccountName(guid), g.Balance(guid))
	fmt.Println(g.Accounts[guid].Child)

}
