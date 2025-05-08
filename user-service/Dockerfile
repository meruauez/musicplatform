FROM golang:1.24

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN go build -o user-service

CMD ["./user-service"]
