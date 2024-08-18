# Build the application from source
FROM golang:1.22-alpine AS build

WORKDIR /app

# Copy the go.mod and go.sum files specific to the service
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire service directory
COPY . .

# Build the specific module
RUN CGO_ENABLED=0 go build -o ./bin ./cmd

# Deploy the application binary into a lean image
FROM alpine:latest AS release

WORKDIR /app

COPY --from=build /app/bin ./bin

EXPOSE 8080
ENTRYPOINT ["./bin"]
