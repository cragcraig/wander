package geo

import "testing"

func TestCircle_CheckCircleCollision_True(t *testing.T) {
	s1 := Circle{1, Vect{1, 1}}
	s2 := Circle{1, Vect{2, 2}}
	if s1.CheckCollision(s2) == nil {
		t.FailNow()
	}
}

func TestCircle_CheckCircleCollision_False(t *testing.T) {
	s1 := Circle{1, Vect{1, 1}}
	s2 := Circle{1, Vect{3, 3}}
	if s1.CheckCollision(s2) != nil {
		t.FailNow()
	}
}

func TestCircle_CheckPolygonCollision_True(t *testing.T) {
	s1 := Circle{1.5, Vect{2, 2}}
	s2 := Polygon{[]Vect{Vect{0, 0}, Vect{0, 1}, Vect{1, 1}, Vect{1, 0}}}
	if s1.CheckCollision(s2) == nil {
		t.FailNow()
	}
}

func TestCircle_CheckPolygonCollision_False(t *testing.T) {
	s1 := Circle{0.5, Vect{2, 2}}
	s2 := Polygon{[]Vect{Vect{0, 0}, Vect{0, 1}, Vect{1, 1}, Vect{1, 0}}}
	if s1.CheckCollision(s2) != nil {
		t.FailNow()
	}
}
