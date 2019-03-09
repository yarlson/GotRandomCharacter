# build stage
FROM golang:alpine as build

RUN apk --no-cache add ca-certificates

ADD got_random.go /
RUN cd / && CGO_ENABLED=0 GOOS=linux go build  -ldflags '-w -s' -a -installsuffix cgo -o /usr/bin/got_random

# final stage
FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/bin/got_random /usr/bin/got_random
CMD ["/usr/bin/got_random"]
