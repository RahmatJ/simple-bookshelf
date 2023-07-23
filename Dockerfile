FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

WORKDIR /app/cmd

RUN go build -o ../binary

WORKDIR /app

ENTRYPOINT ["/app/binary"]
