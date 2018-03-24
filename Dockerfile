FROM golang:1.9

WORKDIR /go/src/github.com/homespendapi
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["homespendapi"]