FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/api-server ./cmd/main

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /bin .

EXPOSE 8080

CMD [ "./api-serv" ]
