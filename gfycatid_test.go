package gfycatid

import "testing"

func TestBadCreate(t *testing.T) {
	_, err := Create(0)

	if err == nil {
		t.Error("Expected err for nadj=0")
	}
}
