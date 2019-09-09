package parser

// Account struct - automatic parse.
type Account struct {
	//XMLName     xml.Name `xml:"account"`
	Name        string   `xml:"name"`
	Type        string   `xml:"type"`
	Description string   `xml:"description"`
	GUID        string   `xml:"id"`
	Parent      string   `xml:"parent"`
	Child       []string // all direct child  accounts GUID
	//RawContent  string   `xml:",innerxml"` // Debug
}

// Transaction struct - manual parse, does not reflect xml structure.
type Transaction struct {
	GUID        string
	Splits      map[string]int // acout => value in cents
	Slots       map[string]string
	DateEntered string
	DatePosted  string
}

// NewTransaction constructs a new Transaction and initialize maps
func NewTransaction() *Transaction {
	t := new(Transaction)
	t.Splits = make(map[string]int) // Internally, storing cents to avoid rounding errors.
	t.Slots = make(map[string]string)
	return t
}
