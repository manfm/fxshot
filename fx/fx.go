package fx

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func Import(path string) []Record {
	f, err := os.Open(path)
	if err != nil {
		return []Record{}
	}
	defer f.Close()

	r := csv.NewReader(f)

	records := []Record{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		r := Record{}
		r.Date = record[0]
		r.Time = record[1]
		r.Open, _ = strconv.ParseFloat(record[2], 64)
		r.High, _ = strconv.ParseFloat(record[3], 64)
		r.Low, _ = strconv.ParseFloat(record[4], 64)
		r.Close, _ = strconv.ParseFloat(record[5], 64)
		r.Volume, _ = strconv.ParseUint(record[6], 0, 16)

		records = append(records, r)
	}

	return records
}
