export GOPATH=$(shell go env GOPATH)

build-generator:
	docker build -t task_wizard/generator:latest -f src/generator/Dockerfile src/generator/Dockerfile

generate:
	# docker run -it -v $(PWD):/src -w /src task_wizard/generator:latest --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $${SERVICE_PATH}/proto/service.proto
	docker run -it -v ${PWD}:/task_wizard -w /task_wizard/src golang:latest go generate ./ent

up $(ENVIRONMENT):
	@if [[ "$(ENVIRONMENT)" = "prod" ]]; then \
		docker compose up --build -d; \
	else \
		docker compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build -d; \
	fi

logs:
	docker compose logs