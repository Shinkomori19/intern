FROM golang:1.19

RUN mkdir /build
ADD . /build
WORKDIR /build
RUN cd /build/

RUN apt-get update && apt install -y protobuf-compiler

RUN GO111MODULE=on
RUN go get -u -v google.golang.org/grpc

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN export PATH="$PATH:$(go env GOPATH)/bin"
RUN go get -u google.golang.org/grpc
RUN cd api && protoc --go_out=../pkg/grpc --go_opt=paths=source_relative --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative echo.proto
RUN cd .. & make

CMD [ "/client_main" ]