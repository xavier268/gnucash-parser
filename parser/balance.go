package parser

import "fmt"

// Balance compute balance for specified account
func (gnc *GNC) Balance(actGUID string) float64 {
	return float64(gnc.Balances[actGUID]) / 100.
}

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

// AccountName get account name from guid
func (gnc *GNC) AccountName(actGUID string) string {
	n, ok := gnc.Accounts[actGUID]
	if !ok {
		return "Unknown account ?! (guid : " + actGUID + ")"
	}
	return n.Name + "(" + n.Description + ")"
}
