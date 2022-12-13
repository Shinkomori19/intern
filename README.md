# Echo server with Go and gRPC

There are two .go files that need to be run, in order to achieve echo server functionality. Those are /cmd/client/client_main.go and /cmd/server/server_main.go.<br>
Here are the steps for executing the code.

Install the protocol compiler plugins for Go using the following commands:
`$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
`$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

Update PATH so that the protoc compiler can find the plugins:
`$ export PATH="$PATH:$(go env GOPATH)/bin"`

Init go.mod file
`$ go mod init github.com/Shinkomori19/intern`

Install Grpc
`$ go get -u google.golang.org/grpc`

Automatically create echo_grpc.pb.go and echo.pb.go (Run this command in intern/api/)
`$ protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \ --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative echo.proto`

Build (Run this command in intern/)
`$ make`

Execute. Open a terminal, and use command below to start server.
`$ ./server_main`

Open another terminal, and use command below to send request to the server.
`$ ./client_main`

This sends a default message of "No specified message." to server. To specify a message, do
`$ ./client_main --reqMessage="<TYPE A MESSAGE HERE>"`

To clean up, run
`$ make clean`
