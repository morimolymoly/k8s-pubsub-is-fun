FROM golang:latest

ADD . /go/src/github.com/morimolymoly/pubsub-is-fun

WORKDIR /go/src/github.com/morimolymoly/pubsub-is-fun
RUN go install github.com/morimolymoly/pubsub-is-fun/src/pubsub-is-fun
CMD ["/go/bin/pubsub-is-fun"]
