package noise

// The following constants represent the details of this implementation of the Noise specification.
const (
	NoiseDraftVersion = "33"
	NoiseDH           = "25519"
	NoiseAEAD         = "ChaChaPoly"
	NoiseHASH         = "SHA256"
)

// The following constants are taken directly from the Noise specification.
const (
	NoiseMessageLength    = 65535 - 2 // 2-byte length
	NoiseTagLength        = 16
	NoiseMaxPlaintextSize = NoiseMessageLength - NoiseTagLength
)

type Config struct {
	// the type of Noise protocol that the client and the server will go through
	HandshakePattern noiseHandshakeType
	// the current peer's keyPair
	KeyPair *KeyPair
	// the other peer's public key
	RemoteKey *[32]byte
	// any messages that the client and the server previously exchanged in clear
	Prologue []byte
	// if the chosen handshake pattern requires the current peer to send a static
	// public key as part of the handshake, this proof over the key is mandatory
	// in order for the other peer to verify the current peer's key
	StaticPublicKeyProof []byte
	// if the chosen handshake pattern requires the remote peer to send an unknown
	// static public key as part of the handshake, this callback is mandatory in
	// order to validate it
	PublicKeyVerifier func(publicKey, proof []byte) bool
	// by default a noise protocol is full-duplex, meaning that both the client
	// and the server can write on the channel at the same time. Setting this value
	// to true will require the peers to write and read in turns. If this requirement
	// is not respected by the application, the consequences could be catastrophic
	HalfDuplex bool
}
