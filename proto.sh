docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --go_out=. -I. ./cmd/api/product.proto