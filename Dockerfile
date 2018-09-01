FROM golang:1.11-stretch

WORKDIR /go/src/github.com/ingenieux/ecr-login

COPY . /go/src/github.com/ingenieux/ecr-login

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && dep ensure && go get github.com/ingenieux/ecr-login/...

ENTRYPOINT /go/bin/ecr-login

