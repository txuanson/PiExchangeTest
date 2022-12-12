package helper

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"go-mailer/types"
	"os"
	"strings"

	"github.com/tebeka/atexit"
)

// Custom JSON marshal function to prevent escaping HTML
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func EnsureDirectoryExists(path string) {
	os.MkdirAll(path, 0644)
}

func EnsureFileExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		seperatedPath := strings.Split(path, "/")
		dir := strings.Join(seperatedPath[:len(seperatedPath)-1], "/")
		os.MkdirAll(dir, 0644)
	}
}

func ReadEMailTemplate(path string) (*types.EmailTemplate, error) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var template types.EmailTemplate

	json.Unmarshal(buffer, &template)

	return &template, nil
}

func CreateCustomerCSVReader(path string) (*csv.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// Reigster the file close function to be called when the program exits
	atexit.Register(func() {
		file.Close()
	})

	reader := csv.NewReader(file)

	// skip the header
	reader.Read()

	return reader, nil
}

func ReadCustomerData(reader *csv.Reader) (*types.Customer, error) {
	record, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var customer types.Customer

	customer.Title = record[0]
	customer.FirstName = record[1]
	customer.LastName = record[2]
	customer.Email = record[3]

	return &customer, nil
}

func WriteEmailToFile(path string, email types.Email) error {
	buffer, err := JSONMarshal(email)

	if err != nil {
		return err
	}

	err = os.WriteFile(path, buffer, 0644)

	if err != nil {
		return err
	}

	return nil
}

func CreateErrorCustomerCSVWriter(path string) (*csv.Writer, error) {
	file, err := os.Create(path)

	// Register the file close function to be called when the program exits gracefully
	atexit.Register(func() {
		file.Close()
	})

	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write the header
	title := []string{"TITLE", "FIRST_NAME", "LAST_NAME", "EMAIL"}
	writer.Write(title)

	return writer, nil
}

func AppendErrorCustomerToCSV(writer *csv.Writer, customer types.Customer) error {
	defer writer.Flush()
	record := []string{customer.Title, customer.FirstName, customer.LastName, customer.Email}

	err := writer.Write(record)
	if err != nil {
		return err
	}

	return nil
}
