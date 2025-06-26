package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewFIleServiceClient(conn)
	callListFIles(client)

}

func callListFIles(client pb.FIleServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalf("could not list files: %v", err)
	}

	fmt.Println(res.GetFilenames())
}
