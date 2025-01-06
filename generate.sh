#!/bin/bash

set -e

PROTO_SRC=proto
PROTO_DEST=service

# Создание директории назначения
mkdir -p $PROTO_DEST

# Переход в директорию назначения
cd $PROTO_DEST

# Инициализация Go модуля
go mod init proto

# Возврат в корневую директорию
cd ..

# Генерация gRPC кода
protoc --go_out=$PROTO_DEST --go_opt=paths=source_relative \
       --go-grpc_out=$PROTO_DEST --go-grpc_opt=paths=source_relative $PROTO_SRC/*.proto

# Загрузка зависимостей в директории service
cd $PROTO_DEST
go mod tidy

# Возврат в корневую директорию после выполнения всех команд
cd ..