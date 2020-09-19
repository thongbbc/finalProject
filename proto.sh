 docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc \
 --go_out=plugins=grpc:.. \
 -I. ./proto/v1/model/error/error.proto

 docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc \
 --go_out=plugins=grpc:.. \
 -I. ./proto/v1/model/product/product.proto

 docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc \
 --go_out=plugins=grpc:.. \
 -I. ./proto/v1/model/user/user.proto

docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc \
--go_out=plugins=grpc:./cmd \
-I. ./proto/v1/user-service.proto

docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc \
--go_out=plugins=grpc:./cmd \
-I. ./proto/v1/product-service.proto