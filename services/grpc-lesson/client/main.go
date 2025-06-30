package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewFIleServiceClient(conn)
	// callListFIles(client)
	// callDownload(client)
	callUpload(client)

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

func callUpload(client pb.FIleServiceClient) {
	fileName := "sports.txt"
	path := "/home/naha/dev/github/go-grpc/services/grpc-lesson/storage/" + fileName
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	stream, err := client.Upload(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	buf := make([]byte, 5)
	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		req := pb.UploadRequest{
			Data: buf[:n],
		}
		sendErr := stream.Send(&req)
		if sendErr != nil {
			log.Fatalf("error while sending data: %v", sendErr)
		}

		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v", err)
	}

	log.Printf("Upload response: %v", res.GetSize())
}
