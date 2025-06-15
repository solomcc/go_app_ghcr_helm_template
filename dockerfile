FROM golang:1.23-alpine as builder

WORKDIR /app
COPY ./cmd .
RUN go build -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
ENTRYPOINT ["./server"]