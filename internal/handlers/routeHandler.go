package handlers

import (
	"context"
	"routes-service/internal/models"
	"routes-service/internal/services"
)

type RouteController struct {
	routeService services.RouteService
}

func NewRouteController(routeService services.RouteService) *RouteController {
	return &RouteController{routeService: routeService}
}

func (c *RouteController) CreateRoute(ctx context.Context, userID string, description string, geojson string) error {
	route := &models.Route{
		UserID:      userID,
		Description: description,
		Geom:        geojson, // Convert GeoJSON to PostGIS type if needed
	}
	return c.routeService.CreateRoute(ctx, route)
}

func (c *RouteController) GetRouteByID(ctx context.Context, id uint) (*models.Route, error) {
	return c.routeService.GetRouteByID(ctx, id)
}

func (c *RouteController) UpdateRoute(ctx context.Context, id uint, description string, geojson string) error {
	route := &models.Route{
		ID:          id,
		Description: description,
		Geom:        geojson,
	}
	return c.routeService.UpdateRoute(ctx, route)
}

func (c *RouteController) DeleteRoute(ctx context.Context, id uint) error {
	return c.routeService.DeleteRoute(ctx, id)
}
