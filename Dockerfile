FROM golang:latest AS builder

RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build -o bin/blockatlas .

FROM debian:latest
COPY --from=builder /build/bin /app
COPY --from=builder /build/config.yml /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app

ENTRYPOINT ["/app/blockatlas"]