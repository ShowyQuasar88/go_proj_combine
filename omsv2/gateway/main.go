package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/showyquasar88/proj-combine/omsv2/common"
	pb "github.com/showyquasar88/proj-combine/omsv2/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var (
	HttpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	OrderServiceAddr = common.EnvString("ORDER_SERVICE_ADDR", "localhost:9001")
)

func main() {
	conn, err := grpc.NewClient(OrderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create grpc connection: %v", err)
	}
	defer conn.Close()

	log.Printf("Dialing order service at %s", OrderServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", HttpAddr)

	if err := http.ListenAndServe(HttpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
