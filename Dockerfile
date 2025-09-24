FROM golang:1.24.5 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o ./hello-app ./cmd/...

FROM scratch AS final
COPY --from=builder /app/hello-app /hello-app
EXPOSE "8080"
ENTRYPOINT ["/hello-app"]