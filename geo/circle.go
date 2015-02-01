package geo

type Circle struct {
    Radius float64
    Offset Vect
}

func (cir Circle) Project(axis Vect) Projection {
    p := axis.Dot(cir.Offset)
    return Projection{p - cir.Radius, p + cir.Radius}
}

func (cir Circle) Rotate(heading Vect) Shape {
    // Rotating a circle is a no-op
    return cir
}

func (cir Circle) Translate(dist Vect) Shape {
    return Circle{cir.Radius, cir.Offset.Add(dist)}
}

func (cir Circle) checkCircleCollision(other Circle) *Overlap {
    offset := other.Offset.Subtract(cir.Offset)
    offset_mag := offset.Magnitude()
    min_dist := cir.Radius + other.Radius
    if offset_mag < min_dist {
        return &Overlap{min_dist - offset_mag, offset.Normalized()}
    }
    return nil
}

func (cir Circle) CheckCollision(other Shape) *Overlap {
    switch other := other.(type) {
    case Polygon:
        // Circle/Polygon collisions are already implemented on Polygon
        return other.CheckCollision(cir)
    case Circle:
        return cir.checkCircleCollision(other)
    }
    return nil
}
