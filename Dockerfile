FROM golang:alpine AS builder
WORKDIR /app
ADD . .
RUN mkdir dist
RUN go build -o gh-pages-internal ./cmd/serve

FROM alpine
WORKDIR /app
COPY --from=builder /app/gh-pages-internal /bin/
ENTRYPOINT ["gh-pages-internal"]
