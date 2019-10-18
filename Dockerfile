FROM golang:latest AS builder
LABEL maintainer="Ejber Ozkan <ejber@ozkan.co.uk>"
ADD . /app/api
WORKDIR /app/api
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
RUN chmod +x ./main
ENTRYPOINT ["./main"]
EXPOSE 8080