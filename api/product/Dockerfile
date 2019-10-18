# build image
FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR /go/src/github.com/cbellee/goShop-productService
ENV GO111MODULE=on
COPY . .
RUN go get -d -v \
  && GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o productservice .

# release image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates \
  && mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/cbellee/goShop-productService/productservice .
COPY --from=builder /go/src/github.com/cbellee/goShop-productService/seed_data.json .

# Metadata params
ARG VERSION
ARG BUILD_DATE
ARG VCS_URL
ARG VCS_REF
ARG NAME
ARG VENDOR

# Metadata
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name=$NAME \
      org.label-schema.description="multi-stage docker build" \
      org.label-schema.url="https://example.com" \
      org.label-schema.vcs-url=https://github.com/cbellee/$VCS_URL \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vendor=$VENDOR \
      org.label-schema.version=$VERSION \
      org.label-schema.docker.schema-version="1.0" \
      org.label-schema.docker.cmd="docker run -d productservice serve"

# port to listen on
EXPOSE 80/tcp

# entry command
ENTRYPOINT [ "./productservice" ]