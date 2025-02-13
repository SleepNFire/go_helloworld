FROM golang:1.24 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/project ./cmd/project/project.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/project .
CMD ["./project"]