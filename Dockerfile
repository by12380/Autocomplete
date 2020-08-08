
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go build -o /go/bin/app main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/assets /assets/
ENTRYPOINT ./app
LABEL Name=autocomplete Version=0.0.1
EXPOSE 8080
