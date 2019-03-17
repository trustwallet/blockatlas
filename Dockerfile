FROM golang:alpine as builder
ADD . /go/src/trustwallet.com/blockatlas
RUN apk add git \
 && go get -d -v trustwallet.com/blockatlas \
 && CGO_ENABLED=0 go install -a \
    -installsuffix cgo \
    -ldflags="-s -w" \
    trustwallet.com/blockatlas

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/blockatlas /bin/
CMD ["/bin/blockatlas"]
