FROM golang

MAINTAINER abhinaba@gmail.com

RUN apt-get install -y git
RUN git clone https://github.com/bonggeek/element.git /go/src/github.com/bonggeek/element
RUN git clone https://github.com/bonggeek/element.git /go/src/github.com/bonggeek/element


EXPOSE 8080

ENTRYPOINT ["git", "version"]
