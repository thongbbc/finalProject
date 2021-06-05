#Generate Error proto
 docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc \
 --go_out=plugins=grpc:./cmd/service/gateway \
 --go_out=plugins=grpc:./cmd/service/user \
 --go_out=plugins=grpc:./cmd/service/product \
 -I. ./proto/v1/model/error/error.proto
 


#Generate model Product
 docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc \
 --go_out=plugins=grpc:./cmd/service/gateway \
 --go_out=plugins=grpc:./cmd/service/user \
 --go_out=plugins=grpc:./cmd/service/product \
 -I. ./proto/v1/model/product/product.proto

#Generate model User
 docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc \
 --go_out=plugins=grpc:./cmd/service/gateway \
 --go_out=plugins=grpc:./cmd/service/user \
 --go_out=plugins=grpc:./cmd/service/product \
 -I. ./proto/v1/model/user/user.proto

docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc \
--go_out=plugins=grpc:./cmd/service/gateway \
--go_out=plugins=grpc:./cmd/service/user \
--go_out=plugins=grpc:./cmd/service/product \
-I. ./proto/v1/user-service.proto

docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc \
--go_out=plugins=grpc:./cmd/service/gateway \
--go_out=plugins=grpc:./cmd/service/user \
--go_out=plugins=grpc:./cmd/service/product \
-I. ./proto/v1/product-service.proto