FROM golang:1.27-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY ./ ./
RUN go build -o /app/xm-web ./cmd/web/main.go


FROM alpine:3.24
COPY --from=builder /app/xm-web /app/xm-web
EXPOSE 8000
CMD ["/app/xm-web"]
