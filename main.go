package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/7yrionLannister/golang-protobuf-tutorial/generated/tutorialpb"
	"google.golang.org/protobuf/proto"

	pb "github.com/7yrionLannister/golang-protobuf-tutorial/generated2/twirptutorial"
)

const fname = "myproto.binary"

func main() {
	book := &tutorialpb.AddressBook{
		People: []*tutorialpb.Person{
			{
				Name:  "Michael",
				Email: "michael@example.com",
				Phones: []*tutorialpb.Person_PhoneNumber{
					{
						Number: "+57 3123456789",
						Type:   tutorialpb.PhoneType_PHONE_TYPE_MOBILE,
					},
					{
						Number: "(606) 31232144",
						Type:   tutorialpb.PhoneType_PHONE_TYPE_HOME,
					},
				},
			},
		},
	}

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := os.WriteFile(fname, out, 0o644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the generated address book binary file.
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	retrievedBook := &tutorialpb.AddressBook{}
	if err := proto.Unmarshal(in, retrievedBook); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Printf("Retrieved book from binary file: %+v\n", retrievedBook)
	fmt.Printf("Original book: %+v\n", book)

	twirpExample()
}

// --------------------------------------------------------------------------------------
// twirp example: No entiendas el código autogenerado, concentrate en esto

// Implementation of pb.HelloWorld
type HelloWorldServer struct{}

// Receives and handles the RPC request through HTTP
func (s *HelloWorldServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	fmt.Println("Server received RPC request from client and It's processing it")
	return &pb.HelloResp{Text: "Hello " + req.Subject}, nil
}

func twirpExample() {
	// CLIENT
	// Because we are creating a Protobuf client, the Content-Type of all requests is going to be "application/protobuf"
	client := pb.NewHelloWorldProtobufClient("http://localhost:8080", &http.Client{})
	timer := time.NewTimer(time.Second * 10)
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for {
			select {
			case <-ticker.C:
				// Sends RPC HTTP request using "application/protobuf" as Content Type
				resp, err := client.Hello(context.Background(), &pb.HelloReq{Subject: "World"})
				if err == nil {
					fmt.Println("Client received response:", resp.Text) // prints "Hello World"
				} else {
					fmt.Println("Error:", err)
				}
			case <-timer.C:
				return
			}
		}
	}()

	// SERVER
	twirpHandler := pb.NewHelloWorldServer(&HelloWorldServer{})
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	http.ListenAndServe(":8080", mux)
}
