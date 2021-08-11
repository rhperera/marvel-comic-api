FROM golang:1.16

WORKDIR /app

COPY go.mod .
COPY go.sum .

# fetch dependencies
RUN go mod download
RUN go mod verify

# copy the source code as the last step
COPY . .

EXPOSE 8080