# syntax=docker/dockerfile:1
FROM golang:1.19-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o . ./cmd/main.go



FROM alpine
WORKDIR /app2
COPY --from=build /app .
CMD "./main"
LABEL autor="ctuzelov"
EXPOSE 7777
