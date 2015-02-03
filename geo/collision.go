package geo

import "math"

type Projection struct {
	Min, Max float64
}

func (p Projection) CheckOverlap(o Projection) bool {
	return p.Min < o.Max && p.Max > o.Min
}

func (p Projection) GetOverlapAmount(o Projection) float64 {
	return math.Min(p.Max, o.Max) - math.Max(p.Min, o.Min)
}

type Overlap struct {
	Amount  float64
	Heading Vect
}

// Implements the Separating Axis Theorm (see http://www.codezealot.org/archives/55)
// Polygons MUST be convex and axes must be normalized.
func checkCollision(shape1 Shape, shape2 Shape, axes []Vect) *Overlap {
	overlap := Overlap{Amount: math.MaxFloat64}
	for i := range axes {
		axis := axes[i]
		proj1 := shape1.Project(axis)
		proj2 := shape2.Project(axis)
		if !proj1.CheckOverlap(proj2) {
			return nil
		} else {
			// Keep track of the minimum magnitude and axis of overlap.
			o := proj1.GetOverlapAmount(proj2)
			if o < overlap.Amount {
				overlap.Amount = o
				overlap.Heading = axis
			}
		}
	}
	return &overlap
}
