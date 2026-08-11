.PHONY: server
server:
	go build -race -ldflags "-s -w" -o ./bin ./server

.PHONY: cli
cli:
	go build -race -ldflags "-s -w" -o ./bin ./cli
