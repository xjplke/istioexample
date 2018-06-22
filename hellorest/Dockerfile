# base image
FROM alpine:latest

# MAINTAINER
MAINTAINER junwu <junwu.shao@gmail.com> 

COPY hellorest /home/hellorest

ENV PORT 8123
ENV HELLO_PORT 50050
ENV HELLO_SERVICE localhost

ENTRYPOINT ["/home/hellorest"]
