package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Lincyaw/SimpleOtelCollector/proto/otlp/collector/profiles/v1experimental"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProfilesServiceServer
}

func (s *server) Export(ctx context.Context, in *pb.ExportProfilesServiceRequest) (*pb.ExportProfilesServiceResponse, error) {
	for _, p := range in.ResourceProfiles {
		fmt.Printf("%+v", p)
	}
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProfilesServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
