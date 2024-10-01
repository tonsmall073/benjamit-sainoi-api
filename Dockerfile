FROM golang:1.23.1-alpine3.20

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8000

RUN go build -o myapp main.go

CMD echo "s" | ./myapp