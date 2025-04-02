# 18.2Mb
FROM golang:1.24.1 AS builder

WORKDIR /usr/local/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "-s -w -extldflags '-static'" \
    -o linker ./cmd/linker/main.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /usr/local/app/linker /bin/linker
COPY --from=builder /usr/local/app/config.yaml .
COPY --from=builder /usr/local/app/ui ./ui

ENTRYPOINT ["/bin/linker"]




# 982Mb

# # Используем базовый образ Go
# FROM golang:1.24.1

# # Устанавливаем рабочую директорию
# WORKDIR /app

# # Копируем файлы проекта
# COPY . .

# # Собираем приложение
# RUN go build -o linker ./cmd/linker/main.go

# # Указываем команду для запуска
# CMD ["./linker"]