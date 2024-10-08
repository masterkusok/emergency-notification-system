package loaders

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/masterkusok/emergency-notification-system/internal/entities"
	"github.com/xuri/excelize/v2"
	"io"
	"os"
	"strconv"
	"testing"
)

func TestCreateContactLoader(t *testing.T) {
	loader := CreateContactLoader()
	if len(loader.parsers) != 4 {
		t.Fail()
	}
}

func TestCsvParser_Parse(t *testing.T) {
	expectedContacts := []entities.Contact{
		entities.Contact{Name: "contact_1", Platform: entities.TG, Address: "@contact_1"},
		entities.Contact{Name: "contact_2", Platform: entities.EMAIL, Address: "@contact_2"},
		entities.Contact{Name: "contact_3", Platform: entities.SMS, Address: "@contact_3"},
	}
	csvString := ""
	for _, contact := range expectedContacts {
		csvString += fmt.Sprintf("%s,%d,%s\n", contact.Name, contact.Platform, contact.Address)
	}
	os.WriteFile("temp.csv", []byte(csvString), 0644)
	parser := CsvParser{}
	file, _ := os.Open("temp.csv")
	resultContacts, _ := parser.Parse(file)

	if len(resultContacts) != len(expectedContacts) {
		t.Logf("wrong number of contacts:\nexpected: 4\ngot: %d", len(resultContacts))
		t.Fail()
	}

	for i := 0; i < len(resultContacts); i++ {
		if resultContacts[i].Name != expectedContacts[i].Name {
			t.Logf("name differs at index %d:\nexpected: %s\ngot: %s",
				i, expectedContacts[i].Name, resultContacts[i].Name)
			t.Fail()
		}
		if resultContacts[i].Platform != expectedContacts[i].Platform {
			t.Logf("platform differs at index %d:\nexpected: %d\ngot: %d",
				i, expectedContacts[i].Platform, resultContacts[i].Platform)
			t.Fail()
		}
		if resultContacts[i].Name != expectedContacts[i].Name {
			t.Logf("address differs at index %d:\nexpected: %s\ngot: %s",
				i, expectedContacts[i].Address, resultContacts[i].Address)
			t.Fail()
		}
	}
	os.Remove("temp.csv")
}

func TestCsvParser_Parse2(t *testing.T) {
	os.WriteFile("temp.csv", []byte("invalid csv,,,,,,,"), 0644)
	parser := CsvParser{}
	file, _ := os.Open("temp.csv")
	resultContacts, err := parser.Parse(file)

	if resultContacts != nil || !errors.Is(err, csv.ErrFieldCount) {
		t.Fail()
	}
	os.Remove("temp.csv")
}

func TestJsonParser_Parse(t *testing.T) {
	expectedContacts := []entities.Contact{
		entities.Contact{Name: "contact_1", Platform: entities.TG, Address: "@contact_1"},
		entities.Contact{Name: "contact_2", Platform: entities.EMAIL, Address: "@contact_2"},
		entities.Contact{Name: "contact_3", Platform: entities.SMS, Address: "@contact_3"},
	}
	jsonEncoded, _ := json.Marshal(expectedContacts)
	os.WriteFile("temp.json", jsonEncoded, 0644)

	parser := JsonParser{}
	file, _ := os.Open("temp.json")
	resultContacts, _ := parser.Parse(file)

	if len(resultContacts) != len(expectedContacts) {
		t.Logf("wrong number of contacts:\nexpected: 4\ngot: %d", len(resultContacts))
		t.Fail()
	}

	for i := 0; i < len(resultContacts); i++ {
		if resultContacts[i].Name != expectedContacts[i].Name {
			t.Logf("name differs at index %d:\nexpected: %s\ngot: %s",
				i, expectedContacts[i].Name, resultContacts[i].Name)
			t.Fail()
		}
		if resultContacts[i].Platform != expectedContacts[i].Platform {
			t.Logf("platform differs at index %d:\nexpected: %d\ngot: %d",
				i, expectedContacts[i].Platform, resultContacts[i].Platform)
			t.Fail()
		}
		if resultContacts[i].Name != expectedContacts[i].Name {
			t.Logf("address differs at index %d:\nexpected: %s\ngot: %s",
				i, expectedContacts[i].Address, resultContacts[i].Address)
			t.Fail()
		}
	}
	os.Remove("temp.json")
}

func TestJsonParser_Parse2(t *testing.T) {
	os.WriteFile("temp.json", []byte("invalid json!.!@#214SD"), 0644)

	parser := JsonParser{}
	file, _ := os.Open("temp.json")
	_, err := parser.Parse(file)

	if err == nil {
		t.Fail()
	}

	os.Remove("temp.json")
}

func TestXlsxParser_Parse(t *testing.T) {
	expectedContacts := []entities.Contact{
		entities.Contact{Name: "contact_1", Platform: entities.TG, Address: "@contact_1"},
		entities.Contact{Name: "contact_2", Platform: entities.EMAIL, Address: "@contact_2"},
		entities.Contact{Name: "contact_3", Platform: entities.SMS, Address: "@contact_3"},
	}
	xlsFile := excelize.NewFile()
	for i := 0; i < len(expectedContacts); i++ {
		xlsFile.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+1), expectedContacts[i].Name)
		xlsFile.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+1), strconv.Itoa(expectedContacts[i].Platform))
		xlsFile.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+1), expectedContacts[i].Address)
	}
	err := xlsFile.SaveAs("temp.xlsx")
	fmt.Println(err)
	parser := XlsxParser{}
	file, _ := os.Open("temp.xlsx")
	resultContacts, _ := parser.Parse(file)

	if len(resultContacts) != len(expectedContacts) {
		t.Logf("wrong number of contacts:\nexpected: 4\ngot: %d", len(resultContacts))
		t.Fail()
	}

	for i := 0; i < len(resultContacts); i++ {
		if resultContacts[i].Name != expectedContacts[i].Name {
			t.Logf("name differs at index %d:\nexpected: %s\ngot: %s",
				i, expectedContacts[i].Name, resultContacts[i].Name)
			t.Fail()
		}
		if resultContacts[i].Platform != expectedContacts[i].Platform {
			t.Logf("platform differs at index %d:\nexpected: %d\ngot: %d",
				i, expectedContacts[i].Platform, resultContacts[i].Platform)
			t.Fail()
		}
		if resultContacts[i].Name != expectedContacts[i].Name {
			t.Logf("address differs at index %d:\nexpected: %s\ngot: %s",
				i, expectedContacts[i].Address, resultContacts[i].Address)
			t.Fail()
		}
	}
	os.Remove("temp.xlsx")
}

type mockLoader struct {
	platform int
	counter  *int
}

func (m mockLoader) Parse(reader io.Reader) ([]entities.Contact, error) {
	*m.counter++
	return make([]entities.Contact, m.platform), nil
}

func TestContactLoader_ParseFrom(t *testing.T) {
	loader := CreateContactLoader()
	counter := new(int)
	loader.parsers[entities.TG] = mockLoader{platform: entities.TG, counter: counter}
	loader.parsers[entities.SMS] = mockLoader{platform: entities.SMS, counter: counter}
	loader.parsers[entities.EMAIL] = mockLoader{platform: entities.EMAIL, counter: counter}

	tgRes, err := loader.ParseFrom(nil, entities.TG)
	if err != nil || len(tgRes) != entities.TG {
		t.Fail()
	}

	smsRes, err := loader.ParseFrom(nil, entities.SMS)
	if err != nil || len(smsRes) != entities.SMS {
		t.Fail()
	}

	emailRes, err := loader.ParseFrom(nil, entities.EMAIL)
	if err != nil || len(emailRes) != entities.EMAIL {
		t.Fail()
	}
}
