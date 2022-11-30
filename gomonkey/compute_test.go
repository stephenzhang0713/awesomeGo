package gomonkey

import (
	"github.com/agiledragon/gomonkey/v2"
	"testing"
)

func TestCompute(t *testing.T) {
	patches := gomonkey.ApplyFunc(networkCompute, func(a, b int) (int, error) {
		return 2, nil
	})

	defer patches.Reset()

	sum, err := Compute(1, 1)
	if sum != 2 || err != nil {
		t.Errorf("Expected %v,got %v", 2, sum)
	}
}
