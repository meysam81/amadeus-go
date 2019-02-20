# Stage I: compile the source code with it's dependencies
FROM golang:latest as builder

WORKDIR /go/src/amadeus-go
COPY amadeus-go .
COPY api /go/src/api


RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o amadeus-go ./cmd/srv/.


# Stage II: run the service
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /go/bin/
COPY --from=builder /go/src/amadeus-go/amadeus-go .

# api token retrieval information
ENV API_KEY kDy2Detv2VNmS6pyf6AFFSQbJVtPLSVe
ENV API_SECRET mt52fS9wcGRn2RHE

## authorization url
ENV API_URL https://test.api.amadeus.com/v1/security/oauth2/token

## url for every service in the server
ENV API_BASE_URL https://test.api.amadeus.com

EXPOSE 8000

CMD ["./amadeus-go"]
