FROM golang:1.20-alpine as base

COPY . /api

WORKDIR /api/src/router

RUN go mod tidy

RUN go build -o router .

ARG PORT
ENV PORT=$PORT

EXPOSE $PORT

CMD ["./router"]
