# Go GRPC tutorial 
## Setup project

Step 1.

To make go mod file run in terminal:

``go mod init FolderName ``
``(e.g: go mod init DISYS/ChittyChat)``

Don't be confused by I have something called a go.sum. It gets made at some point lol

___

Step 2.

Paste into mod:

```
require (
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20211004164453-cedda3a722dd // indirect
	golang.org/x/sys v0.0.0-20211004093028-2c5d950f24ef // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211001223012-bfb93cce50d9 // indirect

)
```

OR: Nikoline said to use in terminal ``go get google.golang.org/grpc`` but when I tried it only 
inserted some of it.

___
Step 3.

Create a Proto file in new proto folder.
I always name the proto file **proto-file.proto** because then the following terminal command
just works.

For very simple proto file see it in this repo.
___

Step 4.

Then in the parent folder of the proto folder paste this into the terminal to auto generate the GRPC go files.
``protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     proto/proto-file.proto``

___

Then run (After you have made proto files, otherwise it just sees that dependencies
are not used and deletes them):

```go mod tidy```

## Make A Server
Step 1. Make a folder called server in the root folder 

Step 2. In this folder create a main.go file.
(In my IDE it puts package server at the very top. To follow my tutorial change this to package main)

Step 3. There are some code you just have to write to set up the server. See in server/main.go.
(There are many ways to tweak this but if you have never done this before just follow along)
Be very aware of the naming if you have changed the Service name in the proto. If you did you will have t put in your services name 
into :

(CoolName._something_ here is something I have made to refer to the proto files. See in server/main.go to understand why this is)

the server:
````
type server struct {
	CoolName.Unimplemented[YOURCOOLNAME]Server 
}
````

the main method:

```` 
CoolName.Register[YOURCOOLNAME]Server(s, &server{}) //Has to have the name of your service (Service for us)

````

___

Step 4.
Make the methods that you defined in the proto file that the service would have. (If you
look in my example it is SayHelloName and SayHelloUser

___

**Little Explanation here:** 

So we make two methods by these names in server so that the client can call them.

We need to match the signature that the autogenerated files dictate so that means:
- The method needs to know that it is connected to the server.
We tell it that by writing (s *server) in front of the method name. 
This refers to the s made in main of the type server (which is the struct we made)
- It has to take a context and a request
- It has to output a CoolName Response and an error (If no error occurred this variable will simply be nil)

There is a log.Printf which is just there to make the server print something out so that we can see what it recived

it then returns a response where the field Hello is set to the string "Hello" and the YourName field is set
to the name to got as input.

These fields are what we defined in the proto file that a response should have 

````
func (s *server) SayHelloName(ctx context.Context, in *CoolName.Request) (*CoolName.Response, error) {
	log.Printf("Received: %v", in.GetName()) 
	return &CoolName.Response{Hello: "Hello", YourName: in.Name}, nil //the nil here is saying no error happned
}
````

## Make a Client
Step 1. Make a folder called client in the root folder 

Step 2. In this folder create a main.go file.
(In my IDE it puts package server at the very top. To follow my tutorial change this to package main)

Step 3. There are many different ways to make a client. Here I give an explanation for the one I made.

I made a scanner that takes input from the terminal in a loop that runs forever (Or until you terminate the program)

Depending on the input it gets it sends one of two requests to the server.
These two request are the ones defined in the proto file.

## Run the program
Step 1. Open two terminal windows

Step 2. Start the server first

You can do this by in the terminal cd'ing into it first ```cd .\server\ ``` (I'm on Windows)
and then running it there: ```go run main.go``` or just ``go run .``

Step 3. Start a client (You can start more)
in another terminal cd into the client folder ```cd .\client\ ``` 
and then running it there: ```go run main.go``` or just ``go run .``

Step 4. Write your name in the client terminal! If everything went well the server will say hello to you!
