package main

import (
	"log"
	"net"
	"routes-service/internal/controllers"
	"routes-service/internal/database"
	"routes-service/internal/handlers"
	"routes-service/internal/repositories"
	"routes-service/internal/services"
	pb "routes-service/pkg/proto"

	"google.golang.org/grpc"
)

func main() {
	database.InitDB()

	// Setup repository, service, and controller
	routeRepo := repositories.NewRouteRepository(database.DB)
	routeService := services.NewRouteService(routeRepo)
	routeController := handlers.NewRouteController(routeService)

	// Setup gRPC handler
	grpcHandler := controllers.NewGRPCHandler(routeController)

	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRoutesServiceServer(grpcServer, grpcHandler)

	log.Println("Server is running on port 50050")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
