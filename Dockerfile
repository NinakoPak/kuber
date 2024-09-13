FROM golang:1.20-alpine3.16 as build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM alpine:3.16
COPY --from=build /bin/demo /bin/demo

ENTRYPOINT ["/bin/demo"]

