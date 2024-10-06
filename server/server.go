package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	pb "grpc-example/pb_files"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	pb.UnimplementedExampleServiceServer
}

// Unary
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		token := md["token"][0]
		fmt.Println("token: ", token)
	}

	var msg map[string]interface{}

	json.Unmarshal([]byte(req.Message), &msg)
	fmt.Println("message", msg)

	rep := &pb.HelloReply{
		Message: "Hello" + req.Name + req.Action.String(),
	}
	return rep, nil
}

// Client Stream
func (s *server) SendMessages(stream pb.ExampleService_SendMessagesServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// End of stream
			break
		}
		if err != nil {
			return err
		}
		messages = append(messages, req.GetMessage()) // Correctly access the "Name" field
	}
	response := "Received messages: " + strings.Join(messages, ", ")
	return stream.SendAndClose(&pb.HelloReply{Message: response})
}

func (s *server) GetMessages(req *pb.HelloRequest, stream pb.ExampleService_GetMessagesServer) error {
	messages := []string{"Hello", "World", "gRPC", "Server Streaming"}
	for _, msg := range messages {
		if err := stream.Send(&pb.HelloReply{Message: msg}); err != nil {
			return err
		}
		time.Sleep(time.Second) // Simulate delay between messages
	}
	return nil
}

func (s *server) Chat(stream pb.ExampleService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received message: %s", msg.GetMessage())

		// Echo the message back to the client
		if err := stream.Send(&pb.ChatMessage{Message: "Server received: " + msg.GetMessage()}); err != nil {
			return err
		}
		time.Sleep(time.Second) // Simulate delay between messages
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterExampleServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
