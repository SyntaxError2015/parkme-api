package dbmodels

// Point contains the left-upper and bottom-lower points of a rectangle in which an entire zone is located
type Point struct {
	X float64 `bson:"x,omitempty" json:"x"`
	Y float64 `bson:"y,omitempty" json:"y"`
}

// Equal compares two Point objects. Implements the Objecter interface
func (point Point) Equal(obj Objecter) bool {
	otherPoint, ok := obj.(Point)
	if !ok {
		return false
	}

	switch {
	case point.X != otherPoint.X:
		return false
	case point.Y != otherPoint.Y:
		return false
	}

	return true
}
