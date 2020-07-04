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
	id := flag.Int("get", 1, "an int")
	updateID := flag.Int("update", 0, "an int")
	delID := flag.Int("del", 0, "an int")

	flag.Parse()
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := product.NewProductServiceClient(conn)
	testAdd(client)
	testGet(client, int32(*id))
	if *updateID != 0 {
		testUpdate(client, int32(*updateID))
	}
	if *delID != 0 {
		testDelete(client, int32(*delID))
	}
}

func testAdd(client product.ProductServiceClient) {
	addReq := product.AddReq{
		Name:     "iphone6",
		Sku:      "IP6",
		Price:    1000000,
		Quantity: 1,
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

func testUpdate(client product.ProductServiceClient, id int32) {
	upateReq := product.UpdateReq{&product.Product{
		Id:    id,
		Name:  "iphone7",
		Sku:   "IP7",
		Price: 10000,
	},
	}
	upateRes, err := client.Update(context.TODO(), &upateReq)

	if err != nil {
		panic(err)
	}
	fmt.Println("Response of update method is:", upateRes)
}

func testDelete(client product.ProductServiceClient, id int32) {
	delReq := product.DeleteReq{
		Id: id,
	}
	delRes, err := client.Delete(context.TODO(), &delReq)

	if err != nil {
		panic(err)
	}
	fmt.Println("Response of delete method is:", delRes)
}
