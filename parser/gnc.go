package parser

import "encoding/xml"

// Account struct
type Account struct {
	XMLName     xml.Name `xml:"account"`
	Name        string   `xml:"name"`
	Type        string   `xml:"type"`
	Description string   `xml:"description"`
	GUID        string   `xml:"id"`
	ParentGUID  string   `xml:"parent"`
	RawContent  string   `xml:",innerxml"` // Debug
}

//Accounts are a set of Account
type Accounts []Account

// Transaction struct
type Transaction struct {
	XMLName    xml.Name `xml:"transaction"`
	RawContent string   `xml:",innerxml"` // Debug

}

// Transactions are a set of Transaction
type Transactions []Transaction

// Book struct
type Book struct {
	XMLName      xml.Name     `xml:"book"`
	Accounts     Accounts     `xml:"account"`
	Transactions Transactions `xml:"transation"`
}

// Gnc is the top level struct
type Gnc struct {
	XMLName xml.Name `xml:"gnc-v2"`
	Book    Book     `xml:"book"`
}
