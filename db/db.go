package db

import (
	"encoding/csv"
	"os"
	"time"
)

var (
	reader  *csv.Reader
	records [][]string
)

type Record struct {
	PhishID          int
	URL              string
	PhishDetailURL   string
	SubmissionTime   time.Time
	Verified         bool
	VerificationTime time.Time
	Online           bool
	Target           string
}

func LoadDatabase(database string) {
	db, _ := os.Open(database)
	reader = csv.NewReader(db)
	records, _ = reader.ReadAll()
}

func CheckURL(url string) (present bool, record []string) {
	for _, record := range records {
		if record[1] == url {
			return true, record
		}
	}
	return false, nil
}
