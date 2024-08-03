package loaders

import (
	"encoding/csv"
	"fmt"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"io"
	"strconv"
)

type CsvParser struct {
}

/*
	csv file should store contact-data with schema:
	[name],[platform],[address]
*/

func (c CsvParser) Parse(reader io.Reader) ([]entities.Contact, error) {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = 3
	result := make([]entities.Contact, 0)
	for {
		record, err := csvReader.Read()
		if err != nil || record == nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		platform, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, fmt.Errorf("%s is not an integer\n", record[1])
		}
		result = append(result, entities.Contact{Name: record[0], Address: record[2], Platform: platform})
	}
	return result, nil
}
