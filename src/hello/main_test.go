package main

import "testing"

func Testsqrt(t *testing.T) {
	got := sqrt(3)
	want := 9.0

	if got != want {
		t.Errorf("greetings('teste') \n got: %v \n want: %v", got, want)
	}
}