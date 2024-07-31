
FROM golang:1.22.4 AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app1 ./cmd/main/


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app2 ./cmd/worker/

# Используем минимальный образ для запуска приложений
FROM alpine:latest

WORKDIR /root/

# Копируем оба собранных приложения и конфигурационный файл
COPY --from=builder /app/bin/app1 ./app1
COPY --from=builder /app/bin/app2 ./app2
COPY --from=builder /app/config.yml ./config.yml

# Запускаем оба приложения
CMD ["sh", "-c", "./app1 & ./app2"]
