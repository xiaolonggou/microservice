package main

import (
	"fmt"
	"testing"

	"github.com/xiaolonggou/microservice/v1/sdk/client"
	"github.com/xiaolonggou/microservice/v1/sdk/client/arts"
)

func TestClient(t *testing.T) {

	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := arts.NewListArtPiecesParams()
	prod, err := c.Arts.ListArtPieces(params)
	//c.Arts.ListArtPieces(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v", prod.GetPayload()[0])
}
