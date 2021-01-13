package gfycatid

import (
	"math/rand"
	"testing"
)

func TestBadCreate(t *testing.T) {
	_, err := Create(0)

	if err == nil {
		t.Error("Expected err for nadj=0")
	}
}

func TestGen(t *testing.T) {
	src := rand.NewSource(12)
	gen, err := CreateWithSource(src, 1)

	if err != nil {
		t.Errorf("Unexpected err: %v", err)
	}

	if act, exp := gen.Gen(), "SlategrayThrasher"; act != exp {
		t.Errorf("Expected %s, got %s", exp, act)
	}

	gen, err = CreateWithSource(src, 2)

	if err != nil {
		t.Errorf("Unexpected err: %v", err)
	}

	if act, exp := gen.Gen(), "NavajowhiteJubilantPolarbear"; act != exp {
		t.Errorf("Expected %s, got %s", exp, act)
	}

	// At this point I'm mostly just doing this for fun.
	gen, err = CreateWithSource(src, 3)

	if err != nil {
		t.Errorf("Unexpected err: %v", err)
	}

	if act, exp := gen.Gen(), "ThirstyVioletMutedChafer"; act != exp {
		t.Errorf("Expected %s, got %s", exp, act)
	}
}
