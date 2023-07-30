# Dockerfile for the container image that would hold the buckets api service which would be created via a multi-stage build.

##
## Build stage
##

# Use the Go official container image (with the alias app-builder) which contains everything needed to compile and run go apps) as the base image for the build stage. 
FROM golang:1.20 AS app-builder

# Copy all the files and directories as it is on your project and add them to the specified directory path (2nd argument) in the image's filesystem 
ADD . /go/src/github.com/Shayne3000/Buckets/

# Make the earlier created directory the working directory as that's where the app would be run from
WORKDIR /go/src/github.com/Shayne3000/Buckets/

# Download and install the required modules into the directory in the container image 
RUN go mod download

# Build the app's main package in main.go along with its dependencies and store the resultant binary Buckets in the /bin directory of the image's filesystem. 
# You can also store it in the root path as /Buckets 
RUN go build -o /bin/Buckets cmd/main.go

##
## Deploy stage
##

# Use the latest apline container image as the base image for deployment
FROM alpine:latest

WORKDIR /

# Copy the resultant binary from the build stage at the given path in the build image(app-builder)'s filesystem into the specified path of the deploy image
COPY --from=app-builder /bin/Buckets /usr/bin/Buckets

# Specify the network port that the deploy container listens on (for http requests) at runtime
EXPOSE 8080

# Run the Buckets binary at the given path
ENTRYPOINT ["/usr/bin/Buckets"]
