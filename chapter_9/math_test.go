package math

import "testing"

type testavg struct {
	values  []float64
	average float64
}

type testmax struct {
	values []float64
	max    float64
}

type testmin struct {
	values []float64
	min    float64
}

// slice of testavgs
var avgtests = []testavg{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

var maxtests = []testmax{
	{[]float64{2, 6, 5, 1, 5, 9, 8}, 9},
	{[]float64{1, 1, 1, 1, 1}, 1},
	{[]float64{-1, -9, 8, 11, 1, 2}, 11},
	{[]float64{22, 3, 12, 15, 91}, 91},
	{[]float64{}, 0},
}

var mintests = []testmin{
	{[]float64{2, 6, 5, 1, 5, 9, 8}, 1},
	{[]float64{1, 1, 1, 1, 1}, 1},
	{[]float64{-1, -9, 8, 11, 1, 2}, -9},
	{[]float64{22, 3, 12, 15, 91}, 3},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range avgtests {
		if len(pair.values) == 0 {
			pair.values = []float64{0}
		}
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMaximum(t *testing.T) {
	for _, pair := range maxtests {
		if len(pair.values) == 0 {
			pair.values = []float64{0}
		}
		max := Maximum(pair.values)
		if max != pair.max {
			t.Error(
				"for", pair.values,
				"expected", pair.max,
				"got", pair.max,
			)
		}
	}
}

func TestMinimum(t *testing.T) {
	for _, pair := range mintests {
		if len(pair.values) == 0 {
			pair.values = []float64{0}
		}
		min := Minimum(pair.values)
		if min != pair.min {
			t.Error(
				"for", pair.values,
				"expected", pair.min,
				"got", pair.min,
			)
		}
	}
}
