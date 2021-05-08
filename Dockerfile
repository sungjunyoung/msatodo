FROM golang:1.16-alpine3.12 as builder

WORKDIR $GOPATH/src/github.com/sungjunyoung/prototodo/
RUN apk update && apk upgrade && apk add --no-cache alpine-sdk
COPY . $GOPATH/src/github.com/sungjunyoung/prototodo/
RUN make build
RUN mv prototodo /

#-------------------------------------------

FROM alpine:3.7
COPY --from=builder /prototodo /bin/

ENTRYPOINT ["/bin/prototodo"]