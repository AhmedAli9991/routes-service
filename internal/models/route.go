package models

type Route struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      string `gorm:"type:text"`
	Description string `gorm:"type:text"`
	Geom        string `gorm:"type:geometry(LineString,4326)"` // Store as GeoJSON string
}
