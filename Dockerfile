FROM golang:1.16-alpine3.12 as builder

WORKDIR $GOPATH/src/github.com/sungjunyoung/msatodo/
RUN apk update && apk upgrade && apk add --no-cache alpine-sdk
COPY . $GOPATH/src/github.com/sungjunyoung/msatodo/
RUN make build
RUN mv msatodo /

#-------------------------------------------

FROM alpine:3.7
COPY --from=builder /msatodo /bin/

ENTRYPOINT ["/bin/msatodo"]