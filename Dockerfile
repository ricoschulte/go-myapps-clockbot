FROM golang:alpine AS builder
ARG VERSION
ENV VERSION=${VERSION} CGO_ENABLED=0
WORKDIR /app
COPY . .

# install go dependencies
RUN go mod tidy

# Test go
RUN go test ./... -v

# Build go
# -s Omit the symbol table and debug information.
# -w Omit the DWARF symbol table.
RUN CGO_ENABLED=0 go build \
    -ldflags "-s -w -X main.version=${VERSION:-dev}" \
    .

# strip the binary to reduce size  

# install strip with package binutils and valid timezone data
RUN apk update && apk add --no-cache binutils tzdata

# This will modify the binary in-place, removing the debugging information and making it smaller.
# Note that stripping a binary will make it difficult to debug if there are issues with it later on, so use this with caution.
RUN strip /app/go-myapps-clockbot

# make it executable
RUN chmod 775 /app/go-myapps-clockbot

FROM scratch 
ARG VERSION
WORKDIR /app
COPY --from=builder /app/go-myapps-clockbot /app/go-myapps-clockbot
COPY --from=builder /usr/share/zoneinfo/ /usr/share/zoneinfo/
LABEL \
  org.opencontainers.image.vendor="Rico Schulte" \
  org.opencontainers.image.title="go-myapps-clockbot" \
  org.opencontainers.image.version="${VERSION:-dev}"
VOLUME ["/data"]
ENTRYPOINT ["/app/go-myapps-clockbot"]