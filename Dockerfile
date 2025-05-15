# -------- DEV STAGE --------
FROM golang:latest AS dev

ENV PATH="/go/bin:${PATH}"

# Install Air and Delve for development
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install golang.org/x/tools/cmd/goimports@latest

WORKDIR /app
CMD ["air"]

# -------- BUILD STAGE --------
FROM golang:latest AS builder

WORKDIR /build

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN go build -o sqlitegen ./cmd

# -------- PRODUCTION STAGE --------
FROM debian:bullseye-slim

WORKDIR /app

# Copy only the built binary
COPY --from=builder /build/sqlitegen /app/sqlitegen

# Run the binary
ENTRYPOINT ["./sqlitegen"]

