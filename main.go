package main

import (
	"encoding/csv"
	"fmt"
	"go-mailer/helper"
	"go-mailer/types"
	"strings"
	"sync"

	"github.com/tebeka/atexit"
)

func main() {
	// Defer the atexit.Exit(0) function to ensure that the program exits cleanly
	defer atexit.Exit(0)

	args, err := helper.ParseArgs()
	helper.HandleError(err)

	// Ensure the output and error directory exists
	helper.EnsureDirectoryExists(args.OutputPath)
	helper.EnsureFileExists(args.ErrorCustomerPath)

	// Read the template file
	template, err := helper.ReadEMailTemplate(args.TemplatePath)
	helper.HandleError(err)

	// Create customer csv reader
	customerCsvReader, err := helper.CreateCustomerCSVReader(args.CustomerPath)
	helper.HandleError(err)

	err = nil

	errorCsvWriter, err := helper.CreateErrorCustomerCSVWriter(args.ErrorCustomerPath)
	helper.HandleError(err)

	errorChan := make(chan error)

	// Create a wait group to wait for all goroutines to finish
	wg := sync.WaitGroup{}
	for {
		customerData, err := helper.ReadCustomerData(customerCsvReader)
		if err != nil {
			break
		}

		// Increment the wait group counter
		wg.Add(1)

		// Create a closure goroutine to send email to customer
		go func(template *types.EmailTemplate, customerData *types.Customer, errorCsvWriter *csv.Writer, errorChan chan error, wg *sync.WaitGroup) {
			defer wg.Done()
			if !helper.IsValidEmail(customerData.Email) {
				err = helper.AppendErrorCustomerToCSV(errorCsvWriter, *customerData)
				if err != nil {
					errorChan <- err
				}
				return
			}

			email := helper.CreateEmail(*template, *customerData)

			emailFileName := fmt.Sprintf("email_%s_%s_%s.json", strings.ToLower(customerData.Title), strings.ToLower(customerData.FirstName), strings.ToLower(customerData.LastName))

			err = helper.WriteEmailToFile(args.OutputPath+"/"+emailFileName, email)
			if err != nil {
				errorChan <- err
			}
		}(template, customerData, errorCsvWriter, errorChan, &wg)
	}

	go func(wg *sync.WaitGroup, errorChan chan error) {
		// Wait for all goroutines to finish
		wg.Wait()

		// Close the error channel
		close(errorChan)
	}(&wg, errorChan)

	doneChan := make(chan bool)

	// Check if any error occurred
	go func(errorChan chan error, doneChan chan bool) {
		for err := range errorChan {
			helper.HandleError(err)
		}
		doneChan <- true
	}(errorChan, doneChan)

	// Wait for the done channel to be closed
	<-doneChan
}
