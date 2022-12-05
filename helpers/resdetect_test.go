package helpers

import (
	"testing"
)

func TestGetMaxX(t *testing.T) {
	if GetMaxX() == 0 {
		t.Errorf("GetMaxX() returned 0")
	}
}

func TestGetMaxY(t *testing.T) {
	if GetMaxY() == 0 {
		t.Errorf("GetMaxY() returned 0")
	}
}
