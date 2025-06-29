package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
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
	callDownload(client)

}

func callListFIles(client pb.FIleServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalf("could not list files: %v", err)
	}

	fmt.Println(res.GetFilenames())
}

func callDownload(client pb.FIleServiceClient) {
	req := &pb.DownloadRequest{
		FileName: "name.txt",
	}
	stream, err := client.Download(context.Background(), req)
	if err != nil {
		log.Fatalf("could not download file: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while receiving data: %v", err)
		}

		log.Printf("Response from download(bytes): %v", res.GetData())
		log.Printf("Response from download(string): %v", string(res.GetData()))
	}

}
