FROM golang:1.20-alpine as build

RUN mkdir -p /go/src/github.com/merge/shopping-card && \
    apk add --no-cache git openssh-client make gcc libc-dev

WORKDIR /go/src/github.com/merge/shopping-card

COPY . .

RUN make build

FROM alpine:latest

COPY --from=build /go/src/github.com/merge/shopping-card/build/bin/apiserver /usr/bin/apiserver
EXPOSE 8000
CMD ["/usr/bin/apiserver"]