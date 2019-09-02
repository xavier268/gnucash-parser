package parser

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

// ReadFile reads (read-only) a Gnucash file
// and parse it into a Gnc structure.
func ReadFile(path string) (*Gnc, error) {

	// Open selected file
	gfile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer gfile.Close()

	// Get an unzipped Reader
	reader, err := gzip.NewReader(gfile)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// Unmarshal into xml
	gnc := new(Gnc)
	err = xml.NewDecoder(reader).Decode(gnc)
	if err != nil {
		return nil, err
	}
	return gnc, nil
}
