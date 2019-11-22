FROM golang:1.12-alpine AS build-env
RUN apk --no-cache add build-base git
ADD . /src
RUN cd /src && go build -o goapp

# final stage
FROM alpine:3.10.3
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build-env /src/goapp /app/
EXPOSE 5000
ENTRYPOINT ./goapp
