FROM golang:1.12.6

ENV GO111MODULE=on

ADD . /
COPY . ./BasicSend
WORKDIR /
RUN go build -o BasicSend .
CMD ["/BasicSend"]