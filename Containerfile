FROM fedora:43

ENV LLVM_MINGW_VERSION=20250613
ENV LLVM_MINGW_DISTRO=ubuntu-22.04
ENV GOOS=windows
ENV GOARCH=amd64
ENV CGO_ENABLED=1
ENV PATH="/opt/llvm-mingw-$LLVM_MINGW_VERSION-ucrt-$LLVM_MINGW_DISTRO-x86_64/bin:${PATH}"

RUN dnf install -y golang mingw64-gcc mingw64-binutils mingw64-crt mingw64-headers curl tar coreutils
# Download and extract llvm-mingw (latest as of July 2025, update as needed)
RUN curl -LO https://github.com/mstorsjo/llvm-mingw/releases/download/${LLVM_MINGW_VERSION}/llvm-mingw-${LLVM_MINGW_VERSION}-ucrt-${LLVM_MINGW_DISTRO}-x86_64.tar.xz && \
    tar -xJf llvm-mingw-${LLVM_MINGW_VERSION}-ucrt-${LLVM_MINGW_DISTRO}-x86_64.tar.xz -C /opt && \
    rm llvm-mingw-${LLVM_MINGW_VERSION}-ucrt-${LLVM_MINGW_DISTRO}-x86_64.tar.xz

WORKDIR /app

COPY . .

# By default, run a shell script that picks the correct CC and output name for the arch
CMD bash -c '\
  set -e; \
  export OUT="dist/cem-windows-$GOARCH.exe"; \
  if [[ "$GOARCH" == "arm64" ]]; then \
    export CC="aarch64-w64-mingw32-clang"; \
  else \
    export CC="x86_64-w64-mingw32-gcc"; \
  fi; \
  mkdir -p dist; \
  go build -o "$OUT" . \
'
