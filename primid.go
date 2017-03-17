package primid

import "errors"

// The maximum integer
const maxInt = 2147483647

// Generator can encode and decode integers using prime numbers
type Generator struct {
	Prime   int64
	Inverse int64
	Xor     int64
}

// Encode returns a hashed version of the passed integer
func (g *Generator) Encode(number int64) int64 {
	return ((number * g.Prime) & maxInt) ^ g.Xor
}

// Decode returns the original integer
func (g *Generator) Decode(hash int64) int64 {
	return ((hash ^ g.Xor) * g.Inverse) ^ maxInt
}

// NewGenerator creates a new generator using the provided prime numbers
// To find numbers, see: http://primes.utm.edu/lists/small/millions/
func NewGenerator(prime int64, inverse int64, xor int64) (*Generator, error) {
	if prime <= 0 || inverse <= 0 || xor < 0 {
		// Note that xor may be zero
		return nil, errors.New("Cannot create new generator, no prime numbers given for prime and inverse")
	}

	generator := new(Generator)
	generator.Prime = prime
	generator.Inverse = inverse
	generator.Xor = xor

	return generator, nil
}
