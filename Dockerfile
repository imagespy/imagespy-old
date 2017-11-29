FROM alpine:3.6

RUN apk add ca-certificates --update-cache

RUN mkdir /app && mkdir /app/ui

WORKDIR /app

COPY ./ui/build /app/ui

COPY ./app/imagespy /app/imagespy

ENTRYPOINT ["./imagespy"]
