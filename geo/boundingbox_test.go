package geo

import "testing"

func TestCheckCollision_True(t *testing.T) {
    b1 := BoundingBox{Coord{0, 0}, []vect{
        vect{0, 0}, vect{0, 1}, vect{1, 1}, vect{1, 0}}}
    b2 := BoundingBox{Coord{0, 0}, []vect{
        vect{0.5, 0.5}, vect{0.5, 1.5}, vect{1.5, 1.5}, vect{1.5, 0.5}}}
    if !b1.CheckCollision(&b2) {
        t.FailNow()
    }
}

func TestCheckCollision_False(t *testing.T) {
    b1 := BoundingBox{Coord{0, 0}, []vect{
        vect{0, 0}, vect{0, 1}, vect{1, 1}, vect{1, 0}}}
    b2 := BoundingBox{Coord{0, 0}, []vect{
        vect{1.5, 1.5}, vect{1.5, 2.5}, vect{2.5, 2.5}, vect{2.5, 1.5}}}
    if b1.CheckCollision(&b2) {
        t.FailNow()
    }
}
