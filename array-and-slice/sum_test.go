package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		expected := 10

		if got != expected {
			t.Errorf("given, %v, expected %d but got %d", numbers, expected, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{0, 9}
	got := SumAll(slice1, slice2)
	expected := []int{3, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("given slices, %v, and %v, expected %v but got %v", slice1, slice2, expected, got)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}
		got := SumAllTails(slice1, slice2)
		expected := []int{2, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("given slices, %v, and %v, expected %v but got %v", slice1, slice2, expected, got)
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{0, 9}
		got := SumAllTails(slice1, slice2)
		expected := []int{0, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("given slices, %v, and %v, expected %v but got %v", slice1, slice2, expected, got)
		}
	})
}
