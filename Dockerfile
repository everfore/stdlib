FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/
RUN go build -o gostdlib
RUN godoc -http=:9000 &
EXPOSE 9000 80
CMD ["/gopath/app/gostdlib"]