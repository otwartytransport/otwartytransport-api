# syntax=docker/dockerfile:1

FROM alpine:3.16.2

COPY ./build/server .

RUN apk add dumb-init
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD /server
