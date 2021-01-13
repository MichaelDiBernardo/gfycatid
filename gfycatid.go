package gfycatid

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// Generator for gfycat-style IDs. Mutex is there to lock access to rand; the
// rand docs suggest that sources made with NewSource are not safe for access by
// multiple goroutines.
type Generator struct {
	rand *rand.Rand
	nadj uint
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
func (*Generator) Gen() string {
	return ""
}

// 	firstAdjective := adjectives[rand.Intn(len(adjectives))]
// 	firstAdjective[0] = firstAdjective[0] - 32

// 	secondAdjective := adjectives[rand.Intn(len(adjectives))]
// 	secondAdjective[0] = secondAdjective[0] - 32

// 	animal := animals[rand.Intn(len(animals))]
// 	animal[0] = animal[0] - 32

// 	id := bytes.Join([][]byte{firstAdjective, secondAdjective, animal}, []byte{})

// 	return string(id)
// }

// // UpdateAssets() attempts to pull the latest assets from gfycat
// func UpdateAssets() error {
// 	adjectiveResp, err := http.Get(adjectivesURL)
// 	if err != nil {
// 		return err
// 	}
// 	defer adjectiveResp.Body.Close()
// 	if b, err := ioutil.ReadAll(adjectiveResp.Body); err == nil {
// 		assets.Adjectives = b
// 	} else {
// 		return err
// 	}

// 	animalResp, err := http.Get(animalsURL)
// 	if err != nil {
// 		return err
// 	}
// 	defer animalResp.Body.Close()
// 	if b, err := ioutil.ReadAll(animalResp.Body); err == nil {
// 		assets.Animals = b
// 	} else {
// 		return err
// 	}

// 	return nil
// }
