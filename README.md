Install the protocol compiler plugins for Go using the following commands:
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

Update your PATH so that the protoc compiler can find the plugins:
$ export PATH="$PATH:$(go env GOPATH)/bin"

 <!-- (無ければ)Protocol Buffers v3 をインストール -->

$ brew install protobuf ("$ which protoc" to check if you have protoc)

% go.mod ファイルを初期化(ここから検証してる)
$ go mod init github.com/Shinkomori19/intern

% Grpc のインストール
$ go get -u google.golang.org/grpc

% echo_grpc.pb.go と echo.pb.go を生成(@api dir で叩く)
$ protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
 --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
 echo.proto

% Build in the intern directory.
$ make

% Execute
Open a terminal, and use command below to start server.
$ ./server_main

Open another terminal, and use command below to send request to the server.
$ ./client_main --reqMessage="<TYPE A MESSAGE HERE>"

This sends a default message of "No specified message.", and server echoes it. To specify a message, do

$ ./client_main
