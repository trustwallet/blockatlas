FROM golang:latest AS builder

ARG SERVICE

RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build -o bin/blockatlas ./cmd/$SERVICE

FROM debian:latest
COPY --from=builder /build/bin /app/bin/$SERVICE
COPY --from=builder /build/config.yml /config/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app/bin/$SERVICE

ENTRYPOINT ["/app/bin/blockatlas", "-c", "/config/config.yml"]
