FROM golang:alpine as builder

RUN apk update && apk add git

COPY . $GOPATH/src/bitbucket.org/malonewebdev/creamy-static
WORKDIR $GOPATH/src/bitbucket.org/malonewebdev/creamy-static

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/creamy-static

FROM scratch

COPY --from=builder /go/bin/creamy-static /go/bin/creamy-static
ENTRYPOINT ["/go/bin/creamy-static"]
