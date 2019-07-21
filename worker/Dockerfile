FROM golang:1.12.6

ENV GO111MODULE=on

ADD . /
COPY . ./Worker
WORKDIR /
RUN go build -o Worker .
CMD ["/Worker"]