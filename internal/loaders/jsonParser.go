package loaders

import (
	"encoding/json"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"io"
)

type JsonParser struct {
}

func (j JsonParser) Parse(reader io.Reader) ([]entities.Contact, error) {
	bytes := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		bytesRead, err := reader.Read(buffer)
		if err != nil || bytesRead == 0 {
			break
		}
		bytes = append(bytes, buffer[:bytesRead]...)
	}

	if len(bytes) == 0 {
		return make([]entities.Contact, 0), nil
	}

	result := make([]entities.Contact, 0)
	err := json.Unmarshal(bytes, &result)
	return result, err
}
