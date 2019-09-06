package parser

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// GNC is the main structure constructed
// from parsing a GnuCash file.
type GNC struct {
	Accounts     map[string]Account
	Transactions []Transaction
	dec          *xml.Decoder
}

// NewGNC constructor.
func NewGNC() *GNC {
	g := new(GNC)
	g.Accounts = make(map[string]Account)
	return g
}

// ParseFile reads (read-only) a Gnucash file
// and parse it with the provided parser.
func (gnc *GNC) ParseFile(path string) {

	// Open selected file
	gfile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer gfile.Close()

	// Get an unzipped Reader
	reader, err := gzip.NewReader(gfile)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	// Store the decoder ...
	gnc.dec = xml.NewDecoder(reader)

	// ... and start parsing
	gnc.parseBook()

}

func (gnc *GNC) parseBook() {
	for {
		tok, err := gnc.dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if tok == nil {
			fmt.Println("Nil Token")
			break
		}
		switch ty := tok.(type) {
		case xml.EndElement:
			if ty.Name.Local == "book" {
				return
			}
		case xml.StartElement:
			switch ty.Name.Local {
			case "account":
				a := gnc.parseAccount(ty)
				if a != nil {
					gnc.Accounts[a.GUID] = *a
				}
			case "transaction":
				a := gnc.parseTransaction(ty)
				if a != nil {
					gnc.Transactions = append(gnc.Transactions, *a)
				}
			default:
				// do nothing
			}
		default:
			// ignore the rest and continue parsing
		}
	}
}

// ParseAccount parse an account
// Auto parsing
// Force copy of parameter, because ty will be overwritten
func (gnc *GNC) parseAccount(ty xml.StartElement) *Account {
	a := new(Account)
	err := gnc.dec.DecodeElement(a, &ty)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return a
}

// ParseTransaction parse a Transaction
// Manual parsing
func (gnc *GNC) parseTransaction(ty xml.StartElement) *Transaction {
	t := NewTransaction()
	for {
		tok, err := gnc.dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if tok == nil {
			fmt.Println("Nil Token")
			break
		}
		switch ty := tok.(type) {
		case xml.StartElement:
			switch ty.Name.Local {
			case "splits":
				s := gnc.parseSplits(ty)
				if s != nil {
					// fmt.Printf("PARSED SPLITS : %#v\n", s)
					for _, ss := range s.Data {
						a := ss.Account
						cleanValue := strings.Replace(ss.Value, "/100", "", -1)
						v, e := strconv.Atoi(cleanValue)
						if e != nil {
							fmt.Println("Converting split value : ", e)
						}
						t.Splits[a] = float64(v) / 100.
					}
				}
			case "slots":
				s := gnc.parseSlots(ty)
				if s != nil {
					//fmt.Printf("PARSED SLOTS : %#v\n", s)
					for _, ss := range s.Data {
						key := ss.Key
						switch key {
						case "notes":
							t.Slots[key] = ss.Value.Text
						case "date-posted":
							t.Slots[key] = ss.Value.Date
						}
					}
				}
			case "date-entered":
				var s struct {
					Date string `xml:"date"`
				}
				gnc.dec.DecodeElement(&s, &ty)
				//fmt.Println("===>", s, "<====")
				t.DateEntered = cleanDateString(s.Date)

			case "date-posted":
				var s struct {
					Date string `xml:"date"`
				}
				gnc.dec.DecodeElement(&s, &ty)
				//fmt.Println("===>", s, "<====")
				t.DatePosted = cleanDateString(s.Date)

			case "id":
				if ty.Name.Space != "http://www.gnucash.org/XML/trn" {
					// There are other ids around ...
					//fmt.Println("Discarding : ", ty.Name.Space)
					break
				}
				var s string
				gnc.dec.DecodeElement(&s, &ty)
				//fmt.Println("TRN ID ===>", s, "<====")
				t.GUID = s

			}
		case xml.EndElement:
			if ty.Name.Local == "transaction" {
				return t
			}
		default:
		}
	}
	return nil
}

// PSplit structure
type PSplit struct {
	Value   string `xml:"value"`
	Account string `xml:"account"`
}

// PSplits array of Split
type PSplits struct {
	Data []PSplit `xml:"split"`
}

func (gnc *GNC) parseSplits(ty xml.StartElement) *PSplits {
	s := new(PSplits)
	err := gnc.dec.DecodeElement(s, &ty)
	if err != nil {
		print(err)
		return nil
	}
	return s
}

// PSlots array of Slot
type PSlots struct {
	Data []PSlot `xml:"slot"`
}

// PSlot structure
type PSlot struct {
	Key   string `xml:"key"`
	Value Value  `xml:"value"`
}

//Value structure
type Value struct {
	Text string `xml:",chardata"`
	Date string `xml:"gdate"`
}

func (gnc *GNC) parseSlots(ty xml.StartElement) *PSlots {
	s := new(PSlots)
	err := gnc.dec.DecodeElement(s, &ty)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return s
}

func cleanDateString(in string) (out string) {
	r, e := regexp.Compile(`^[0-9]{4}-[0-9]{1,2}-[0-9]{1,2}`)
	if e != nil {
		panic(e)
	}
	out = r.FindString(in)
	return out
}
