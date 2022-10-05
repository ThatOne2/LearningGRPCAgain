package main

//Many of these come automatically when the IDE sees that we use them
//The ones that often gives problem is context (My IDE likes to just erase it if I'm not using it -.-´
import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"

	CoolName "LearningGRPCAgain/proto"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //Making our scanner
	fmt.Println("Please write your name and press enter")
	fmt.Println("To end the program pres control and c at the same time (it just kills it)") //You can make a prettier way to end the program
	//But I will not do that here

	for scanner.Scan() { //a loop that will run forever taking inputs
		UserName := scanner.Text()
		UserName = strings.TrimSpace(UserName) //Removes all space characters and newlines from the username
		//So if a user just wrote space and then entered the string will be seen as less than one (og if a user just wrote enter)
		if len(UserName) < 1 {
			//Because the user didn't give us there name we will call the SendHelloUserRequest so our server just will say hello user
			reqUser := CoolName.Empty{} //Making an empty request
			SendHelloUserRequest(&reqUser)
		} else {
			req := CoolName.Request{Name: UserName} //Making a request with the user name
			SendHelloNameRequest(&req)
		}

	}
}

func SendHelloNameRequest(req *CoolName.Request) {
	//We make a connection to the server at port 8080
	//We open up this connection every time we want to send a message
	connection, err := grpc.Dial(fmt.Sprintf("localhost:%d", 8080), grpc.WithInsecure())
	//Defer close is a way to tell the program "when this method is done please close the connection"
	defer connection.Close()

	//error handeling, go likes that
	if err != nil {
		log.Fatal("Failed to dial: %v", err)
	}

	//Get the context background, we need to send this to our server along with our request
	ctx := context.Background()

	//Making a client that has a connection to our server
	client := CoolName.NewServiceClient(connection)

	//Here we asy that we want to save the response we get from calling the server with the SayHelloName method
	response, err := client.SayHelloName(ctx, req) //we send the context and our request
	if err != nil {
		fmt.Println("Request did not go through")
		return
	}

	//We print our response
	fmt.Printf("%s %s\n", response.Hello, response.YourName)
}

//Functions in much the same way as SendHelloNameRequest
//This methoed however takes and *CoolName.Empty request
func SendHelloUserRequest(req *CoolName.Empty) {
	connection, err := grpc.Dial(fmt.Sprintf("localhost:%d", 8080), grpc.WithInsecure())
	defer connection.Close()

	if err != nil {
		log.Fatal("Failed to dial: %v", err)
	}

	ctx := context.Background()
	client := CoolName.NewServiceClient(connection)

	response, err := client.SayHelloUser(ctx, req) //We send the method call SayHelloUser here
	if err != nil {
		fmt.Println("Request did not go through")
		return
	}

	fmt.Printf("%s %s\n", response.Hello, response.YourName)
}
