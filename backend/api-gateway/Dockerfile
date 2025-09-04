FROM golang:1.24.2

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o api-gateway ./cmd/main.go


EXPOSE 8000

CMD ["./api-gateway"]

