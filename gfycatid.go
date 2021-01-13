package gfycatid

import (
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/MichaelDiBernardo/gfycatid/assets"
)

// Generator for gfycat-style IDs. Mutex is there to lock access to rand; the
// rand docs suggest that sources made with NewSource are not safe for access by
// multiple goroutines.
type Generator struct {
	rand *rand.Rand
	nadj uint // Number of adjectives to generate.
	m    sync.Mutex
}

// Create a generator that will make IDs with `nadj` preceding adjectives,
// followed by one animal.
func Create(nadj uint) (*Generator, error) {
	src := rand.NewSource(time.Now().UnixNano())
	return CreateWithSource(src, nadj)
}

// CreateWithSource makes a generator that will use the given source to generate
// random numbers.
func CreateWithSource(s rand.Source, nadj uint) (*Generator, error) {
	if nadj == 0 {
		return nil, errors.New("nadj must be non-zero")
	}
	return &Generator{
		rand: rand.New(s),
		nadj: nadj,
	}, nil
}

// Gen generates a new randomized ID with the # of adjectives specified at
// construction.
func (gen *Generator) Gen() string {
	var result strings.Builder
	for n := uint(0); n < gen.nadj; n++ {
		i := gen.r(assets.NumAdjectives)
		adj := strings.Title(assets.Adjectives[i])
		result.WriteString(adj)
	}

	i := gen.r(assets.NumAnimals)
	animal := strings.Title(assets.Animals[i])
	result.WriteString(animal)
	return result.String()
}

func (gen *Generator) r(n int) int {
	gen.m.Lock()
	defer gen.m.Unlock()
	return gen.rand.Intn(n)
}
