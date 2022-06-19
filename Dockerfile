FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build

EXPOSE 8080

CMD ["/app/KubectlServer","serve"]
