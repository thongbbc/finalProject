FROM  golang:latest
WORKDIR /Volumes/HDWEBSOFT/golang/finalProject
COPY . .
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
COPY /cmd/service/.env /cmd/service/.env

COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh

RUN chmod 777 /usr/local/bin/docker-entrypoint.sh \
    && ln -s /usr/local/bin/docker-entrypoint.sh /
RUN go mod download
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/service/gateway/" -command="./gateway"

