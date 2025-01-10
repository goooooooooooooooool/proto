#!/bin/bash

set -e

PROTO_SRC=proto

# Генерация gRPC кода
protoc --go_out=. --go-grpc_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  *.proto

# Загрузка зависимостей в директории service
go mod tidy
