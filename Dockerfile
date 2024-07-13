ARG GO_VERSION=1.22

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./internal ./internal

WORKDIR /app/internal

ENV CGO_ENABLED=0
ENV GO_OSARCH="linux/amd64"
RUN go build -o ./app .
RUN go build -o /app/simple-go-blog-webserver

FROM gcr.io/distroless/base:latest

COPY --from=builder /app/simple-go-blog-webserver .

EXPOSE 8080

CMD ["./simple-go-blog-webserver"]
