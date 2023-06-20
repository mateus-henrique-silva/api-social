FROM golang:1.16

WORKDIR /src/router
COPY . .

# Install Git
RUN apt-get update && apt-get install -y git

# Remove specific revision if present
RUN sed -i '/github.com\/uticket\/rest@v1.2.0/d' go.mod

# Download Go modules
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

ENTRYPOINT [ "gon", "run", "*.go" ]
