package keccak256sha3

import "crypto"

// func New224() hash.Hash {
// 	return new224()
// }


type state struct{
	rate int
	outputLen int
	dsbyte byte
}

func new224Generic() *state {
	return &state{rate: 144, outputLen: 28, dsbyte: 0x06}
}

func new256Generic() *state {
	return &state{rate: 136, outputLen: 32, dsbyte: 0x06}
}

func new384Generic() *state {
	return &state{rate: 104, outputLen: 48, dsbyte: 0x06}
}

func new512Generic() *state {
	return &state{rate: 72, outputLen: 64, dsbyte: 0x06}
}

func init() {
	crypto.RegisterHash(crypto.SHA3_224, New224)
	crypto.RegisterHash(crypto.SHA3_256, New256)
	crypto.RegisterHash(crypto.SHA3_384, New384)
	crypto.RegisterHash(crypto.SHA3_512, New512)
}
