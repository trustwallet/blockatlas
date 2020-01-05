FROM golang:alpine as builder
ADD . /go/src/github.com/trustwallet/blockatlas
RUN apk add git \
 && go get -d -v github.com/trustwallet/blockatlas \
 && CGO_ENABLED=0 go install -a \
    -ldflags='-s -w -extldflags "-static"' \
    github.com/trustwallet/blockatlas

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/blockatlas /bin/blockatlas
COPY --from=builder /go/src/github.com/trustwallet/blockatlas/config.yml /config.yml
CMD ["/bin/blockatlas", "api"]
