package ecclient

import (
	"context"
	"encoding/base64"
	"github.com/fxamacker/cbor/v2"
)

var m struct {
	PostTransaction Transaction `graphql:"postTransaction(transaction: $transaction)"`
}

func (c *Client) PostTransaction(transaction TransactionRequest) error {
	transaction.Nonce = c.GetNextNonce()
	transactionBytes, _ := cbor.Marshal(transaction)
	sign1Bytes := sign1(transactionBytes, c.PrivateKey)
	variables := map[string]interface{}{
		"transaction": Bytes(base64.StdEncoding.EncodeToString(sign1Bytes)),
	}
	err := c.client.Mutate(context.Background(), &m, variables)
	check(err)
	return nil

}
