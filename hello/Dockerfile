FROM golang:1.13.2-alpine3.10 as builder

WORKDIR /app
COPY . .
RUN go build main.go

FROM alpine:3.10
WORKDIR /app
COPY --from=builder /app/main /app
CMD ["./main"]