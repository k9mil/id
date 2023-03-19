FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server main.go

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/server .

RUN apk --no-cache add tzdata

EXPOSE 8080

CMD ["./server"]