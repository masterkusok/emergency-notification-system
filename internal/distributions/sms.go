package distributions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	apiURL = "https://api.exolve.ru/messaging/v1/SendSMS"
)

type smsContent struct {
	Number      string `json:"number"`
	Destination string `json:"destination"`
	Text        string `json:"text"`
}

type SMSDistributor struct {
	number string
	apiKey string
}

func NewSMSDistributor(apiKey, number string) (*SMSDistributor, error) {
	return &SMSDistributor{
		apiKey: apiKey,
		number: number,
	}, nil
}

func (s *SMSDistributor) Send(message string, contact entities.Contact) error {
	sms := smsContent{
		Number:      s.number,
		Destination: contact.Address,
		Text:        message,
	}

	jsonData, err := json.Marshal(sms)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bytesBody, _ := io.ReadAll(resp.Body)
		log.Printf("%s", bytesBody)
		return fmt.Errorf("error: response status: %d", resp.StatusCode)
	}

	return nil
}
