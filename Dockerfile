FROM golang:latest AS builder
LABEL maintainer="Ejber Ozkan <ejber.ozkan@bbc.co.uk>"
ADD . /app/api
WORKDIR /app/api
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# stage to install telegraf
FROM alpine:latest AS setup-telegraf
ENV INFLUX_HOSTS \"http://influxdb:8086\"
ENV INFLUX_DB telegraf
ENV EXTRA_PLUGINS=""
#RUN echo "http://dl-2.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories
RUN echo "http://dl-2.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories
#RUN apk --no-cache add --update telegraf
RUN apk --no-cache add telegraf
RUN mkdir /etc/telegraf
COPY telegraf.conf /etc/telegraf/telegraf.conf
RUN sed -i -e "s|{{INFLUX_HOSTS}}|$INFLUX_HOSTS|g" -e "s|{{INFLUX_DB}}|$INFLUX_DB|g" -e "s|#{{EXTRA_PLUGINS}}|$EXTRA_PLUGINS\n#{{EXTRA_PLUGINS}}|" /etc/telegraf/telegraf.conf

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
RUN mkdir /etc/telegraf
COPY --from=setup-telegraf /etc/telegraf/telegraf.conf /etc/telegraf/telegraf.conf
COPY --from=setup-telegraf /usr/bin/telegraf /usr/bin/telegraf
COPY two-processes.sh two-processes.sh
RUN chmod +x ./two-processes.sh
#ENTRYPOINT ["./main"]
ENTRYPOINT ["./two-processes.sh"]
EXPOSE 8080