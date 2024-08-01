FROM golang:1.20-alpine3.18
RUN mkdir /build
WORKDIR /build
COPY . .

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN go install -v ./...

FROM alpine:3.18
RUN apk add --no-cache \
    ca-certificates \
    tzdata

RUN ln -fs /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
WORKDIR /

COPY --from=0 /go/bin/* /usr/bin/

ENTRYPOINT [ "server" ]
