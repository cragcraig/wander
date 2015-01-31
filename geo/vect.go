package geo

import "math"

type Vect struct {
    Value, Theta float64
}

func (v *Vect) Add(o Vect) {
    dx := v.Value * math.Cos(v.Theta) + o.Value * math.Cos(o.Theta)
    dy := v.Value * math.Sin(v.Theta) + o.Value * math.Sin(o.Theta)
    v.Value = math.Sqrt(dx * dx + dy * dy)
    v.Theta = math.Atan2(dy, dx)
}
