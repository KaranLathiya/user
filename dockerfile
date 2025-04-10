# Start a new stage from scratch
FROM alpine:latest as wait

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait

RUN chmod +x /wait

FROM golang:1.22.4 as builder

WORKDIR /app

COPY . . 

RUN go mod vendor -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main user/cmd

# Start a new stage from scratch
FROM alpine:latest

COPY --from=wait /wait /wait

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /cmd/main

COPY --from=builder /app/config/.env /config/.env

RUN cat /config/.env

EXPOSE 8000

CMD /wait && cmd/main
