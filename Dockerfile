FROM golang:1.24.5 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /event-booking-api ./cmd/api

FROM alpine:latest
RUN adduser -D appuser
USER appuser
WORKDIR /root/
COPY --from=builder /event-booking-api .
EXPOSE 8080
CMD ["./event-booking-api"]