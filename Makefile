BINARY = server
FILE = src/main.go

main: $(FILE)
	go mod download
	go build -o $(BINARY) $(FILE)
	bash ./scripts/env.sh

test: $(FILE)
	bash ./scripts/test.sh

prod: $(FILE)
	go mod download
	go mod tidy
	go build -o $(BINARY) $(FILE)
