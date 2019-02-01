# Use the official go docker image built on debian.
FROM golang:1.11.2

# Grab the source code and add it to the workspace.
ADD . /go/src/github.com/skylerjaneclark/buddy-api

# Install revel and the revel CLI.
RUN go get github.com/golang/dep/cmd/dep && \
    go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel

#Install glide and build utilities
RUN go get golang.org/x/tools/go/buildutil &&\
    go get github.com/Masterminds/glide

#Install redis

# Use the revel CLI to start up our application.
ENTRYPOINT revel run github.com/skylerjaneclark/buddy-api dev  8080

# Open up the port where the app is running.
EXPOSE 8080
