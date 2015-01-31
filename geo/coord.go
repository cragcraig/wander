package geo

import "math"

type Coord struct {
    X, Y float64
}

func (c *Coord) AngleTo(o Coord) float64 {
    return math.Atan2(o.Y - c.Y, o.X - c.X)
}

func (c *Coord) Add(v Vect) {
    c.X += v.Value * math.Cos(v.Theta)
    c.Y += v.Value * math.Sin(v.Theta)
}
