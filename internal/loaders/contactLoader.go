// Package loaders provides functions to parse contact data from different types of files
package loaders

import (
	"fmt"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"io"
)

const (
	JSON = iota
	CSV
	XLSX
)

type parser interface {
	Parse(reader io.Reader) ([]entities.Contact, error)
}

type ContactLoader struct {
	parsers []parser
}

func CreateContactLoader() *ContactLoader {
	cl := ContactLoader{}
	cl.parsers = make([]parser, 4)
	cl.parsers[JSON] = JsonParser{}
	cl.parsers[CSV] = CsvParser{}
	cl.parsers[XLSX] = XlsxParser{}
	return &cl
}

// ParseFrom godoc
// This method is used to parse contact data from file of specific format
func (c *ContactLoader) ParseFrom(reader io.Reader, format int) ([]entities.Contact, error) {
	if format < 0 || format >= len(c.parsers) {
		return nil, fmt.Errorf("wrong parse format")
	}
	parser := c.parsers[format]
	return parser.Parse(reader)
}
