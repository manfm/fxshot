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
		r.Open = priceToFl32(record[2])
		r.High = priceToFl32(record[3])
		r.Low = priceToFl32(record[4])
		r.Close = priceToFl32(record[5])
		r.Volume, _ = strconv.ParseUint(record[6], 0, 16)

		records = append(records, r)
	}

	return records
}

func priceToFl32(price string) float32 {
	p, _ := strconv.ParseFloat(price, 32)
	return float32(p)
}
