FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD ./src /src
RUN cd /src && go build -o lyte

# final stage
FROM golang:alpine AS runtime-env
WORKDIR /app
COPY --from=build-env /src/lyte /app/
ENTRYPOINT ./lyte