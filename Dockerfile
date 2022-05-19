FROM golang:1.18.2-alpine3.14


# Set the Current Working Directory inside the container
WORKDIR /telebot

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Run the executable
CMD ["./main"]