package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	//We need to teach it where to find the auto generated files so we do that here and give it a name we can remember
	//The name I choose is very cool thank you
	CoolName "LearningGRPCAgain/proto"
)

// server is used to implement Service from the proto.
type server struct {
	CoolName.UnimplementedServiceServer //Has to have the name of your service (Service for us)
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080)) //the port here is not too important

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() //make a new grpc server

	//Connect the made server to the proto generated Service
	CoolName.RegisterServiceServer(s, &server{}) //Has to have the name of your service (Service for us)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil { //Start serving (It will start listening for inputs from clients)
		log.Fatalf("failed to serve: %v", err)
	}
}

//making the say hello name function we defined in the proto file (It just says hello and takes the name it got as an input)
func (s *server) SayHelloName(ctx context.Context, in *CoolName.Request) (*CoolName.Response, error) {
	log.Printf("Received: %v", in.GetName())
	return &CoolName.Response{Hello: "Hello", YourName: in.Name}, nil //the nil here is saying no error happned
}

//making the say hello name function we defined in the proto file (It just says hello user)
func (s *server) SayHelloUser(ctx context.Context, in *CoolName.Empty) (*CoolName.Response, error) {
	log.Printf("Received Empty message")
	return &CoolName.Response{Hello: "Hello", YourName: "User"}, nil //the nil here is saying no error happned
}
