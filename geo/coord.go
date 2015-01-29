package geo

import "math"

type Coord struct {
    x, y float64
}

func (c *Coord) AngleTo(o Coord) float64 {
    return math.Atan2(o.y - c.y, o.x - c.x)
}

func (c *Coord) Add(v Vect) {
    c.x += v.Value * math.Cos(v.Theta)
    c.y += v.Value * math.Sin(v.Theta)
}
