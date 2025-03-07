package main

import (
	"fmt"
	"log"
	"os"

	"github.com/7yrionLannister/golang-protobuf-tutorial/generated/tutorialpb"
	"google.golang.org/protobuf/proto"
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
}
