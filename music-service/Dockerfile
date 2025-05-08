# music-service/Dockerfile

FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o music-service

EXPOSE 8082

CMD ["./music-service"]
