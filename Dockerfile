FROM golang:1.23 as builder
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o provider ./cmd/provider

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/provider .
USER nonroot:nonroot
ENTRYPOINT ["/provider"]
