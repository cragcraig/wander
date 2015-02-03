package geo

import "testing"

func TestPolygon_CheckPolygonCollision_True(t *testing.T) {
	s1 := Polygon{[]Vect{Vect{0, 0}, Vect{0, 1}, Vect{1, 1}, Vect{1, 0}}}
	s2 := Polygon{[]Vect{Vect{0.5, 0.5}, Vect{0.5, 1.5}, Vect{1.5, 1.5}, Vect{1.5, 0.5}}}
	if s1.CheckCollision(s2) == nil {
		t.FailNow()
	}
}

func TestPolygon_CheckPolygonCollision_False(t *testing.T) {
	s1 := Polygon{[]Vect{Vect{0, 0}, Vect{0, 1}, Vect{1, 1}, Vect{1, 0}}}
	s2 := Polygon{[]Vect{Vect{1.5, 1.5}, Vect{1.5, 2.5}, Vect{2.5, 2.5}, Vect{2.5, 1.5}}}
	if s1.CheckCollision(s2) != nil {
		t.FailNow()
	}
}

func TestPolygon_CheckCircleCollision_True(t *testing.T) {
	s1 := Polygon{[]Vect{Vect{0, 0}, Vect{0, 1}, Vect{1, 1}, Vect{1, 0}}}
	s2 := Circle{1.5, Vect{2, 2}}
	if s1.CheckCollision(s2) == nil {
		t.FailNow()
	}
}

func TestPolygon_CheckCircleCollision_False(t *testing.T) {
	s1 := Polygon{[]Vect{Vect{0, 0}, Vect{0, 1}, Vect{1, 1}, Vect{1, 0}}}
	s2 := Circle{0.5, Vect{2, 2}}
	if s1.CheckCollision(s2) != nil {
		t.FailNow()
	}
}
