FROM --platform=linux/amd64 alpine:3.16.2

RUN mkdir /gobin/

RUN apk update \
    && apk upgrade \
    && apk add --no-cache libc-dev \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates 2>/dev/null || true

# RUN apk add --no-cache libc6-compat
# RUN apk add --no-cache gcompat

# COPY /lib64/ld-linux-x86-64.so.2 /lib/
# RUN ln -s /lib/libc.musl-x86_64.so.1 /lib/ld-linux-x86-64.so.2
COPY homemanager /gobin

WORKDIR /gobin

ENTRYPOINT ["/gobin/homemanager"]

EXPOSE 8080
