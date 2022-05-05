package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("Then result must be 5")
	}
}
func TestSub(t *testing.T) {
	result := sub(3, 2)

	if result != 1 {
		t.Error("Then result must be 1")
	}
}
func TestTimes(t *testing.T) {
	result := times(2, 3)

	if result != 6 {
		t.Error("Then result must be 6")
	}
}
func TestSunX(t *testing.T) {
	result := sumX(2, 3)

	if result != 7 {
		t.Error("Then result must be 7")
	}
}
func TestSunY(t *testing.T) {
	result := sumY(3, 1)

	if result != 5 {
		t.Error("Then result must be 5")
	}
}
func TestSunt(t *testing.T) {
	result := sumt(3, 1)

	if result != 5 {
		t.Error("Then result must be 2")
	}
}
