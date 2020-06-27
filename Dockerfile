FROM  golang:1.10
WORKDIR /Volumes/HDWEBSOFT/golang/finalProject
COPY . .
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "get", "github.com/golang/protobuf/proto"]
COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
RUN chmod 777 /usr/local/bin/docker-entrypoint.sh \
    && ln -s /usr/local/bin/docker-entrypoint.sh /
ENTRYPOINT ["docker-entrypoint.sh"]
