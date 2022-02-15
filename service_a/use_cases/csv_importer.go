package use_cases

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

type Response struct {
	Results chan map[string]interface{}
	Errors  chan error
}

func ReadCSV(filePath string) (*Response, error) {
	resp := &Response{
		Results: make(chan map[string]interface{}, 100),
		Errors:  make(chan error),
	}

	// Load a csv file.
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)

	go func() {
		keys := map[int]string{}
		var counter uint

	Reader:
		for {
			record, err := r.Read()
			if err == io.EOF {
				file.Close()
				break Reader
			}
			if len(record) == 0 {
				file.Close()
				break Reader
			}
			if counter == 0 {
				for idx, key := range record {
					keys[idx] = key
				}
			}
			counter++
			if counter == 1 {
				continue
			}
			output := map[string]interface{}{}
			for idx, value := range record {
				// Trim added to handle with strings inconsistency
				key := strings.Trim(keys[idx], " ")
				output[key] = strings.Trim(value, " ")
			}
			resp.Results <- output
		}
		close(resp.Results)
	}()
	return resp, nil

}
