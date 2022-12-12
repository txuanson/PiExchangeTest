# PI EXCHANGE TEST

## How to run
Before this you need to prepare required files `customers.csv`, `email_template.json`

### Local

* Install required packages
```sh
go mod download
```

* Build the binary
```sh
go build -o main .
```

* Run the binary
```sh
# ./main ${EMAIL_TEMPLATE_PATH} ${CUSTOMER_PATH} ${OUTPUT_DIR} ${ERROR_PATH}
# Example
./main assets/email_template.json \
  assets/customers.csv \
  output error/errors.csv
```

### Docker

* Build the image
```sh
docker build -t go-mailer .
```

* Run the container
```sh
docker run --name go-mailer \
  --rm --privileged \
  -v /$(pwd)/assets:/assets:ro \
  -v /$(pwd)/error:/error \
  -v /$(pwd)/output:/output \
  go-mailer \
  assets/email_template.json \
  assets/customers.csv output \
  error/errors.csv
```

## How to test

```sh
go test ./...
```

Or you can run the test with VSCode Testing

