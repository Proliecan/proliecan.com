# compile in container 1
# run in container 2

FROM golang:latest as builder   

WORKDIR /app

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY src/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# ====

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/server /app/server

WORKDIR /app

EXPOSE 8080

# command: server -d localhost -p 8080 -v
# ALWAYS USE 0.0.0.0 FOR DOCKER!!! Why? no clue.
CMD ["./server", "-d", "0.0.0.0", "-p", "8080", "-v"]

# build command: docker build -t proliecan-website .
# run command: docker run -p 8080:8080 --name proliecan-website proliecan-website 