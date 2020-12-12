FROM golang:1.15.2-alpine3.12 as builder
# All these steps will be cached
WORKDIR /build
COPY . .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# Build the binary
RUN CGO_ENABLED=0 go build -o fs-backuper main.go
FROM alpine
COPY --from=builder /build/fs-backuper  .
ENTRYPOINT [ "./fs-backuper" ]