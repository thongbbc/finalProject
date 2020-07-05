package main

import (
	"finalProject/cmd/service/gateway/router"
	"fmt"
	"strconv"
)


func main() {
	r := router.InitGateway()
	fmt.Println("====================================")
	fmt.Println("START gateway service at 3000")
	fmt.Println("====================================")
	// Connect grpc
	r.Run(":"+strconv.Itoa(3000))
}