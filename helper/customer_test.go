package helper

import (
	"go-mailer/types"
	"reflect"
	"testing"
)

func TestGenerateDate(t *testing.T) {
	got := GenerateDate()

	// must update this value to the current date
	want := "12 Dec 2022"

	if got != want {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestCreateEmail(t *testing.T) {
	template := types.EmailTemplate{
		From:     "The Marketing Team<marketing@example.com>",
		Subject:  "A new product is being launched soon...",
		MimeType: "text/plain",
		Body:     "Hi {{TITLE}} {{FIRST_NAME}} {{LAST_NAME}},\nToday, {{TODAY}}, we would like to tell you that... Sincerely,\nThe Marketing Team",
	}

	customer := types.Customer{
		Title:     "Mr",
		FirstName: "John",
		LastName:  "Smith",
		Email:     "john.smith@example.com",
	}

	got := CreateEmail(template, customer)
	want := types.Email{
		From:     "The Marketing Team<marketing@example.com>",
		To:       "john.smith@example.com",
		Subject:  "A new product is being launched soon...",
		MimeType: "text/plain",
		Body:     "Hi Mr John Smith,\nToday, 12 Dec 2022, we would like to tell you that... Sincerely,\nThe Marketing Team",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestIsValidEmail(t *testing.T) {
	email := "a@a.com"
	got := IsValidEmail(email)
	want := true

	if got != want {
		t.Errorf("Test case: %v -> got = %v, want %v", email, got, want)
	}

	email = "a"
	got = IsValidEmail(email)
	want = false

	if got != want {
		t.Errorf("Test case: %v -> got = %v, want %v", email, got, want)
	}

	email = "a@a"
	got = IsValidEmail(email)
	want = true

	if got != want {
		t.Errorf("Test case: %v -> got = %v, want %v", email, got, want)
	}

	email = "a@a."
	got = IsValidEmail(email)
	want = false

	if got != want {
		t.Errorf("Test case: %v -> got = %v, want %v", email, got, want)
	}
}
