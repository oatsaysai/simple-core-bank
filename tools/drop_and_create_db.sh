#!/bin/sh

DB_NAME="go-http-server-template"
CONTAINER_NAME="go-http-server-template-postgres"

docker exec -i "$CONTAINER_NAME" psql -U postgres -c "DROP DATABASE IF EXISTS \"$DB_NAME\";" &&
docker exec -i "$CONTAINER_NAME" psql -U postgres -c "CREATE DATABASE \"$DB_NAME\";"