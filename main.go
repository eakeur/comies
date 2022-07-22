package main

import (
	"comies/app/gateway/api/gen/menu"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	ctx := context.Background()

	conn, err := grpc.Dial("comies.herokuapp.com:443", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect with server: %v", err)
	}

	cli := menu.NewMenuClient(conn)

	res, err := cli.ListProducts(ctx, &menu.ListProductsRequest{})
	if err != nil {
		log.Fatalf("Failed to connect with server: %v", err)
	}

	log.Println(res)

}
