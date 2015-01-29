package tactical

type CollidableType int

const (
    SHIP = iota
    REEF
    LAND
)

type Collidable interface {
    getType() CollidableType
    getBoundingBox() BoundingBox
}
