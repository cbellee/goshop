# build image
FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
ENV GO111MODULE=on
WORKDIR /app
COPY /go.mod .
COPY /go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o customerservice

# release image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates \
  && mkdir /app
WORKDIR /app
COPY --from=builder /customerservice .
COPY --from=builder /seed_data.json .

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
      org.label-schema.docker.cmd="docker run -d service serve"

# port to listen on
EXPOSE 80/tcp

# entry command
ENTRYPOINT [ "./customerservice" ]