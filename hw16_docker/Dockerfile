# Stage 1: Build
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Копируем файлы с зависимостями
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

RUN ls -la

# Собираем приложение
RUN go build -o main .

# Stage 2: Run
FROM alpine:3.16

WORKDIR /app

# Устанавливаем необходимые пакеты
RUN apk --no-cache add ca-certificates

# Копируем собранный бинарный файл из предыдущего этапа
COPY --from=builder /app/main .

# Копируем SQL файлы и другие необходимые ресурсы
COPY --from=builder /app/sqlc/hw14_db_new.sql /app/sqlc/hw14_db_new.sql

# Открываем порт приложения
EXPOSE ${APP_PORT}