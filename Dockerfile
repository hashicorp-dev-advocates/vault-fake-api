FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download

COPY . .

RUN go mod download cloud.google.com/go/kms

RUN go build -o fakevault

EXPOSE 8200

CMD ["./fakevault"]
