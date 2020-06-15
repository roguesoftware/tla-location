package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/roguesoftware/tla-proto"
)

const port = ":50505"

var initialLocations []*pb.LocationItem

type server struct {
	pb.UnimplementedLocationServiceServer
}

func (s *server) GetLocations(ctx context.Context, in *pb.LocationRequest) (*pb.LocationReply, error) {
	lon := in.GetLongitude()
	lat := in.GetLatitude()
	rad := in.GetRadius()

	log.Printf("Received: %v %v %v", lon, lat, rad)

	var locations []*pb.LocationItem
	locations = initialLocations[1:2]

	return &pb.LocationReply{Locations: locations}, nil
}

func main() {
	// load initial locations
	fileName := "locations.json"
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening %v: %v", fileName, err)
	}
	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(jsonBytes, &initialLocations)

	// create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLocationServiceServer(s, &server{})
	log.Printf("Registered location server with %v stories", len(initialLocations))
	if err := s.Serve(lis); err != nil {
		log.Fatal(s.Serve(lis))
	}
}
