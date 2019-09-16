package primid

import (
	"errors"
	"math"
	"math/big"
)

// Errors
var ErrNotPrime = errors.New("number is not a prime");
var ErrTooLarge = errors.New("cannot calclulate inverse because prime is larger than maximum value of int64")

// The maximum integer
const maxInt = uint64(math.MaxInt32) 

// Generator can encode and decode integers using prime numbers
type Generator struct {
	prime   uint64
	inverse uint64
	xor     uint64
}

// Encode returns a hashed version of the passed integer
func (g *Generator) Encode(number uint64) uint64 {
	return ((number * g.prime) & maxInt) ^ g.xor
}

// Decode returns the original integer
func (g *Generator) Decode(hash uint64) uint64 {
	return ((hash ^ g.xor) * g.inverse) & maxInt
}

// NewGenerator creates a new generator using the provided prime numbers
// To find numbers, see: http://primes.utm.edu/lists/small/millions/
func NewGenerator(prime, random uint64) (*Generator, error) {
	inverse, err := Inverse(prime)
	if err != nil {
		return nil, err
	}

	generator := &Generator{
		prime: prime,
		inverse: inverse,
		xor: random,
	}

	return generator, nil
}

// NewGeneratorRaw creates a new generator using the provided values without testing for validity of the provided numbers
func NewGeneratorRaw(prime, inverse, xor uint64) *Generator {
	return &Generator{
		prime: prime,
		inverse: inverse,
		xor: xor,
	}
}

// Inverse calculates the inverse of prime
// Limitation is that prime must be less than the maximum uint64 value
func Inverse(prime uint64) (uint64, error) {
	if prime > math.MaxInt64 {
		return 0, ErrTooLarge
	}
	
	if !IsPrime(prime) {
		return 0, ErrNotPrime
	}

	var max = big.NewInt(int64(maxInt + 1))
	var target big.Int

	return target.ModInverse(big.NewInt(int64(prime)), max).Uint64(), nil
}

// Check if a number is a prime using Miller-Rabin test
func IsPrime(prime uint64) bool {
	return big.NewInt(int64(prime)).ProbablyPrime(20)
}