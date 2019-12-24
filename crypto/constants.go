package crypto

import "math/big"

const (
	Ed25519KeySize        = 32
	AESKeySize            = 32
	CommitmentRingSize    = 8
	CommitmentRingSizeExp = 3
	CStringBulletProof    = "bulletproof"
	CStringBurningAddress = "burningaddress"
)

const (
	MaxSizeInfoCoin = 255 		// bytes
)

const (
	MaxInputNumber  = 255
	MaxOutputNumber = 255

	MaxInputNumberV2  = 32
	MaxOutputNumberV2 = 32
)

var LInt = new(big.Int).SetBytes([]byte{0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14, 0xde, 0xf9, 0xde, 0xa2, 0xf7, 0x9c, 0xd6, 0x58, 0x12, 0x63, 0x1a, 0x5c, 0xf5, 0xd3, 0xed})