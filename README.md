# Go-Jek battleship

## Running the project

### With Docker

Put the input data in the root of the repository and call the file `input.txt`.
Then simply run:

	make run

### Without Docker

Run

	go build -o battleship cmd/main.go
	./battleship -input input.txt

`input.txt` would be the name of you're input file.

## Running tests

### With Docker

	make test

### Without Docker

	go test -cover ./...
