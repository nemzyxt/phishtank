package db

import (
	"encoding/csv"
	"os"
)

var (
	reader  *csv.Reader
	records [][]string
)

// load the offline database (csv file)
func LoadDatabase(database string) error {
	db, err := os.Open(database)
	if err != nil {
		return err
	}

	reader = csv.NewReader(db)

	records, err = reader.ReadAll()
	if err != nil {
		return err
	}

	return nil
}

// check whether a url exists in the database
func CheckURL(url string) (record []string) {
	for _, record := range records {
		if record[1] == url {
			return record
		}
	}

	return nil
}
