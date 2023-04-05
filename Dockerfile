FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env .

RUN go build -o safekeep .

EXPOSE 3033
CMD ["./safekeep"]