package helper

import (
	"errors"
	"go-mailer/types"
	"os"
)

// ParseArgs parses the command line arguments and returns a pointer to a types.Args struct
// Cannot test this function because it uses the os.Args variable
func ParseArgs() (*types.Args, error) {
	if(len(os.Args) != 5) {
		return nil, errors.New("invalid number of arguments, expected 4: templatePath, customerPath, outputPath, errorCustomerPath")
	}

	args := types.Args{
		TemplatePath:      os.Args[1],
		CustomerPath:      os.Args[2],
		OutputPath:        os.Args[3],
		ErrorCustomerPath: os.Args[4],
	}

	return &args, nil
}
