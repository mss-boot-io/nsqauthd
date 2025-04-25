FROM alpine

LABEL authors="lwnmengjing"

WORKDIR /app

COPY ./application /app/application

ENTRYPOINT ["/app/application"]