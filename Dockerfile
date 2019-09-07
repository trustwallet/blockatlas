FROM golang:alpine as builder
ADD . /go/src/github.com/trustwallet/blockatlas
RUN apk add git \
 && go get -d -v github.com/trustwallet/blockatlas/cmd \
 && CGO_ENABLED=0 go install -a \
    -installsuffix cgo \
    -ldflags="-s -w" \
    github.com/trustwallet/blockatlas/cmd

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/cmd /bin/blockatlas
COPY --from=builder /go/src/github.com/trustwallet/blockatlas/coins.yml /coins.yml
COPY --from=builder /go/src/github.com/trustwallet/blockatlas/config.yml /config.yml
CMD ["/bin/blockatlas", "api"]
