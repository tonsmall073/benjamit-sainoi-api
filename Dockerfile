FROM golang:1.23.1-alpine3.20

WORKDIR /app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go mod tidy

RUN swag init

RUN go build -o myapp main.go

EXPOSE 8000
EXPOSE 50051

CMD ["sh", "-c", "echo 'hg' | ./myapp"]