# base image
FROM alpine:latest

# MAINTAINER
MAINTAINER junwu <junwu.shao@gmail.com> 

COPY service /home/service

ENV PORT 50051

ENTRYPOINT ["/home/service"]
