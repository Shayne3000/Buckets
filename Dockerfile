# Dockerfile for the container image that would hold the buckets api service and be created via a multi-stage build.

##
## Build stage
##

# Go official container image contains everything needed to compile and run go apps
FROM golang:1.20 AS app-builder

# Create a directory in the image that follows a similar path to the local project
WORKDIR /go/src/github.com/Shayne3000/Buckets/

# Copy go.mod and go.sum inot the working directory
COPY go.mod go.sum ./

# Download and install the required modules into a directory in the container image 
RUN go mod download

# Copy all the remaining app files into the working directory
COPY . ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/Buckets

##
## Run/Deploy stage
##
FROM alpine:latest
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=app-builder /bin/Buckets /usr/bin/Buckets
EXPOSE 8080
ENTRYPOINT ["/usr/bin/Buckets"]
