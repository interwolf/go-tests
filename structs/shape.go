package structs

// Shape has an method of Area()
type Shape interface {
	Area() float64
}

// Rect represents a rectangle with width a and height b
type Rect struct {
	a float64
	b float64
}

// Perimeter calc the perimeter of rect(a, b)
func (r Rect) Perimeter() float64 {
	return 2 * (r.a + r.b)
}

// Area calc the area of rect(a, b)
func (r Rect) Area() float64 {
	return r.a * r.b
}

// Circle has a radius
type Circle struct {
	radius float64
}

// Area calc the area of circle(radius)
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Triagle has a base and a height
type Triagle struct {
	base   float64
	height float64
}

// Area calc the area of circle(radius)
func (t Triagle) Area() float64 {
	return t.base * t.height / 2
}
