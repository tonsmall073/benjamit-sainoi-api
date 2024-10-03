FROM golang:1.23.1-alpine3.20

WORKDIR /app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go mod tidy

RUN swag init -o docs/v1

RUN go build -o myapp main.go

EXPOSE 8000

CMD echo "s" | ./myapp