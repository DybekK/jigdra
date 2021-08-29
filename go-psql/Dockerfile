FROM golang:latest

WORKDIR /go-psql

COPY /go-psql/go.mod .
COPY /go-psql/go.sum .

RUN go mod download

COPY /go-psql /go-psql

RUN go build

CMD ["./go-psql"]
