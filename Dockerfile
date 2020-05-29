FROM golang:alpine AS builder
LABEL maintainer="Fajarhide <fajarhide@gmail.com>"

WORKDIR /build
# Let's cache modules retrieval - those don't change so often
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code necessary to build the application
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main

#multistage
FROM alpine:latest  
RUN apk update && apk --no-cache add ca-certificates
RUN apk add --update tzdata
ENV TZ=Asia/Jakarta
WORKDIR /app
COPY --from=builder /build/main .
CMD ["./main"] 