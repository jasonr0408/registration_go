# Start from golang:1.12-alpine base image
FROM golang:1.12-alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 9487
CMD ["./main"]