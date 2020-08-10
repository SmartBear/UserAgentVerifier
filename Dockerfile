#FROM golang:alpine
#WORKDIR /app
#
## Fetch dependencies
#COPY app/go.mod ./
#RUN go mod download
#
## Build
#COPY app/main.go ./
#RUN CGO_ENABLED=0 go build
#
### Create final image
##FROM golang:alpine
##COPY --from=builder /build/app/main .
#EXPOSE 8080
#CMD ["./main"]

FROM golang:alpine as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /build/app

# Fetch dependencies
COPY app/go.mod ./
RUN go mod download

# Build
COPY app/main.go ./
RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build

# Create final image
FROM golang
WORKDIR /root
COPY --from=builder /build/app/main .
EXPOSE 8080
CMD ["./main"]
#FROM golang
#
#WORKDIR /
#COPY app/. .
#
#RUN go mod download
#
#EXPOSE 8080
#
#CMD ["go", "run", "main.go"]