FROM golang:1.19.2-buster

WORKDIR /src

RUN apt update && apt install -y protobuf-compiler=3.6.1.3-2

# TODO pin versions / put in dependency file
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

ENTRYPOINT [ "protoc" ]