package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, v := range cases {
		t.Run(testName(v.time), func(t *testing.T) {
			got := secondsInRadians(v.time)
			if got != v.angle {
				t.Fatalf("wanted %v radians, but got %v", v.angle, got)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, v := range cases {
		t.Run(testName(v.time), func(t *testing.T) {
			got := secondHandPoint(v.time)

			if !roughlyEqualPoint(got, v.point) {
				t.Fatalf("wanted %v Point, but got %v", v.point, got)
			}
		})
	}
}

func TestMinutesInRadiant(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{
			simpleTime(0, 30, 0),
			math.Pi,
		},
	}

	for _, v := range cases {
		t.Run(testName(v.time), func(t *testing.T) {
			got := minutesInRadians(v.time)

			if got != v.angle {
				t.Fatalf("Wanted %v radians, but got %v", v.angle, got)
			}
		})
	}
}

func roughlyEqualFloat(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat(a.X, b.X) &&
		roughlyEqualFloat(a.Y, b.Y)
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
