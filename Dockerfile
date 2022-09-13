FROM alpine

RUN mkdir /gobin/

RUN apk update \
    && apk upgrade \
    && apk add --no-cache libc-dev \
    && apk add --no-cache ca-certificates \
    && apk add --no-cache libc6-compat \
    && update-ca-certificates 2>/dev/null || true

COPY /lib64/ld-linux-x86-64.so.2 /lib/
COPY homemanager /gobin

WORKDIR /gobin

ENTRYPOINT ["/gobin/homemanager"]

EXPOSE 8080
