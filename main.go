package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/roguesoftware/tla-proto"
)

const (
	port = ":50505"
)

type server struct {
	pb.UnimplementedLocationServiceServer
}

func (s *server) GetLocations(ctx context.Context, in *pb.LocationRequest) (*pb.LocationReply, error) {
	name := in.GetName()
	// lon := in.GetLongitude()
	// lat := in.GetLatitude()
	// rad := in.GetRadius()

	// log.Printf("Received: %v %v %v", lon, lat, rad)
	log.Printf("Received: %v", name)

	// var locations []pb.Location
	var location pb.LocationItem

	location.Id = "abcdef"
	// location.Longitude = -112.5
	// location.Latitude = 32.3
	// location.Address = "123 Main St, Leander, TX 78641"
	// location.Title = "City Hall"
	// location.Distance = 10468.4

	// locations = append(locations, location)

	return &pb.LocationReply{Message: &location}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLocationServiceServer(s, &server{})
	log.Printf("Registered location server")
	if err := s.Serve(lis); err != nil {
		log.Fatal(s.Serve(lis))
	}
}
