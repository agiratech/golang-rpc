package main

import (
  "log"
  "net"
  "strings"

  "golang.org/x/net/context"
  "google.golang.org/grpc"

  pb "github.com/agiratech/golang-rpc/person"
)

const (
  port = ":3333"
)

// server is used to implement person.PersonServer.
type server struct {
  savedPersons []*pb.PersonRequest
}

// CreatePerson creates a new Person
func (s *server) CreatePerson(ctx context.Context, in *pb.PersonRequest) (*pb.PersonResponse, error) {
  s.savedPersons = append(s.savedPersons, in)
  return &pb.PersonResponse{Id: in.Id, Success: true}, nil
}

// GetPersons returns all persons by given filter
func (s *server) GetPersons(filter *pb.PersonFilter, stream pb.Person_GetPersonsServer) error {
  for _, person := range s.savedPersons {
    if filter.Keyword != "" {
      if !strings.Contains(person.Name, filter.Keyword) {
        continue
      }
    }
    if err := stream.Send(person); err != nil {
      return err
    }
  }
  return nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  // Creates a new gRPC server
  s := grpc.NewServer()
  pb.RegisterPersonServer(s, &server{})
  s.Serve(lis)
}