package controllers

import (
	"context"
	"routes-service/internal/handlers"
	pb "routes-service/pkg/proto"
)

type GRPCHandler struct {
	pb.UnimplementedRoutesServiceServer
	controller *handlers.RouteController
}

func NewGRPCHandler(controller *handlers.RouteController) *GRPCHandler {
	return &GRPCHandler{controller: controller}
}

func (h *GRPCHandler) CreateRoute(ctx context.Context, req *pb.Route) (*pb.RouteResponse, error) {
	err := h.controller.CreateRoute(ctx, req.UserId, req.Description, req.Geojson)
	if err != nil {
		return nil, err
	}
	return &pb.RouteResponse{Message: "Route created successfully"}, nil
}

func (h *GRPCHandler) GetRoute(ctx context.Context, req *pb.RouteID) (*pb.Route, error) {
	route, err := h.controller.GetRouteByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Route{
		Id:          int64(route.ID),
		UserId:      route.UserID,
		Description: route.Description,
		Geojson:     route.Geom,
	}, nil
}

func (h *GRPCHandler) UpdateRoute(ctx context.Context, req *pb.Route) (*pb.RouteResponse, error) {
	err := h.controller.UpdateRoute(ctx, uint(req.Id), req.Description, req.Geojson)
	if err != nil {
		return nil, err
	}
	return &pb.RouteResponse{Message: "Route updated successfully"}, nil
}

func (h *GRPCHandler) DeleteRoute(ctx context.Context, req *pb.RouteID) (*pb.RouteResponse, error) {
	err := h.controller.DeleteRoute(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.RouteResponse{Message: "Route deleted successfully"}, nil
}
