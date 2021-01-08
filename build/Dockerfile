FROM golang:alpine AS builder

WORKDIR $GOPATH/src/nosoxon.net
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo \
    -ldflags "-s -w" \
    -o /web-server ./cmd/web-server

FROM scratch
COPY --from=builder /web-server /web-server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY web/ /web/

CMD ["/web-server"]