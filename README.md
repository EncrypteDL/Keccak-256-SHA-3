# Keccak-256-SHA-3

## Description
Keccak-256 (SHA-3) is a specific instance of the Keccak cryptographic hash function, which was selected as the winner of the NIST hash function competition and subsequently standardized as SHA-3 (Secure Hash Algorithm 3).

### Technical Details:
- State Size: Keccak operates on a state of 1600 bits, organized as a 5x5x64 matrix.
- Rate and Capacity: For Keccak-256, the rate (r) is 1088 bits and the capacity (c) is 512 bits. The rate determines how much of the input is absorbed into the state at each iteration, and the capacity provides security.
- Permutation Function: The core of Keccak is its permutation function, Keccak-f, which applies a series of bitwise operations to transform the state.


## Installation

### Prerequisites
```
Go version 1.22
Python 3.12
deep knowledge of cryptographic
```
```
# Clone the repository
git clone https://github.com/EncrypteID/Keccak-256-SHA-3.git

# Change to the project directory
cd Keccak-256-SHA-3

# Build the project
go build
python -u
```

# Usage
- Cryptocurrencies: Keccak-256 is widely used in blockchain and cryptocurrency applications, including Ethereum.
- Digital Signatures: Used in digital signature schemes to ensure the integrity and authenticity of messages.
- General Security: Utilized in various security protocols and systems that require secure hash functions.

## Features
1.	**Hash Length**: Produces a 256-bit (32-byte) hash value.
2.	**Security:** Designed to provide security against a wide range of attacks, including collision attacks, preimage attacks, and length extension attacks.
3.	**Sponge Construction:** Uses a sponge construction that allows the same algorithm to be used for hashing, pseudorandom number generation, and other cryptographic functions.

### Example Usage in Python:
```
from Crypto.Hash import keccak

def keccak256(data):
    k = keccak.new(digest_bits=256)
    k.update(data.encode('utf-8'))
    return k.hexdigest()

# Example
print(keccak256("Hello, world!"))  # Outputs the Keccak-256 hash of the input string
```
``
Output: Hello, world! 
Input: Keccak-256 Hash: 644bcc7e564373040e3fd35074d0f9201bb59e4c661ad49d9d0e60734531e4dc
``
### Example Usage in Go:
``
go get golang.org/x/crypto/sha3
``

```
package main

import (
	"crypto/sha3"
	"encoding/hex"
	"fmt"
	"os"
)

func keccak256(input string) string {
	hash := sha3.New256()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run keccak256_client.go <input_string>")
		return
	}

	input := os.Args[1]
	hashOutput := keccak256(input)

	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Keccak-256 Hash: %s\n", hashOutput)
}
```

## Difference Between Keccak-256 and SHA3-256
### Difference Between Keccak-256 and SHA3-256

**Keccak-256** and **SHA3-256** refer to closely related cryptographic hash functions, but there are important distinctions between them due to standardization and minor technical differences. Here's a detailed comparison:

#### Origins:

- **Keccak-256**: This is the original hash function that was submitted to the NIST hash function competition by its creators, which eventually won and became the basis for SHA-3.
- **SHA3-256**: This is the standardized version of Keccak-256, published by NIST in the FIPS 202 standard.

#### Technical Differences:

1. **Padding**:
   - **Keccak-256**: Uses a specific padding scheme known as "Keccak-p".
   - **SHA3-256**: Uses a slightly different padding scheme called "SHA-3 padding" or "SHA-3 variant of Keccak".

2. **Domain Separation**:
   - **Keccak-256**: The original Keccak implementation does not include domain separation.
   - **SHA3-256**: Introduces domain separation, which allows the use of the same permutation function for different purposes within SHA-3 (e.g., for hashing, message authentication codes, etc.).

#### Usage in Cryptocurrencies:

- **Keccak-256**: 
  - Commonly used in Ethereum for hashing purposes, including generating addresses and signing transactions. 
  - The Ethereum community typically refers to the function as `keccak256`.
  
- **SHA3-256**: 
  - Follows the standardized version as per NIST's FIPS 202.
  - While technically more robust due to the additional standardization steps, it is less commonly referenced in certain cryptocurrency communities that had already adopted Keccak before SHA-3 was finalized.

#### Code Examples:

Below are examples of how to compute both Keccak-256 and SHA3-256 in Python and Go.

#### Python Example:

```python
# Keccak-256
from Crypto.Hash import keccak

def keccak256(data):
    k = keccak.new(digest_bits=256)
    k.update(data.encode('utf-8'))
    return k.hexdigest()

print("Keccak-256:", keccak256("Hello, world!"))

# SHA3-256
from hashlib import sha3_256

def sha3256(data):
    return sha3_256(data.encode('utf-8')).hexdigest()

print("SHA3-256:", sha3256("Hello, world!"))
```

#### Go Example:

```go
package main

import (
	"crypto/sha3"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func keccak256(input string) string {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func sha3256(input string) string {
	hash := sha3.New256()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	input := "Hello, world!"

	// Keccak-256
	keccakHash := keccak256(input)
	fmt.Printf("Keccak-256: %s\n", keccakHash)

	// SHA3-256
	sha3Hash := sha3256(input)
	fmt.Printf("SHA3-256: %s\n", sha3Hash)
}
```

### Summary

- **Keccak-256** and **SHA3-256** are cryptographic hash functions based on the same underlying algorithm but differ in padding and domain separation.
- **Keccak-256** is widely used in Ethereum and other pre-standardization applications.
- **SHA3-256** is the standardized version as defined by NIST, incorporating slight modifications for added security and versatility.
- Both provide strong cryptographic security, but it's important to use the appropriate one depending on the application context, especially when interoperability or compliance with standards is required.


## Contributing
