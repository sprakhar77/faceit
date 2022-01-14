FROM golang:1.14-alpine

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o ./application ./cmd/http/main.go

EXPOSE 8080:8080

ENTRYPOINT ./application
