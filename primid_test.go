package primid

import (
	"testing"
	"math"
)

// TestInverse tests the inverse calculator
func TestInverse(t *testing.T){
	var inv,err = Inverse(1580030173)
	if err != nil {
		t.Fatalf("failed to invert prime: %s", err.Error())
	}

	if inv != 59260789 {
		t.Fatalf("inverse not valid, expected %d, got %d", 59260789, inv)
	}
}

// TestGenerator crudely tests the generator by checking whether some numbers go through successfully
func TestGenerator(t *testing.T) {
	var generator, err = NewGenerator(1580030173, 12312)
	if err != nil {
		t.Fatalf("failed to initialize generator: %s", err.Error())
	}

	const increment =  math.MaxInt32 / 100
	for num := uint64(0); num < math.MaxInt32; num += increment {
		var enc=generator.Encode(num)
		var dec=generator.Decode(enc)

		t.Logf("%d => %d => %d\n", num,enc,dec)

		if num != dec {
			t.Fatal("decoding failed")
		}
	}
}