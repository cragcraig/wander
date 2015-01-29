package geo

type BoundingBox struct {
    Coord Coord
    Theta float64
    Rect [4]Coord
}

func (box *BoundingBox) checkCollision(other *BoundingBox) bool {
    // TODO: Check for collison.
    return false
}
