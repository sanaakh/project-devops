FROM golang:1.23 as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o myapp .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/myapp .
CMD ["./myapp"]
EXPOSE 8080
