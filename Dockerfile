FROM golang:1.17.2-alpine
MAINTAINER kemper0530

ENV GOPATH /go
ENV PATH=$PATH:$GOPATH/src

# 以下、Docker run 用の設定
ENV PATH=$PATH:$GOPATH/src/github.com/kemper0530/go-lambda-api-demo/src
WORKDIR $GOPATH/src/github.com/kemper0530/go-lambda-api-demo
COPY  /src $GOPATH/src/github.com/kemper0530/go-lambda-api-demo/src

RUN go mod init go-lambda-api-demo
RUN go mod tidy

RUN GOOS=linux go build -o go-lambda-api-demo ./src

ENTRYPOINT ["/go/src/github.com/kemper0530/go-lambda-api-demo/go-lambda-api-demo"]
