FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
RUN cd /src && go build -o lyte

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/lyte /app/
ENTRYPOINT ./lyte