FROM golang:1.23

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN swag init --parseDependency --parseInternal

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
