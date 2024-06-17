package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Skele878/gRPC_CRUD_App_Example/db"
	pb "github.com/Skele878/gRPC_CRUD_App_Example/proto"
	"github.com/Skele878/gRPC_CRUD_App_Example/server"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	Port      = flag.Int("port", 50051, "gRPC server port")
	srv       = &server.Server{}
	newServer = grpc.NewServer()
)
var DB *gorm.DB

func init() {
	db.DatabaseConnection()
}

func main() {
	fmt.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *Port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pb.RegisterMovieServiceServer(newServer, srv)

	log.Printf("Server listening at %v", lis.Addr())

	if err := newServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
