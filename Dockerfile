# ---------- BUILD STAGE ----------
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# ---------- RUNTIME STAGE ----------
FROM alpine:3.20
WORKDIR /app
RUN adduser -D appuser
COPY --from=builder /app/main .
COPY .env .
COPY data_example.json .
USER appuser
EXPOSE 3000
CMD ["./main"]