package geo

import "math"

type BoundingBox struct {
    Coord Coord
    Vertices []vect
}

func (box *BoundingBox) Rotate(theta float64) {
    // TODO
}

func (box *BoundingBox) computeNormals() []vect {
    normals := make([]vect, len(box.Vertices))
    for i := range box.Vertices {
        v1 := box.Vertices[i]
        v2 := box.Vertices[(i + 1) % len(box.Vertices)]
        edge := subtract(v1, v2)
        normals[i] = normalize(perpendicular(edge))
    }
    return normals
}

func (box *BoundingBox) project(axis vect) projection {
    min := dot(axis, box.Vertices[0])
    max := min
    for i := range box.Vertices {
        // The axis must be normalized to get accurate projections
        p := dot(axis, box.Vertices[i])
        if p < min {
            min = p
        } else if p > max {
            max = p
        }
    }
    return projection{min, max}
}

type vect struct {
    x, y float64
}

func subtract(a vect, b vect) vect {
    return vect{a.x - b.x, a.y - b.y}
}

func dot(a vect, b vect) float64 {
    return a.x * b.x + a.y * b.y
}

func normalize(v vect) vect {
    l := math.Sqrt(v.x * v.x + v.y * v.y)
    return vect{v.x / l, v.y / l}
}

func perpendicular(v vect) vect {
    // Could be either (x, y) -> (-y, x) or (y, -x)
    return vect{-v.y, v.x}
}

type projection struct {
    min, max float64
}

func (p *projection) checkOverlap(o *projection) bool {
    return p.min < o.max && p.max > o.min
}

func (p *projection) getOverlap(o *projection) float64 {
    return math.Min(p.max, o.max) - math.Max(p.min, o.min)
}

type Overlap struct {
    Amount float64
    Axis vect
}

// Implements the Separating Axis Theorm (see http://www.codezealot.org/archives/55)
func (box *BoundingBox) CheckCollision(other *BoundingBox) *Overlap {
    overlap := Overlap{Amount: math.MaxFloat64}
    axes := append(box.computeNormals(), other.computeNormals()...)
    for i := range axes {
        axis := axes[i]
        proj1 := box.project(axis)
        proj2 := other.project(axis)
        if !proj1.checkOverlap(&proj2) {
            return nil
        } else {
            // Keep track of the minimum magnitude and axis of overlap.
            o := proj1.getOverlap(&proj2)
            if (o < overlap.Amount) {
                overlap.Amount = o
                overlap.Axis = axis
            }
        }
    }
    return &overlap
}

