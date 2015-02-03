package geo

import "math"

type Polygon struct {
	Vertices []Vect
}

func (poly Polygon) ComputeNormals() []Vect {
	normals := make([]Vect, len(poly.Vertices))
	for i := range poly.Vertices {
		v1 := poly.Vertices[i]
		v2 := poly.Vertices[(i+1)%len(poly.Vertices)]
		edge := v1.Subtract(v2)
		normals[i] = edge.Perpendicular().Normalized()
	}
	return normals
}

func (poly Polygon) Project(axis Vect) Projection {
	min := axis.Dot(poly.Vertices[0])
	max := min
	for i := range poly.Vertices {
		// The axis must be normalized to get accurate projections
		p := axis.Dot(poly.Vertices[i])
		if p < min {
			min = p
		} else if p > max {
			max = p
		}
	}
	return Projection{min, max}
}

func (poly Polygon) applyTransform(f func(Vect) Vect) Polygon {
	ret := Polygon{make([]Vect, len(poly.Vertices))}
	for i := range poly.Vertices {
		ret.Vertices[i] = f(poly.Vertices[i])
	}
	return ret
}

func (poly Polygon) Rotate(heading Vect) Shape {
	return poly.applyTransform(func(v Vect) Vect {
		return v.Rotate(heading)
	})
}

func (poly Polygon) Translate(dist Vect) Shape {
	return poly.applyTransform(func(v Vect) Vect {
		return v.Add(dist)
	})
}

func (poly Polygon) getClosestVertex(pos Vect) Vect {
	closest := poly.Vertices[0]
	mindist := math.MaxFloat64
	for i := range poly.Vertices {
		v := poly.Vertices[i]
		d := v.X*v.X + v.Y*v.Y
		if d < mindist {
			mindist = d
			closest = v
		}
	}
	return closest
}

func (poly Polygon) CheckCollision(other Shape) *Overlap {
	var axes []Vect
	switch other := other.(type) {
	case Polygon:
		axes = append(poly.ComputeNormals(), other.ComputeNormals()...)
	case Circle:
		axes = append(poly.ComputeNormals(), poly.getClosestVertex(other.Offset).Subtract(other.Offset).Normalized())
	default:
		return nil
	}
	return checkCollision(poly, other, axes)
}
