package main

import "testing"


func TestRemoveDuplicates(t *testing.T){
	arr_original := []int{7, 3, 5, 7, 3, 11}
	arr_fact := []int{7, 3, 5, 11}
	
	got := removeDuplicates(arr_original)
	want := arr_fact

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}