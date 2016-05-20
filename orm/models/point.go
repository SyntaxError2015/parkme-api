package models

import "parkme-api/orm/dbmodels"

// Point contains the left-upper and bottom-lower points of a rectangle in which an entire zone is located
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"Y"`
}

// Expand copies the dbmodels.Point to a Point expands all
// the components by fetching them from the database
func (point *Point) Expand(dbPoint dbmodels.Point) {
	point.X = dbPoint.X
	point.Y = dbPoint.Y
}

// Collapse coppies the Point to a dbmodels.Point and
// only keeps the unique identifiers from the inner components
func (point *Point) Collapse() *dbmodels.Point {
	dbPoint := dbmodels.Point{
		X: point.X,
		Y: point.Y,
	}

	return &dbPoint
}
