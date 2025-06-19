FROM golang:1.24-bullseye

# Set build arguments with defaults
ARG GOARCH=amd64
ARG CC=x86_64-w64-mingw32-gcc

# Install cross-compilers for both x64 and arm64
RUN apt-get update && \
    apt-get install -y gcc-mingw-w64-x86-64 aarch64-w64-mingw32-gcc

# Set environment variables based on build args
ENV GOOS=windows
ENV CGO_ENABLED=1

# Set dynamically from build args
ENV GOARCH=${GOARCH}
ENV CC=${CC}

WORKDIR /app
COPY . .

RUN go build -o cem.exe .
