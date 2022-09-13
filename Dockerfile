FROM alpine

RUN mkdir /gobin/

RUN apk update \
    && apk upgrade \
    && apk add --no-cache libc-dev \
    && apk add --no-cache ca-certificates \
    && apk add libc6-compat \
    && update-ca-certificates 2>/dev/null || true

COPY homemanager /gobin

WORKDIR /gobin

ENTRYPOINT ["/gobin/homemanager"]

EXPOSE 8080
