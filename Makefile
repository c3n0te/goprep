.PHONY: server
server:
	go build -race -ldflags "-s -w" -o ./bin ./server

.PHONY: cli
cli:
	go build -race -ldflags "-s -w" -o ./bin ./cli

.PHONY: thread
thread:
	go build -race -ldflags "-s -w" -o ./bin ./thread

.PHONY: scrape
scrape:
	go build -race -ldflags "-s -w" -o ./bin ./scrape

.PHONY: url
url:
	go build -race -ldflags "-s -w" -o ./bin ./url

.PHONY: cipher
cipher:
	go build -race -ldflags "-s -w" -o ./bin ./cipher
