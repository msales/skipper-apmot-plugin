FROM golang:1.12

WORKDIR /plugin
COPY ./ .

ENV GO111MODULE on
RUN GOOS=linux go build -buildmode=plugin -o apmot.so plugin.go

# You can run `docker cp <container_id>:/plugin/ampot.so ./` to retrieve the plugin for use inside a docker container.
CMD ["bash"]