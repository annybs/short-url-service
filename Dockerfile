# Build image
FROM golang:1.21.3 AS build

WORKDIR /build

# Configure known hosts
RUN mkdir -p /root/.ssh && ssh-keyscan github.com >> /root/.ssh/known_hosts

# Copy source
COPY go.mod go.sum main.go ./

# Install dependencies and build
RUN go install
RUN go build -o short-url-service

# Distribution image
FROM debian:12

# Copy binary and set as command
COPY --from=build /build/short-url-service /usr/local/bin/short-url-service
CMD ["short-url-service"]
