package parser

import (
	"fmt"
	"testing"
)

var g *GNC
var dates = []string{"2010", "2011", "2012", "2013", "2014",
	"2015", "2016", "2017", "2018", "2019", "2020", "2021"}
var prefix = "\n"
var accts = []string{
	"d5789b335f3858a2df8f6be3148ee665",
	"24bf3d6b788b27db1ac1047eae930eb0",
	"01ca102dbb628766a4031103815d10fd",
	"41b423a28ca11c434a4289281310c46f",
}

func init() {
	// Initialize and run all tests
	//path := "sample.xml.gz"
	path := "../../compta2015/comptes2014.gnucash"
	g = NewGNC()
	g.ParseFile(path)

}

func TestParseIsOK(t *testing.T) {
	if g == nil {
		t.Fatal("The parse should have finished ?")
	}
	if len(g.Roots) != 1 {
		t.Fatal("There should be exactly 1 Root account")
	}

	g.PrintStats()
	g.PrintRoots()
	fmt.Println()
}
func TestBalances(t *testing.T) {

	guid := "d5789b335f3858a2df8f6be3148ee665"
	if g == nil {
		t.Fatal("The parse should have finished ?")
	}
	g.PrintAccountOwnBalance("\n", guid)

	guid = "24bf3d6b788b27db1ac1047eae930eb0"
	g.PrintAccountOwnBalance("\n", guid)

	guid = "01ca102dbb628766a4031103815d10fd"
	g.PrintAccountOwnBalance("\n", guid)

	guid = "41b423a28ca11c434a4289281310c46f"
	g.PrintAccountOwnBalance("\n", guid)

	fmt.Println()

}

func TestBalanceByDate(t *testing.T) {

	for _, a := range accts {
		g.PrintBalanceByDate(prefix, a, dates...)
		fmt.Println()
	}
}

func TestAllBalaces(t *testing.T) {

	g.PrintAccountBalances("\n", g.Roots[0])

	fmt.Println()
}
