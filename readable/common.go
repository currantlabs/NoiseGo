package noise

// The following constants represent the details of this Noise implementation.
const (
	NoiseDraftVersion = "33"
	NoiseDH           = "25519"
	NoiseAEAD         = "ChaChaPoly"
	NoiseHASH         = "SHA256"
)

// The following constants represent constants from the Noise specification.
const (
	// A noise message's length. The specification says it should be 65535,
	// but the specification forgets to mention the 2-byte of required header (length)
	NoiseMessageLength    = 65535 - 2
	NoiseTagLength        = 16
	NoiseMaxPlaintextSize = NoiseMessageLength - NoiseTagLength
)

type Config struct {
	// This is MANDATORY to verify public keys sent by the other peer
	publicKeyVerifier func([]byte) bool

	handshakePattern              noiseHandshakeType
	prologue                      []byte
	handshakeDataToSend           [][]byte
	handshakeReceivedDataCallBack []func([]byte) error
	keyPair, ephemeralKeyPair     *keyPair
	remoteKey, ephemeralRemoteKey *keyPair
}