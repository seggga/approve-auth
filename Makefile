USERNAME := seggga
APP_NAME := auth-srv
VERSION := 1.0.1

check_fs:
	docker build -t docker.io/$(USERNAME)/$(APP_NAME):$(VERSION) . && \
	docker create --name $(APP_NAME) $(USERNAME)/$(APP_NAME):$(VERSION) && \
	docker export $(APP_NAME) | tar t > $(APP_NAME)_filesystem.txt

run_container:
	docker run -ti docker.io/$(USERNAME)/$(APP_NAME):$(VERSION) sh

run_app:
	AUTH_PORT_3000_TCP_PORT=3000 AUTH_PORT_4000_TCP_PORT=4000 go run ./cmd/server/main.go -c ./configs/config.yaml


gen_proto:
	mkdir -p pkg/proto && \
	protoc  proto/*.proto --go-grpc_out=pkg --go_out=pkg

build_cpu_profile:
	go tool pprof -svg http://172.31.193.58:3000/debug/pprof/profile\?seconds\=5 > ./pprof/pprf.svg

build_memory_profile:
	go tool pprof http://172.31.193.58:3000/debug/pprof/heap