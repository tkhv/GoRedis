FROM golang:1.21.6-alpine3.19 AS build
WORKDIR /app
COPY go.mod .
COPY *.go .
RUN go build -o main .

FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE 6379

ENTRYPOINT [ "/app/main" ]