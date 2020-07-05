package driver

import (
	"finalProject/cmd/service/product"
	"fmt"
	"google.golang.org/grpc"
)


func ConnectProduct(host string, port string) product.ProductServiceClient {
	grpcUrl := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(grpcUrl, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := product.NewProductServiceClient(conn)
	return client
}

