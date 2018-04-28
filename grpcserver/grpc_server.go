package grpcserver

import (
	"net"

	"golang.org/x/net/context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "github.com/navinds25/grpcGoExpts/eaconnproto"
)

const (
	port = ":50051"
)

type server struct{}

//grpc functions

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello" + in.Name}, nil
}

func (s *server) SendConfig(ctx context.Context, in *pb.Config) (*pb.Ack, error) {
	name := in.Name
	filename := in.Filename
	log.Println(name, filename)
	return &pb.Ack{Valid: "got config "}, nil
}

func GrpcServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterPublisherServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

