# Build image
FROM golang:1.21 AS build

WORKDIR /build

# Configure known hosts
RUN mkdir -p /root/.ssh && ssh-keyscan github.com >> /root/.ssh/known_hosts

# Copy source
COPY go.mod go.sum main.go ./
COPY internal ./internal

# Install dependencies and build
RUN go install
RUN go build -o shorty

# Distribution image
FROM debian:12

# Set default database path
ENV SHORTY_DATABASE_PATH /shorty/data

# Copy binary and set as command
COPY --from=build /build/shorty /usr/local/bin/shorty
ENTRYPOINT ["shorty"]
CMD ["start"]
