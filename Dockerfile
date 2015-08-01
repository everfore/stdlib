FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/
RUN go build -o gostdlib
RUN cp gostdlib /gopath/app/bin/
EXPOSE 80
CMD ["/gopath/app/bin/gostdlib"]