FROM docker.io/golang:alpine AS builder

# RUN apk update && apk add --no-cache git
RUN apk update

WORKDIR /tmp/jp

COPY src/ src/
RUN go version
RUN go mod init src
RUN go mod download
# RUN go get github.com/eclipse/paho.mqtt.golang

RUN go build -o cattled src/cattled.go

#@ step 2: build a minimal image with our binary
FROM scratch

COPY --from=builder /tmp/jp/cattled /cattled

ENTRYPOINT ["/cattled"]
