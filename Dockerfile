FROM golang:1.22

WORKDIR /dn/ondatra
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
#COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN apt-get -y update && apt-get install -y protobuf-compiler && apt-get install -y golang-goprotobuf-dev

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest\
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 \
    #&& go get google.golang.org/grpc/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go@latest \
    && go install golang.org/x/tools/cmd/goimports@latest

