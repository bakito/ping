FROM golang:1.18 as builder

WORKDIR /build

RUN apt-get update && apt-get install -y upx
COPY . .

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN go build -a -installsuffix cgo -ldflags="-w -s" -o ping && \
    upx --ultra-brute -q ping

# application image

FROM scratch

LABEL maintainer="bakito <github@bakito.ch>"
USER 1001
ENTRYPOINT ["/go/bin/ping"]

COPY --from=builder /build/ping /go/bin/ping
