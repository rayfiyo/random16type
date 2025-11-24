FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o random16type

FROM gcr.io/distroless/base-debian12
WORKDIR /app

COPY --from=builder /app/random16type /app/random16type

EXPOSE 8080
USER nonroot:nonroot
CMD ["/app/random16type"]
