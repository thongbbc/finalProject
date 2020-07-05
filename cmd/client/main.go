package main

import (
	"context"
	modelProduct "finalProject/cmd/service/grpc-model/product"
	"finalProject/cmd/service/product"
	"flag"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	server := "0.0.0.0:5000"
	getId := flag.Int("id", 1, "an int")
	flag.Parse()
	fmt.Println("getId: ", *getId)
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := product.NewProductServiceClient(conn)
	testAdd(client)
	testGet(client, int32(*getId))
}

func testAdd(client product.ProductServiceClient) {
	addReq := modelProduct.AddReq{
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

func testGet(client product.ProductServiceClient, id int32) {

	findReq := modelProduct.GetReq{
		Id: id,
	}
	findRes, err := client.Get(context.TODO(), &findReq)

	if err != nil {
		panic(err)
	}
	fmt.Println("Response of get method is:", findRes)
}