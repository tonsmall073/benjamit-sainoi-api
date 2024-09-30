FROM golang:1.23.1-alpine3.20

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8000

CMD ["go", "run","main.go"]