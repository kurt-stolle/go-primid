# PrimID for Go
Transform your internal id's to obfuscated integers based hashes.
This has the advantage that is is incredibly fast compared to other methods.

For the original Primid project written in JavaScript, see the [GitHub page](https://github.com/kurt-stolle/node-primid).

# Installation
`go get github.com/kurt-stolle/go-primid`

# Methods
PrimID uses only three functions.
```go
import (
    "github.com/kurt-stolle/go-primid"
)

// Main function, for this small example
func main(){
    // Get a new generator using desired settings (should be unique settings per application).
    // These should be prime numbers!
    var generator=primid.NewGenerator(prime,inverse,xor)

    // Encode a number, get a integer based hash
    var hash=generator.Encode(some_number)

    // Decode a hash, get back a number
    var number=generator.Decode(hash)
}
```

# Numbers
You need to pick your own prime numbers unique to your application. For suggestions, look at [this list](http://primes.utm.edu/lists/small/millions/).

# Usage example
1. Create a new `primid` instance using
```go
var generator = primid.NewGenerator(1580030173, 1163945558)
```
2. Generate a hash, in this example we want to obfruscate the value `15`
```go
var hash=generator.Encode(15);
```
3. Use the hash as output of your API or other system.
4. Turn the hash back into a number. If the hash from the example above is used, then `id` is equal to `15`.
```go
var id=generator.Decode(hash);
```
