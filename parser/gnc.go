package parser

import (
	"encoding/xml"
)

// Account struct - automatic parse
type Account struct {
	XMLName     xml.Name `xml:"account"`
	Name        string   `xml:"name"`
	Type        string   `xml:"type"`
	Description string   `xml:"description"`
	GUID        string   `xml:"id"`
	ParentGUID  string   `xml:"parent"`
	//RawContent  string   `xml:",innerxml"` // Debug
}

// Transaction struct - manual parse, does not reflect xml structure
type Transaction struct {
	Splits      map[string]float64
	Slots       map[string]string
	DateEntered string
	DatePosted  string
}

// NewTransaction constructs a new Transaction and initialize maps
func NewTransaction() *Transaction {
	t := new(Transaction)
	t.Splits = make(map[string]float64)
	t.Slots = make(map[string]string)
	return t
}
