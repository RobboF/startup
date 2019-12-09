FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD ./src /go/src
RUN cd /go/src && go build -o goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/goapp /app/
ENTRYPOINT ./goapp