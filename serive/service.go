package main

import (
	"context"
	pb "go-grpc-example/proto/search"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SearchService struct {
	
}

const PORT = "9001"

func (s *SearchService)Search(ctx context.Context,r *pb.SearchRequest)(*pb.SearchResponse,error)  {
	request := r.GetRequest()

	return &pb.SearchResponse{Response: "入参为:"+request+"server"},nil
}

func (s *SearchService)List(ctx context.Context,r *pb.ListRequest)(*pb.ListResponse,error)  {
	return &pb.ListResponse{
		Status: true,
		Code:   200,
		Msg:    "list",
	}, nil
}

func main()  {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis,err := net.Listen("tcp",":"+PORT)
	if err != nil{
		log.Fatalf("net.Listen err: %v", err)
	}
	err = server.Serve(lis)
	if err != nil {
		return 
	}
}