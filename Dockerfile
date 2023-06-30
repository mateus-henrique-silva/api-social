FROM golang:1.16

WORKDIR /src/router
COPY . .

# Install Git
RUN apt-get update && apt-get install -y git

# Remove specific revision if present
RUN go clean -modcache

# Download Go modules
RUN go mod download

# RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

RUN go run *go
