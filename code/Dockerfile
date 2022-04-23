# The build image
FROM golang:1.18-alpine as build-image

# We need these to do the build
RUN apk update && apk add build-base

# Copy over the source and config
WORKDIR src
COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY Makefile ./

# Build the application
RUN make

# The runtime image
FROM alpine:3.14

# Add a new user to run as
RUN addgroup gouser && adduser -G gouser -D gouser
USER gouser

# Copy over the binary
WORKDIR app
COPY --from=build-image go/src/example-bind-variables .

# Runs the application at startup
ENTRYPOINT /app/example-bind-variables