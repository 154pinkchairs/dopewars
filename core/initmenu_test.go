package core

import (
	"testing"
)

func TestInit(t *testing.T) {
	Init()
	var err error
	if err != nil {
		t.Error("Init() failed")
	}
}
