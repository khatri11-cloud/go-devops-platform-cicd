# ---------- Stage 1: Build ----------
FROM golang:1.26.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# ---------- Stage 2: Runtime ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]