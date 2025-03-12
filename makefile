run: build run-without-build

start-db:
	./tools/start_postgresql_docker.sh
	sleep 5
	./tools/drop_and_create_db.sh

stop-db:
	./tools/stop_postgresql_docker.sh

build-docker-image:
	docker build -t go-http-server-template .
	./tools/remove_all_none_image.sh

build:
	go build -o go-http-server-template src/main.go

run-without-build:
	./go-http-server-template migrate-db
	./go-http-server-template serve-http-api
