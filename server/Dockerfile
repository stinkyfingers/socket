FROM golang

ADD . /go/src/github.com/stinkyfingers/socket/server

RUN go get github.com/stinkyfingers/socket/server
RUN go install github.com/stinkyfingers/socket/server

ENTRYPOINT /go/bin/server

EXPOSE 7000