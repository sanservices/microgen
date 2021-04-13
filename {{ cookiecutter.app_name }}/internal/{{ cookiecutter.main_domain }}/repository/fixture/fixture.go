package fixture

import (
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
)

// Fixture implementation of {{ cookiecutter.main_domain }} repo
// returns random data for testing
// NOT FOR PRODUCTION USE, TESTING PURPOSES ONLY
type Fixture struct {
}

// New returns a new Fixture struct
func New() (im *Fixture) {
	gofakeit.Seed(rand.Int63())
	gofakeit.SetGlobalFaker(gofakeit.NewCrypto())
	return &Fixture{}
}
