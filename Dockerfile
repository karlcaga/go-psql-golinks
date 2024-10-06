# syntax=docker/dockerfile:1

FROM golang:1.23.1 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build .

FROM debian:bookworm-slim AS run

COPY --from=build /app/go-psql-golinks /bin/go-psql-golinks

EXPOSE 8080

CMD [ "go-psql-golinks" ]