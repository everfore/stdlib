FROM google/golang
MAINTAINER Shaalx Shi "60026668.m@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/
RUN go build -o gostdlib
RUN chmod u+x /gopath/app/gostdlib
RUN chmod u+x /gopath/app/run.sh
EXPOSE 9000:9000 80:80
CMD ["/gopath/app/run.sh"]