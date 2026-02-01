# Stage 1: Build Frontend Assets
FROM node:18-alpine AS node-builder
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN chmod +x scripts/build-assets.sh && ./scripts/build-assets.sh

# Stage 2: Build Go App
FROM golang:1.25-alpine AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

# Stage 3: Final Image
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ffmpeg

# Copy binary from Go builder
COPY --from=go-builder /app/app .

# Copy built assets/docs from Node builder
COPY --from=node-builder /app/docs ./docs

EXPOSE 3002
CMD ["./app"]
