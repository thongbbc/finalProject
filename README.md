# finalProject
BUILD PROJECT
- docker build --no-cache -f ./Dockerfile -t service .

RUN PROJECT
- docker-compose up

OR

BUILD PROJECT
- make build

RUN PROJECT
- make start-server

STOP SERVER
- make stop-server

WATCH LOG
- make watch-logs
