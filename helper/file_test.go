package helper

import (
	"bytes"
	"go-mailer/types"
	"reflect"
	"testing"
)

func TestReadEMailTemplate(t *testing.T) {
	got, err := ReadEMailTemplate("../assets/email_template.json")

	TestErrorHandler(t, err)

	want := types.EmailTemplate{
		From:     "The Marketing Team<marketing@example.com>",
		Subject:  "A new product is being launched soon...",
		MimeType: "text/plain",
		Body:     "Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}},\nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
	}

	if *got != want {
		t.Errorf("got = %v, want %v", *got, want)
	}
}

func TestCreateCustomerCSVReader(t *testing.T) {
	_, err := CreateCustomerCSVReader("../assets/customers.csv")

	TestErrorHandler(t, err)
}

func TestReadCustomerData(t *testing.T) {
	reader, err := CreateCustomerCSVReader("../assets/customers.csv")

	TestErrorHandler(t, err)

	got := make([]types.Customer, 0)

	for {
		customerData, err := ReadCustomerData(reader)
		if err != nil {
			break
		}
		got = append(got, *customerData)
	}

	want := []types.Customer{
		{
			Title:     "Mr",
			FirstName: "John",
			LastName:  "Smith",
			Email:     "john.smith@example.com",
		},
		{
			Title:     "Mrs",
			FirstName: "Michelle",
			LastName:  "Smith",
			Email:     "michelle.smith@example.com",
		},
	}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestWriteEmailToFile(t *testing.T) {
	email := types.Email{
		From:     "The Marketing Team<marketing@example.com>",
		To:       "john.smith@example.com",
		Subject:  "A new product is being launched soon...",
		MimeType: "text/plain",
		Body:     "Hi Mr John Smith,\nToday, 31 Dec 2020, we would like to tell you that... Sincerely,\nThe Marketing Team",
	}

	got, err := JSONMarshal(email)

	TestErrorHandler(t, err)

	err = WriteEmailToFile("../output/test_email_output.json", email)
	TestErrorHandler(t, err)

	// Do NOT delete the last new line! It is required for the test to pass
	want := []byte(`{"from":"The Marketing Team<marketing@example.com>","to":"john.smith@example.com","subject":"A new product is being launched soon...","mimeType":"text/plain","body":"Hi Mr John Smith,\nToday, 31 Dec 2020, we would like to tell you that... Sincerely,\nThe Marketing Team"}
`)

	if bytes.Equal(got, want) == false {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestAppendErrorCustomerToCSV(t *testing.T) {
	const errorFilePath = "../error/errors.csv"
	EnsureFileExists(errorFilePath)

	writer, err := CreateErrorCustomerCSVWriter(errorFilePath)

	TestErrorHandler(t, err)

	customer := types.Customer{
		Title:     "Mr",
		FirstName: "John",
		LastName:  "Smith",
		Email:     "john.s",
	}

	err = AppendErrorCustomerToCSV(writer, customer)

	TestErrorHandler(t, err)
}
