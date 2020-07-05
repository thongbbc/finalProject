package main

import (
	"bytes"
	"encoding/json"
	"finalProject/cmd/service/gateway/router"
	productModel "finalProject/cmd/service/grpc-model/product"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestProductRoute(t *testing.T) {
	router := router.InitGateway()
	wGet := httptest.NewRecorder()
	wPost := httptest.NewRecorder()
	message := map[string]interface{}{
		"sku": "P123",
		"price":  1000,
	}

	bytesRepresentation, err := json.Marshal(message)
	assert.Equal(t, nil, err)
	reqPost, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(bytesRepresentation))
	router.ServeHTTP(wPost, reqPost)

	// Expect post success
	assert.Equal(t, 200, wPost.Code)

	product := &productModel.AddRes{}


	r := bytes.NewReader(wPost.Body.Bytes())
	decoder := json.NewDecoder(r)

	errDecode := decoder.Decode(product)

	assert.Equal(t, nil, errDecode)
	fmt.Println(product)
	idProduct := strconv.Itoa(int(product.GetProduct().Id))
	expectedRes := fmt.Sprintf("{\"product\":{\"id\":%s,\"sku\":\"P123\",\"price\":1000}}", idProduct)

	assert.Equal(t, expectedRes, wPost.Body.String())
	reqGet, _ := http.NewRequest("GET", fmt.Sprintf("/product/%s", idProduct), nil)
	router.ServeHTTP(wGet, reqGet)
	// Expect get success
	expectedRes = fmt.Sprintf("{\"product\":{\"id\":%s,\"sku\":\"P123\",\"price\":1000}}", idProduct)
	assert.Equal(t, 200, wGet.Code)
	assert.Equal(t, expectedRes, wGet.Body.String())
}
