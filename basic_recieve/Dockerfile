FROM golang:1.12.6

ENV GO111MODULE=on

ADD . /
COPY . ./BasicRecieve
WORKDIR /
RUN go build -o BasicRecieve .
CMD ["/BasicRecieve"]