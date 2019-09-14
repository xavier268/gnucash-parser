package parser

import (
	"fmt"
	"sort"
)

// PrintAccountOwnBalance displays account own balance,
// without its child accounts
func (gnc *GNC) PrintAccountOwnBalance(prefix string, actGUID string) {
	fmt.Printf("%s%-60.60s\t:%20.2f €", prefix, gnc.AccountName(actGUID), gnc.Balance(actGUID))
}

// PrintAccountBalances includes the child accounts
func (gnc *GNC) PrintAccountBalances(prefix string, actGUID string) {
	gnc.PrintAccountOwnBalance(prefix, actGUID)
	fmt.Printf(" (total : %8.2f)", gnc.CumulBalance(actGUID))
	pp := prefix + "  "
	a, ok := gnc.Accounts[actGUID]
	if ok {
		for _, g := range a.Child {
			gnc.PrintAccountBalances(pp, g)
		}
	}
}

// PrintBalanceByDate display balance at various POSTED dates
func (gnc *GNC) PrintBalanceByDate(
	prefix string,
	actGUID string,
	dates ...string) {
	for _, d := range dates {
		fmt.Printf("%s[%s] %-40.40s\t:%20.2f €",
			prefix, d,
			gnc.AccountName(actGUID),
			gnc.BalanceAtDatePosted(actGUID, d))
	}
}

// AccountName get account name from guid
func (gnc *GNC) AccountName(actGUID string) string {
	n, ok := gnc.Accounts[actGUID]
	if !ok {
		return "?! guid : " + actGUID
	}
	return n.Name + "(" + n.Description + " " + actGUID + ")"
}

// PrintRoots list the root (ie, no parent) accounts
func (gnc *GNC) PrintRoots() {
	fmt.Printf("\nRoot accounts are : ")
	for _, guid := range gnc.Roots {
		fmt.Printf("\n   %20.20s\t:\t%s", guid, gnc.AccountName(guid))
	}
}

// PrintStats prints statistics
func (gnc *GNC) PrintStats() {
	fmt.Printf("\nThere are %d accounts and %d transactions\nLast entered %s, last posted %s",
		len(gnc.Accounts), len(gnc.Transactions), gnc.LastEntered, gnc.LastPosted)
}

// PrintAccountDetails display all splits for account, and cumulative balance
// Sort by date
func (gnc *GNC) PrintAccountDetails(actid string) {
	type resentry struct {
		string  // date posted
		float64 // value
	}
	var res []resentry

	fmt.Printf("\nDetails for : %s\n    Date\t  Amount\t\t  Total", gnc.AccountName(actid))

	for _, t := range gnc.Transactions {
		for a, s := range t.Splits {
			if a == actid {
				v := float64(s) / 100
				res = append(res, resentry{t.DatePosted, v})
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].string < res[j].string
	})

	b := 0.
	for _, r := range res {
		b += r.float64
		fmt.Printf("\n %s\t%10.2f\t\t%10.2f", r.string, r.float64, b)
	}
}
