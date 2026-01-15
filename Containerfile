FROM fedora:43

ENV LLVM_MINGW_VERSION=20251216
ENV LLVM_MINGW_DISTRO=ubuntu-22.04
ENV GOOS=windows
ENV GOARCH=amd64
ENV CGO_ENABLED=1
ENV PATH="/opt/llvm-mingw-$LLVM_MINGW_VERSION-ucrt-$LLVM_MINGW_DISTRO-x86_64/bin:${PATH}"

RUN dnf install -y golang mingw64-gcc mingw64-binutils mingw64-crt mingw64-headers curl tar coreutils
# Download and extract llvm-mingw (update version as needed from https://github.com/mstorsjo/llvm-mingw/releases)
RUN curl -LO https://github.com/mstorsjo/llvm-mingw/releases/download/${LLVM_MINGW_VERSION}/llvm-mingw-${LLVM_MINGW_VERSION}-ucrt-${LLVM_MINGW_DISTRO}-x86_64.tar.xz && \
    tar -xJf llvm-mingw-${LLVM_MINGW_VERSION}-ucrt-${LLVM_MINGW_DISTRO}-x86_64.tar.xz -C /opt && \
    rm llvm-mingw-${LLVM_MINGW_VERSION}-ucrt-${LLVM_MINGW_DISTRO}-x86_64.tar.xz

WORKDIR /app

COPY . .

# By default, run a shell script that picks the correct CC and output name for the arch
# Supports OUTPUT_DIR env var for go-release-workflows compatibility (defaults to dist/)
CMD bash -c '\
  set -e; \
  OUTDIR="${OUTPUT_DIR:-dist}"; \
  if [[ "$GOARCH" == "arm64" ]]; then \
    export CC="aarch64-w64-mingw32-clang"; \
    ARCH_SUFFIX="arm64"; \
  else \
    export CC="x86_64-w64-mingw32-gcc"; \
    ARCH_SUFFIX="x64"; \
  fi; \
  if [[ "$OUTDIR" == "dist/bin" ]]; then \
    OUT="$OUTDIR/cem-win32-$ARCH_SUFFIX.exe"; \
  else \
    OUT="$OUTDIR/cem-windows-$GOARCH.exe"; \
  fi; \
  mkdir -p "$OUTDIR"; \
  go build -ldflags="-s -w" -o "$OUT" . \
'
