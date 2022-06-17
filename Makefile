BINARY = server
FILE = src/main.go

main: $(FILE)
	go mod download
	go build -o $(BINARY) $(FILE)
	bash ./scripts/env.sh

clear:
	sudo rm -rf ./uploads/**

test: $(FILE)
	go test