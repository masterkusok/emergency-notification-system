package loaders

import (
	"fmt"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"github.com/thedatashed/xlsxreader"
	"io"
	"strconv"
)

type XlsxParser struct {
}

/*
	xls file should store contact-data with schema:
	[name] [platform] [address]
*/

func (x XlsxParser) Parse(reader io.Reader) ([]entities.Contact, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	xlsxReader, err := xlsxreader.NewReader(bytes)
	if err != nil {
		return nil, err
	}

	result := make([]entities.Contact, 0)
	for row := range xlsxReader.ReadRows(xlsxReader.Sheets[0]) {
		if len(row.Cells) != 3 {
			return nil, fmt.Errorf("invalid xlsx format")
		}
		name := row.Cells[0].Value
		platform, err := strconv.Atoi(row.Cells[1].Value)
		if err != nil {
			return nil, fmt.Errorf("invalid xlsx format")
		}
		address := row.Cells[2].Value
		result = append(result, entities.Contact{Name: name, Address: address, Platform: platform})
	}
	return result, nil
}
