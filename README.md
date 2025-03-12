# go-http-server-template

## Prerequisites

- Make
- Docker 24 or later
- Go 1.24.1 or later
- PostgreSQL 16 or later

## How to run with docker-compose

```sh
docker-compose up -d
```

## How to rebuild image after edit code

```sh
docker-compose up -d --build
```

## Getting started

1. Start PostgreSQL docker

```sh
make start-db
```

2. Serve HTTP API

```sh
make run
```

3. Stop PostgreSQL docker

```sh
make stop-db
```

## Load test

1. Please set log level to "error" before load test

2. Load test create account API

```sh
./tools/load_test_create_account.sh 
```