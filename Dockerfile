FROM golang:alpine as build-env
LABEL maintainer="Reza Chrismardianto <rezachrismardianto20@gmail.com>"
ARG SERVICE_NAME=service-user
RUN mkdir /app
ADD . /app/

RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta
WORKDIR /app
RUN go build -o ${SERVICE_NAME} .

FROM alpine
WORKDIR /app
COPY --from=build-env /app/${SERVICE_NAME}          /app/${SERVICE_NAME}

RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta

ENTRYPOINT ["/app/service-user"]
