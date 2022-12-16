# Creates two binary files.
all: server_main client_main

server_main: cmd/server/server_main.go
	go build cmd/server/server_main.go

client_main: cmd/client/client_main.go
	go build cmd/client/client_main.go

fmt:
	gofmt -w -s *.go

# Delete two binary files.
clean:
	rm server_main client_main