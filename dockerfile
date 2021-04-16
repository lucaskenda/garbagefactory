FROM golang:1.16

# Set the Current Working Directory inside the container
WORKDIR $GOPATH

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Build files
RUN go build garbageFactory.go

# Run the executable
ENTRYPOINT ["./garbageFactory"]
