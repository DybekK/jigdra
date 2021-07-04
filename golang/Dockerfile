FROM golang:latest

WORKDIR /golang

COPY /golang/go.mod .
COPY /golang/go.sum .

RUN go mod download

COPY /golang /golang

RUN go build

CMD ["./golang"]
