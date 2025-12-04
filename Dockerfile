FROM golang:1.24-alpine AS builder

RUN apk update && \
    apk add --no-cache bash make protobuf-dev ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /app

COPY . .

RUN make compile

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/bin/server /server

ENTRYPOINT ["/server"]
