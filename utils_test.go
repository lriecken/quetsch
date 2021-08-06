package main

import (
	"testing"
)

func TestScaleDims(t *testing.T) {
	minWidth, minHeight := scaleDims(1000, 1000, 0.1)
	if minWidth != 100 || minHeight != 100 {
		t.Fatal("Scaling has wrong results")
	}
}

func TestGetMinimalScale(t *testing.T) {
	minWidth, minHeight, minScale := getMinimalScale(2000, 1000, 100, 100)
	if minWidth != 200 || minHeight != 100 || minScale != 0.1 {
		t.Fatal("Scale calculated wrong when width > height")
	}
	minWidth, minHeight, minScale = getMinimalScale(1000, 2000, 100, 100)
	if minWidth != 100 || minHeight != 200 || minScale != 0.1 {
		t.Fatal("Scale calculated wrong when width > height")
	}
	minWidth, minHeight, minScale = getMinimalScale(1000, 1000, 100, 100)
	if minWidth != 100 || minHeight != 100 || minScale != 0.1 {
		t.Fatal("Scale calculated wrong when width > height")
	}
}
