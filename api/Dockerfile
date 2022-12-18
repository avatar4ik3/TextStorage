FROM golang:alpine as build-env

ARG app_env
# ARG app_port

ENV APP_ENV $app_env
# ENV APP_PORT $app_port
COPY api/ /go/src/avatar4ik3/TextStorage/api

WORKDIR /go/src/avatar4ik3/TextStorage/api/

RUN go get ./...

RUN go build -o /my_app
EXPOSE 5000:5000

# EXPOSE ${APP_PORT}:${APP_PORT}
CMD ["/my_app"]

