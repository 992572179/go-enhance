package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	tests := []struct {
		a, b, c int
	}{
		{2, 5, 10},
		{11, 23, 546},
	}
	for _, test := range tests {
		if result := add(test.a, test.b); result != test.c {
			fmt.Println("error!")
		}
	}

}
