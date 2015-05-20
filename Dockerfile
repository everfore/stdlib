FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/src/app
ADD . /gopath/src/app/

RUN go get github.com/shaalx/callback
RUN go install github.com/shaalx/callback

EXPOSE 80
CMD ["/gopath/app/bin/callback"]
