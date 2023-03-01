FROM golang:1.19

WORKDIR /go/app/src

CMD ["go", "run", "main.go"]