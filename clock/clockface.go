package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

// Point represents a two-dimensional coordinate
type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	centerX          = 150
	centerY          = 150
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// SVGWriter writes an svg representation to the writer w
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	io.WriteString(w, secondHandTag(t))
	io.WriteString(w, minuteHandTag(t))
	io.WriteString(w, hourHandTag(t))
	io.WriteString(w, svgEnd)
}

func secondHandTag(t time.Time) string {
	p := SecondHand(t)
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHandTag(t time.Time) string {
	p := MinuteHand(t)
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func hourHandTag(t time.Time) string {
	p := HourHand(t)
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

// SecondHand is the point of the second hand on clock
func SecondHand(t time.Time) Point {

	p := SecondHandUnit(t)
	p = Point{p.X * secondHandLength, -p.Y * secondHandLength}
	p = Point{p.X + centerX, p.Y + centerY}

	return p
}

// MinuteHand is the point of the second hand on clock
func MinuteHand(t time.Time) Point {

	p := minuteHandUnit(t)
	p = Point{p.X * minuteHandLength, -p.Y * minuteHandLength}
	p = Point{p.X + centerX, p.Y + centerY}

	return p
}

// HourHand is the point of the second hand on clock
func HourHand(t time.Time) Point {

	p := hourHandUnit(t)
	p = Point{p.X * hourHandLength, -p.Y * hourHandLength}
	p = Point{p.X + centerX, p.Y + centerY}

	return p
}

// SecondHandUnit calcs the unit point
func SecondHandUnit(t time.Time) Point {
	angle := SecondInRadian(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

// minuteHandUnit calcs the unit point
func minuteHandUnit(t time.Time) Point {
	angle := MinuteInRadian(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

// hourHandUnit calcs the unit point
func hourHandUnit(t time.Time) Point {
	angle := HourInRadian(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

// SecondInRadian calculates the radian
func SecondInRadian(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

// MinuteInRadian calculates the radian
func MinuteInRadian(t time.Time) float64 {
	return (math.Pi / (minutesInHalfClock / (float64(t.Minute()))))
}

// HourInRadian calculates the radian
func HourInRadian(t time.Time) float64 {
	return (math.Pi / (hoursInHalfClock / (float64(t.Hour() % 12)))) +
		(MinuteInRadian(t) / 12)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
