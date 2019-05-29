FROM golang

RUN go get github.com/leomarangoni/logmock
RUN go install github.com/leomarangoni/logmock

ENTRYPOINT /go/bin/logmock
