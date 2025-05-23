FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/app/main.go

FROM alpine:latest

# Установка необходимых пакетов для работы
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копирование бинарного файла и конфигурационного файла из этапа сборки, а также миграций
COPY --from=builder /app/app .
COPY --from=builder /app/config/config.yml ./config/config.yml
COPY --from=builder /app/migration ./migration

CMD ["./app"]
