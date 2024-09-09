# dockerfile
FROM golang:1.22.5 as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main


FROM alpine:latest

RUN apk add --no-cache gcompat

WORKDIR /

COPY --from=builder /app/main .
COPY --from=builder /app/internal/assets ./internal/assets
COPY --from=builder /app/internal/templates ./internal/templates
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /

ENV ZONEINFO=/zoneinfo.zip

EXPOSE 8080

ENTRYPOINT ["./main"]