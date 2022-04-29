FROM golang:1.18.1-alpine3.14 AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ARG BuildNumber

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o claneventsbot -ldflags "-X main.buildNumber=$BuildNumber" *.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/claneventsbot .

## Export necessary port
#EXPOSE ${PORT}
#
## Command to run when starting the container
#CMD ["/dist/claneventsbot"]

# Build a small image
FROM scratch

ENV PORT=8080

COPY --from=builder /dist/claneventsbot /

# Export necessary port
EXPOSE ${PORT}

# Command to run
ENTRYPOINT ["/claneventsbot"]