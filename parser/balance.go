package parser

// Balance compute balance for specified account
func (gnc *GNC) Balance(actGUID string) float64 {
	b := 0

	for _, t := range gnc.Transactions {
		v, ok := t.Splits[actGUID]
		if ok {
			b += v
		}
	}
	// From cents to EUR...
	return float64(b) / 100.

}

// AccountName get account name from guid
func (gnc *GNC) AccountName(actGUID string) string {
	n, ok := gnc.Accounts[actGUID]
	if !ok {
		return "Unknown account ?! (guid : " + actGUID + ")"
	}
	return n.Name + "(" + n.Description + ")"
}
