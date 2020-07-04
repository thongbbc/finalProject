package main

import (
	"context"
	product "finalProject/cmd/service/proto"
	"flag"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	server := "0.0.0.0:5000"
	id := *flag.Int("get", 1, "an int")
	flag.Parse()
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := product.NewProductServiceClient(conn)
	testAdd(client)
	testGet(client, int32(id))
}

func testAdd(client product.ProductServiceClient) {
	addReq := product.AddReq{
		Name:  "iphone6",
		Sku:   "IP6",
		Price: 1000000,
	}
	addRes, err := client.Add(context.TODO(), &addReq)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response of add method is:", addRes)
}

func testGet(client product.ProductServiceClient, id int32) {
	findReq := product.GetReq{
		Id: id,
	}
	findRes, err := client.Get(context.TODO(), &findReq)

	if err != nil {
		panic(err)
	}
	fmt.Println("Response of get method is:", findRes)
}
