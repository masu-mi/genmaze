FROM golang:1.14.1-buster AS build

COPY . /work
WORKDIR /work

ARG COMMIT
ENV COMMIT=$COMMIT
RUN make fast


FROM alpine:3.11 AS release

RUN apk add --no-cache --update ca-certificates
COPY --from=build /work/bin/maze /bin/maze
COPY gen/http/openapi.* /gen/http/
COPY public/swagger-ui /public/swagger-ui

ENTRYPOINT ["/bin/maze"]
