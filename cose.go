package ecclient

import (
	"crypto/ed25519"
	"fmt"
	"reflect"

	"github.com/fxamacker/cbor/v2"
)

// Unprotected protected
type Unprotected struct {
	Kid []byte `cbor:"4,keyasint,omitempty"`
}

// Protected Protected
type Protected struct {
	Alg         int  `cbor:"1,keyasint,omitempty"`
	ContentType uint `cbor:"3,keyasint"`
}

type sign1Message struct {
	_           struct{} `cbor:",toarray"`
	Protected   []byte
	Unprotected Unprotected
	Payload     []byte
	Signature   []byte
}

func sign1(message []byte, privateKey ed25519.PrivateKey) []byte {
	tags := cbor.NewTagSet()
	if err := tags.Add(
		cbor.TagOptions{EncTag: cbor.EncTagRequired, DecTag: cbor.DecTagRequired},
		reflect.TypeOf(sign1Message{}),
		18); err != nil {
		fmt.Println("error:", err)
	}
	em, _ := cbor.EncOptions{}.EncModeWithTags(tags)
	Protected, _ := cbor.Marshal(Protected{
		Alg:         -8,
		ContentType: 0,
	})
	Payload, err := cbor.Marshal([]interface{}{"Signature1", Protected, []byte{}, message})
	check(err)

	Signature := ed25519.Sign(privateKey, Payload)
	b, err := em.Marshal(sign1Message{
		Protected: Protected,
		Unprotected: Unprotected{
			Kid: privateKey.Public().(ed25519.PublicKey),
		},
		Payload:   message,
		Signature: Signature,
	})
	return b

}
