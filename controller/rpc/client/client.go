package main
import (
	"context"
	pb "go-grpc/controller/rpc/proto/search"
	"google.golang.org/grpc"
	"log"
)
const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())

	resp1, err := client.List(context.Background(), &pb.ListRequest{
		ProjectCode: "gRPC",
		ProductId: 100,
	})

	if err != nil {
		log.Fatalf("client.List err: %v", err)
	}

	log.Printf("resp: %v", resp1.GetMsg())



}