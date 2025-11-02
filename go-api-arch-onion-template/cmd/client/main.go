package main

import (
	"context"
	"fmt"
	"go-api-arch-onion-template/pkg"
	"go-api-arch-onion-template/userinterface/web/gin/api"
)

func main() {
	c, _ := api.NewClientWithResponses(pkg.GetEndpoint("/api/v1"))
	res, _ := c.GetAlbumByIdWithResponse(context.TODO(), 1)
	fmt.Println("sample", res)
}
