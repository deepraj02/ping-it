FROM golang:alpine AS builder

RUN mkdir /build 

ADD . /build/

WORKDIR /build

RUN go build -o main .





FROM alpine 

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

ENV BASE_URL = ${BASE_URL}

CMD [ "./main" ]