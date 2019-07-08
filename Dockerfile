FROM golang:1.12-alpine3.10 as builder
ADD . /redirect
WORKDIR /redirect
RUN go build -o /tmp/redirect

FROM alpine:3.10
COPY --from=builder /tmp/redirect /usr/bin/redirect

ENTRYPOINT ["/usr/bin/redirect"]
