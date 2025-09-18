FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/server .

ENV PORT=8080
ENV JWT_SECRET_KEY=supersecret

EXPOSE 8080
CMD ["./server"]
