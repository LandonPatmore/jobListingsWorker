FROM golang:1.11.5

LABEL maintainer = "Landon Patmore <landon.patmore@gmail.com>"

WORKDIR $GOPATH/src/dataPullerWorker

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["dataPullerWorker"]