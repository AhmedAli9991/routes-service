package repositories

import (
	"context"
	"routes-service/internal/models"

	"gorm.io/gorm"
)

type RouteRepository interface {
	CreateRoute(ctx context.Context, route *models.Route) error
	GetRouteByID(ctx context.Context, id uint) (*models.Route, error)
	UpdateRoute(ctx context.Context, route *models.Route) error
	DeleteRoute(ctx context.Context, id uint) error
}

type routeRepository struct {
	db *gorm.DB
}

func NewRouteRepository(db *gorm.DB) RouteRepository {
	return &routeRepository{db: db}
}

func (r *routeRepository) CreateRoute(ctx context.Context, route *models.Route) error {
	return r.db.WithContext(ctx).Create(route).Error
}

func (r *routeRepository) GetRouteByID(ctx context.Context, id uint) (*models.Route, error) {
	var route models.Route
	if err := r.db.WithContext(ctx).First(&route, id).Error; err != nil {
		return nil, err
	}
	return &route, nil
}

func (r *routeRepository) UpdateRoute(ctx context.Context, route *models.Route) error {
	return r.db.WithContext(ctx).Save(route).Error
}

func (r *routeRepository) DeleteRoute(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Route{}, id).Error
}
