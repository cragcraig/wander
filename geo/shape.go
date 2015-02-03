package geo

type Shape interface {
	Project(axis Vect) Projection
	Rotate(heading Vect) Shape
	Translate(dist Vect) Shape
	CheckCollision(other Shape) *Overlap
	// GetClosestPoint(pos Vect)
	// GetCenter(pos Vect)
}
