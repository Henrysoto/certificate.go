package cert

import (
	"encoding/csv"
	"os"
)

func ParseCSV(filename string) ([]*Cert, error) {
	certs := make([]*Cert, 0)
	f, err := os.Open(filename)
	if err != nil {
		return certs, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return certs, err
	}

	for _, record := range records {
		course := record[0]
		name := record[1]
		date := record[2]
		c, err := New(course, name, date)
		if err != nil {
			return certs, err
		}
		certs = append(certs, c)
	}

	return certs, nil
}
