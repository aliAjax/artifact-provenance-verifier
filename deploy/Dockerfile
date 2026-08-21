FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . .
RUN go build -trimpath -o /out/provenance ./cmd/server
FROM alpine:3.20
RUN adduser -D -u 10001 app
COPY --from=build /out/provenance /usr/local/bin/provenance
USER app
EXPOSE 8092
ENTRYPOINT ["/usr/local/bin/provenance"]
