FROM golang:1.14

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

WORKDIR /app/cmd/api

# build
RUN go build -o main .

EXPOSE 8080

CMD ["/app/cmd/api/main"]



