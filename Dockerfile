FROM golang

RUN go get github.com/gorilla/mux
RUN go get github.com/leomarangoni/logmock
RUN go install github.com/leomarangoni/logmock

EXPOSE 8000
ENTRYPOINT /go/bin/logmock
