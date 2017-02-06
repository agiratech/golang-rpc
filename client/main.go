package main

import (
  "io"
  "log"

  "golang.org/x/net/context"
  "google.golang.org/grpc"

  pb "github.com/agiratech/golang-rpc/person"
)

const (
  address = "localhost:3333"
)

// createPerson calls the RPC method CreatePerson of PersonServer
func createPerson(client pb.PersonClient, person *pb.PersonRequest) {
  resp, err := client.CreatePerson(context.Background(), person)
  if err != nil {
    log.Fatalf("Could not create Person: %v", err)
  }
  if resp.Success {
    log.Printf("A new Person has been added with id: %d", resp.Id)
  }
}

// getPersons calls the RPC method GetPersons of PersonServer
func getPersons(client pb.PersonClient, filter *pb.PersonFilter) {
  // calling the streaming API
  stream, err := client.GetPersons(context.Background(), filter)
  if err != nil {
    log.Fatalf("Error on get persons: %v", err)
  }
  for {
    person, err := stream.Recv()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatalf("%v.GetPersons(_) = _, %v", client, err)
    }
    log.Printf("Person: %v", person)
  }
}
func main() {
  // Set up a connection to the gRPC server.
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  // Creates a new PersonClient
  client := pb.NewPersonClient(conn)

  person := &pb.PersonRequest{
    Id:    1001,
    Name:  "Reddy",
    Email: "reddy@xyz.com",
    Phone: "9898982929",
    Addresses: []*pb.PersonRequest_Address{
      &pb.PersonRequest_Address{
        Street:            "Tripilcane",
        City:              "Chennai",
        State:             "TN",
        Zip:               "600002",
        IsShippingAddress: false,
      },
      &pb.PersonRequest_Address{
        Street:            "Balaji colony",
        City:              "Tirupati",
        State:             "AP",
        Zip:               "517501",
        IsShippingAddress: true,
      },
    },
  }

  // Create a new person
  createPerson(client, person)

  person = &pb.PersonRequest{
    Id:    1002,
    Name:  "Raj",
    Email: "raj@xyz.com",
    Phone: "5000510001",
    Addresses: []*pb.PersonRequest_Address{
      &pb.PersonRequest_Address{
        Street:            "Marathahalli",
        City:              "Bangalore",
        State:             "KS",
        Zip:               "560037",
        IsShippingAddress: true,
      },
    },
  }

  // Create a new person
  createPerson(client, person)
  // Filter with an empty Keyword
  filter := &pb.PersonFilter{Keyword: ""}
  getPersons(client, filter)
}