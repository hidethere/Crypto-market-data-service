# Use official Go image
FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o /usr/local/bin/app ./cmd/server/main.go

EXPOSE 8080

CMD ["app"]
