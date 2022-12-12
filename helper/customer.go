package helper

import (
	"go-mailer/types"
	"net/mail"
	"strings"
	"time"
)

func GenerateDate() string {
	return time.Now().Format("01 Jan 2006")
}

func CreateEmail(template types.EmailTemplate, customer types.Customer) types.Email {
	body := template.Body

	// strings.Replace the template variables with the customer data
	body = strings.Replace(body, "{{TITLE}}", customer.Title, -1)
	body = strings.Replace(body, "{{FIRST_NAME}}", customer.FirstName, -1)
	body = strings.Replace(body, "{{LAST_NAME}}", customer.LastName, -1)
	body = strings.Replace(body, "{{TODAY}}", GenerateDate(), -1)

	return types.Email{
		From:     template.From,
		To:       customer.Email,
		Subject:  template.Subject,
		Body:     body,
		MimeType: template.MimeType,
	}
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}