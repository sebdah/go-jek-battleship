FROM golang:1.7-alpine

RUN mkdir -p /go/src/github.com/sebdah/go-jek-battleship

ADD . /go/src/github.com/sebdah/go-jek-battleship/
WORKDIR /go/src/github.com/sebdah/go-jek-battleship/

RUN go build -o /go/bin/battleship ./cmd

CMD ["true"]
