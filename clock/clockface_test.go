package clockface_test

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"

	"github.com/interwolf/go-tests/clockface"
)

// Svg is for XML parsing of svg
type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	centerX          = 150
	centerY          = 150
)

func TestSVGWriterHand(t *testing.T) {
	testCases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{centerX, centerY, centerX, centerY - secondHandLength}},
		{simpleTime(0, 0, 30), Line{centerX, centerY, centerX, centerY + secondHandLength}},
		{simpleTime(0, 0, 45), Line{centerX, centerY, centerX - secondHandLength, centerY}},
		{simpleTime(0, 0, 15), Line{centerX, centerY, centerX + secondHandLength, centerY}},
		{simpleTime(0, 0, 0), Line{centerX, centerY, centerX, centerY - minuteHandLength}},
		{simpleTime(0, 30, 0), Line{centerX, centerY, centerX, centerY + minuteHandLength}},
		{simpleTime(0, 45, 0), Line{centerX, centerY, centerX - minuteHandLength, centerY}},
		{simpleTime(0, 15, 0), Line{centerX, centerY, centerX + minuteHandLength, centerY}},
		{simpleTime(0, 0, 0), Line{centerX, centerY, centerX, centerY - hourHandLength}},
		{simpleTime(6, 0, 0), Line{centerX, centerY, centerX, centerY + hourHandLength}},
		{simpleTime(9, 0, 0), Line{centerX, centerY, centerX - hourHandLength, centerY}},
		{simpleTime(15, 0, 0), Line{centerX, centerY, centerX + hourHandLength, centerY}},
	}

	for _, test := range testCases {
		t.Run(testName(test.time), func(t *testing.T) {
			time := test.time
			line := test.line
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(line, svg.Line) {
				t.Errorf("Expected to find hand line (%v) in svg lines (%v), but not!\n",
					line, svg.Line)
			}
		})
	}
}

func containsLine(line Line, lines []Line) bool {
	for _, svgLine := range lines {
		if line == svgLine {
			return true
		}
	}

	return false
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(2020, time.January, 1, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("00:00:00")
}

func TestSecondUnit(t *testing.T) {
	testCases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{simpleTime(0, 0, 0), clockface.Point{0, 1}},
		{simpleTime(0, 0, 30), clockface.Point{0, -1}},
		{simpleTime(0, 0, 45), clockface.Point{-1, 0}},
		{simpleTime(0, 0, 15), clockface.Point{1, 0}},
	}

	for _, test := range testCases {
		t.Run(testName(test.time), func(t *testing.T) {
			want := test.point
			got := clockface.SecondHandUnit(test.time)
			if !equalPoint(want, got) {
				t.Fatalf("got: %v, want:%v\n", got, want)
			}
		})
	}
}

func equalFloat64(a, b float64) bool {
	const threshold = 1e-7
	return math.Abs(a-b) < threshold
}

func equalPoint(a, b clockface.Point) bool {
	return equalFloat64(a.X, b.X) && equalFloat64(a.Y, b.Y)
}
