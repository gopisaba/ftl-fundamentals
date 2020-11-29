package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 3, want: 5},
		{a: 4, b: 3, want: 7},
	}
	for i, tc := range testCases {
		t.Logf("Test case %d \n", i)
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 5, b: 3, want: 2},
		{a: 6, b: 3, want: 3},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 5, b: 3, want: 15},
		{a: 6, b: 3, want: 18},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type divideTestCase struct {
		a, b        float64
		want        float64
		errExpected bool
	}
	testCases := []divideTestCase{
		{a: 5, b: 2.5, want: 2, errExpected: false},
		{a: 5, b: 0, want: 0, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if (tc.errExpected && err == nil) || (!tc.errExpected && err != nil) {
			t.Error("Unexpected Errors")
		}
		if tc.want != got {
			t.Errorf("want %f, got %f, err %s", tc.want, got, err)
		}
	}
}
