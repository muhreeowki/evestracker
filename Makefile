build:
	@go build -o bin/evestracker

run: build
	@./bin/evestracker
