# Stage I: compile the source code with it's dependencies
FROM golang:latest as builder

WORKDIR /go/src/amadeus-go
COPY . .


RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o amadeus-go ./cmd/srv/.


# Stage II: run the service
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /go/bin/
COPY --from=builder /go/src/amadeus-go/amadeus-go .

EXPOSE 8000

CMD ["./amadeus-go"]
