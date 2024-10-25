FROM golang:1.22.2

WORKDIR /home/dev/Projects/hl-server-go/

COPY . .
COPY cmd/.env ./cmd/.env


EXPOSE 8080

RUN go build -o main ./cmd/main.go

CMD ["./main"]