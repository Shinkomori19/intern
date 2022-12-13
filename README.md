# Echo server with Go and gRPC

There are two .go files that need to be run, in order to achieve echo server functionality.<br>
Those files are /cmd/client/client_main.go and /cmd/server/server_main.go.<br>

Here are the steps for executing the code.

### 1. Init go.mod file:<br>

`$ go mod init github.com/Shinkomori19/intern`

### 2. Install the protocol compiler plugins for Go using the following commands:<br>

`$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`<br>
`$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`<br>

### 3. Update PATH so that the protoc compiler can find the plugins:<br>

`$ export PATH="$PATH:$(go env GOPATH)/bin"`

### 4. Install Grpc:<br>

`$ go get -u google.golang.org/grpc`

### 5. Automatically create echo_grpc.pb.go and echo.pb.go (Run this command in intern/api/):

`$ protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \ --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative echo.proto`

### 6. Build (Run this command in intern/)<br>

`$ make`

### 7. To start a server, open a terminal and:<br>

`$ ./server_main`

### 8. To send a request to the server, open another terminal and:<br>

`$ ./client_main`
This sends a default message of "No specified message." to server.<br>

### 9. To specify a message, do:<br>

`$ ./client_main --reqMessage="<TYPE A MESSAGE HERE>"`

### To clean executables, run:<br>

`$ make clean`
