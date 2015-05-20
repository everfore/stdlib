FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/src/callback
ADD . /gopath/src/callback/

RUN go get github.com/shaalx/callback
RUN go install github.com/shaalx/callback

EXPOSE 80
CMD ["/gopath/app/bin/callback"]
