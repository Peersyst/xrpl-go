FROM golang:1.23-alpine AS install
RUN apk add --no-cache make git ca-certificates

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM install AS lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8
RUN make lint

FROM lint AS test
RUN make test-ci