package services

import (
	"context"
	"routes-service/internal/models"
	"routes-service/internal/repositories"
)

type RouteService interface {
	CreateRoute(ctx context.Context, route *models.Route) error
	GetRouteByID(ctx context.Context, id uint) (*models.Route, error)
	UpdateRoute(ctx context.Context, route *models.Route) error
	DeleteRoute(ctx context.Context, id uint) error
}

type routeService struct {
	repo repositories.RouteRepository
}

func NewRouteService(repo repositories.RouteRepository) RouteService {
	return &routeService{repo: repo}
}

func (s *routeService) CreateRoute(ctx context.Context, route *models.Route) error {
	return s.repo.CreateRoute(ctx, route)
}

func (s *routeService) GetRouteByID(ctx context.Context, id uint) (*models.Route, error) {
	return s.repo.GetRouteByID(ctx, id)
}

func (s *routeService) UpdateRoute(ctx context.Context, route *models.Route) error {
	return s.repo.UpdateRoute(ctx, route)
}

func (s *routeService) DeleteRoute(ctx context.Context, id uint) error {
	return s.repo.DeleteRoute(ctx, id)
}
