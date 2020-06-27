FROM  golang:1.10
WORKDIR /Volumes/HDWEBSOFT/golang/finalProject
COPY . .
RUN ["go", "get", "github.com/go-playground/justdoit"]

ENTRYPOINT justdoit -watch="./" -include="(.+\\.go|.+\\.c)$" -build="go build ./cmd/api/" -run="./api"