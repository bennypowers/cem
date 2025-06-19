FROM golang:1.24-bullseye

# Install mingw-w64 for cross-compilation
RUN apt-get update && apt-get install -y gcc-mingw-w64-x86-64

# Set environment variables for cross-compilation
ENV GOOS=windows
ENV GOARCH=amd64
ENV CGO_ENABLED=1
ENV CC=x86_64-w64-mingw32-gcc

WORKDIR /app
COPY . .

# Output binary to a location that matches the mounted volume
ARG OUTPUT=.
RUN go build -o ${OUTPUT}/cem.exe .
