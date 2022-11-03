FROM golang:1.19.2-buster as build

WORKDIR /service

RUN apt update && apt install -y protobuf-compiler=3.6.1.3-2

# TODO pin versions / put in dependency file
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

COPY src/services/project_service/src/proto/service.proto ./src/services/project_service/src/proto/service.proto
# Generate protobuf files
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./src/services/project_service/src/proto/service.proto

FROM golang:1.19.2-buster
ARG ENVIRONMENT

WORKDIR /service

COPY --from=build /service /service

COPY src/ent ./src/ent

WORKDIR /service/src/services/project_service

COPY src/services/project_service/go.mod ./
COPY src/services/project_service/go.sum ./

# Download dependencies
RUN go mod download

# Copy source files
COPY src/services/project_service/src/*.go ./

# build binary
RUN [ "${ENVIRONMENT}" != "prod" ] || go build -o /main

ENV ENVIRONMENT=${ENVIRONMENT}
COPY src/services/project_service/docker-entrypoint.sh /docker-entrypoint.sh
CMD [ "/docker-entrypoint.sh" ]