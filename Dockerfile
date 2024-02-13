# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH environment variable at /go.
FROM golang:1.17 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY cmd/go.mod ./
COPY cmd/go.sum ./
RUN go mod download

# Copy local code to the container image.
COPY cmd/ ./

# Build the binary.
# -o app specifies the output name of the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

# Use the official Alpine image for a lean production container.
FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/app /app

# Run the web service on container startup.
CMD ["/app"]
