FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go
RUN go build -o migrate ./cmd/migrate/main.go

EXPOSE 8080

CMD ["sh", "-c", "./migrate up && ./main"]