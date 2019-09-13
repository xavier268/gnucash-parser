package parser

import "fmt"

// initBalance compute all balances for all accounts.
func (gnc *GNC) initBalance() {
	fmt.Println("Initializing balances ...")
	gnc.Balances = make(map[string]int)
	for _, t := range gnc.Transactions {
		for a, v := range t.Splits {
			gnc.Balances[a] = gnc.Balances[a] + v
		}
	}
	//fmt.Println(gnc.Balances)
}

// Balance compute balance for specified account
func (gnc *GNC) Balance(actGUID string) float64 {
	return float64(gnc.Balances[actGUID]) / 100.
}

// CumulBalance provides the balance of the account AND its child
func (gnc *GNC) CumulBalance(actGUID string) float64 {

	a, ok := gnc.Accounts[actGUID]
	if !ok {
		return 0.0
	}
	b := gnc.Balance(actGUID)
	for _, c := range a.Child {
		b += gnc.CumulBalance(c)
	}
	return b
}

// BalanceAtDatePosted gives the balance at the provided date.
// Date is specified as YYYY-MM-DD, or only part of that string.
func (gnc *GNC) BalanceAtDatePosted(actGUID string, date string) float64 {

	b := 0
	for _, t := range gnc.Transactions {
		if t.DatePosted <= date {
			for a, v := range t.Splits {
				if a == actGUID {
					b += v
				}
			}
		}
	}
	return float64(b) / 100.
}
