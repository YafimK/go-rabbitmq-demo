FROM golang:1.12.6

ENV GO111MODULE=on

ADD . /
COPY . ./Tasker
WORKDIR /
RUN go build -o Tasker .
EXPOSE 8080
CMD ["/Tasker"]