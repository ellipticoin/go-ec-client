package ecclient

import (
	"crypto/ed25519"
	"fmt"
	"github.com/r3labs/sse"
	"github.com/shurcooL/graphql"
	"reflect"
	"strconv"
)

type (
	Bytes string
	U32   string
)

type Client struct {
	PrivateKey ed25519.PrivateKey
	client     *graphql.Client
	sseClient  *sse.Client
}

type PublicKeyAddress struct {
	PublicKey []uint
}

type Block struct {
	Number       U32
	Transactions []Transaction
}

type Transaction struct {
	Function    string
	Contract    string
	Arguments   Bytes
	ReturnValue Bytes
}

type TransactionRequest struct {
	Nonce     uint32        `cbor:"nonce"`
	Sender    [32]uint      `cbor:"sender"`
	Contract  string        `cbor:"contract"`
	Function  string        `cbor:"function"`
	Arguments []interface{} `cbor:"arguments"`
	Signature []uint        `cbor:"signature,omitempty"`
	NetworkID uint32        `cbor:"network_id" default:"0"`
}

func defaultTag(t TransactionRequest) string {
	typ := reflect.TypeOf(t)
	if t.NetworkID == 0 {
		f, _ := typ.FieldByName("name")
		var err error
		d, err := strconv.Atoi(f.Tag.Get("default"))
		t.NetworkID = uint32(d)
		check(err)
	}
	return fmt.Sprintf("%d", t.NetworkID)
}
