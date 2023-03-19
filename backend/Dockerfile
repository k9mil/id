FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server id.go

FROM alpine:3.15

WORKDIR /root

COPY --from=builder /app/server .

RUN apk --no-cache add tzdata

EXPOSE 8080

CMD ["./server"]