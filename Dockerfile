##
## Build
##
FROM docker.io/golang as portunus_builder
WORKDIR /opt/app-root
ADD . /opt/app-root
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/portunus

##
## Deploy
##
FROM docker.io/alpine:latest
COPY --from=portunus_builder /opt/app-root/bin/portunus /usr/bin/portunus
EXPOSE 8080
ENTRYPOINT [ "portunus" ]
