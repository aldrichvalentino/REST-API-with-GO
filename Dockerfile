# image
FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY . .

# dependency
RUN apk add --no-cache git curl

# package manager
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# install dependency
RUN dep ensure

# monitoring
RUN go get -u github.com/pilu/fresh

# execute a watcher
CMD ["fresh", "-c", "runner.conf"]

EXPOSE 8080
