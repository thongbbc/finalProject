package main

import (
	"context"
	product "finalProject/cmd/service/proto"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	server := "0.0.0.0:5000"
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := product.NewProductServiceClient(conn)
	testAdd(client)
	testGet(client)
}

func testAdd(client product.ProductServiceClient) {
	addReq := product.AddReq{
		Name: "iphone6",
		Sku: "IP6",
		Price: 1000000,
	}
	addRes, err := client.Add(context.TODO(), &addReq)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response of add method is:", addRes)
}

func testGet(client product.ProductServiceClient) {
	findReq := product.GetReq{
		Id:1,
	}
	findRes, err := client.Get(context.TODO(), &findReq)

	if err != nil {
		panic(err)
	}
	fmt.Println("Response of get method is:", findRes)
}