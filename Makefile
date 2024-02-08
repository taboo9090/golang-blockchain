run:
	@echo "Running..."
	@go run ./main.go

build:
	@echo "Building..."
	@go build -o ./bin/blockchain .

clear:
	@echo "Clearing..."
	@rm -f ./bin/blockchain