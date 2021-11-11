FROM golang:1.17.2-alpine
MAINTAINER kemper0530

ENV GOPATH /go
ENV PATH=$PATH:$GOPATH/src

RUN apk update && \
    apk upgrade && \
    apk add vim && \
    apk add git && \
    apk add build-base

RUN apk add tzdata
ENV TZ=Asia/Tokyo

# 以下、Docker run 用の設定
ENV PATH=$PATH:$GOPATH/src/github.com/kemper0530/go-lambda-api-demo/src

WORKDIR $GOPATH/src/github.com/kemper0530/go-lambda-api-demo

COPY  /src $GOPATH/src/github.com/kemper0530/go-lambda-api-demo/src

RUN go mod init github.com/kemper0530/go-lambda-api-demo
RUN go get -v -t -d ./...
RUN go mod tidy

RUN GOOS=linux go build -o go-lambda-api-demo ./src

RUN chmod -R 777 $GOPATH/src/github.com/kemper0530/go-lambda-api-demo

ENTRYPOINT ["/go/src/github.com/kemper0530/go-lambda-api-demo/go-lambda-api-demo"]
