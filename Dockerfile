# Dockerfile for the container image that would hold the buckets api service which would be created via a multi-stage build.

##
## Build stage
##

# Go official container image contains everything needed to compile and run go apps
FROM golang:1.20 AS app-builder

# Create a directory in the image that follows a similar path to the local project's main.go and make it the working directory as that's where the app would be run from
WORKDIR /go/src/github.com/Shayne3000/Buckets/cmd/

# Copy go.mod and go.sum into the working directory
COPY go.mod go.sum ./

# Download and install the required modules into the directory in the container image 
RUN go mod download

# Copy all the remaining app files as it is on your project into the respective directory, ideally it should be the working directory but we're using a cmd folder.
COPY . /go/src/github.com/Shayne3000/Buckets/

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/Buckets

##
## Run/Deploy stage
##
FROM alpine:latest

WORKDIR /

COPY --from=app-builder /bin/Buckets /usr/bin/Buckets

EXPOSE 8080

ENTRYPOINT ["/usr/bin/Buckets"]
