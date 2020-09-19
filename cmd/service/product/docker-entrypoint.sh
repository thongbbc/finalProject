#!/bin/bash
CompileDaemon -log-prefix=false -build="go build ./cmd/service/product" -command="./product"
