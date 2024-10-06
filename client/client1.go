package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"time"

	pb "grpc-example/pb_files"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:50052"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewExampleServiceClient(conn)

	ctx := context.Background()

	UnaryClient(ctx, client)

	// ClientStream(ctx, client)

	// ServerStream(ctx, client)

	// BidirectionalStream(ctx, client)
}

func UnaryClient(ctx context.Context, client pb.ExampleServiceClient) {
	// Create metadata
	md := metadata.New(map[string]string{
		"token":        "dververbt3",
		"organization": "xyz",
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	jsonMsg := map[string]string{
		"Key1": "value1",
		"Key2": "value2",
	}
	jm, _ := json.Marshal(jsonMsg)

	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: "amizhthan", Action: pb.Actions_CREATE, Message: string(jm)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func ClientStream(ctx context.Context, client pb.ExampleServiceClient) {
	stream, err := client.SendMessages(ctx)
	if err != nil {
		log.Fatalf("could not send messages: %v", err)
	}

	messages := []string{"Hello", "World", "gRPC", "Client Streaming"}
	for _, msg := range messages {
		if err := stream.Send(&pb.HelloRequest{Message: msg}); err != nil {
			log.Fatalf("could not send message: %v", err)
		}
		time.Sleep(time.Second) 
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not receive reply: %v", err)
	}
	log.Printf("Response: %s", reply.GetMessage())
}

func ServerStream(ctx context.Context, client pb.ExampleServiceClient) {
	stream, err := client.GetMessages(ctx, &pb.HelloRequest{Name: "Client"})
	if err != nil {
		log.Fatalf("could not get messages: %v", err)
	}

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not receive message: %v", err)
		}
		log.Printf("Received message: %s", reply.GetMessage())
	}
}

func BidirectionalStream(ctx context.Context, client pb.ExampleServiceClient) {
	stream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalf("could not chat: %v", err)
	}

	messages := []string{"Hello", "World", "gRPC", "Bidirectional Streaming"}
	waitc := make(chan struct{})

	go func() {
		for _, msg := range messages {
			if err := stream.Send(&pb.ChatMessage{Message: msg}); err != nil {
				log.Fatalf("could not send message: %v", err)
			}
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			reply, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("could not receive message: %v", err)
			}
			log.Printf("Received message: %s", reply.GetMessage())
		}
	}()

	<-waitc
}
