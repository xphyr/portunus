FROM docker.io/golang as portunus_builder
WORKDIR /opt/app-root/
COPY portunus.go go.mod pkg ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -o bin/portunus portunus.go


FROM docker.io/alpine:latest
COPY --from=portunus_builder /opt/app-root/bin/portunus /usr/bin/portunus
ENTRYPOINT [ "portunus" ]
