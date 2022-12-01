FROM golang:1.17.9-alpine3.15 as build-env

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY . .

#RUN go clean -modcache

RUN go mod tidy 

RUN go get -d -v ./...

RUN CGO_ENABLED=0 go test ./... -v
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static

COPY --from=build-env /go/bin/app /
COPY --from=build-env /app/config.toml /
COPY --from=build-env /app/map_data.json /
EXPOSE 8080
CMD ["/app"]
CMD ["/app"]
