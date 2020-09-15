package ecclient

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"strconv"
)

var nonce struct {
	NextNonce interface{} `graphql:"nextNonce(address: $address)"`
}

var block struct {
	Block Block `graphql:"block(blockNumber: $blockNumber)"`
}

var transaction struct {
	Transaction Transaction `graphql:"transaction(transactionId: $transactionId)"`
}

func (c *Client) GetNextNonce() uint32 {
	variables := map[string]interface{}{
		"address": Bytes(base64.StdEncoding.EncodeToString(c.PrivateKey.Public().(ed25519.PublicKey))),
	}
	err := c.client.Query(context.Background(), &nonce, variables)
	check(err)
	nonce, err := strconv.Atoi(nonce.NextNonce.(string))
	check(err)
	return uint32(nonce)
}

func (c *Client) GetBlock(blockNumber uint32) Block {
	variables := map[string]interface{}{
		"blockNumber": U32(strconv.Itoa(int(blockNumber))),
	}
	err := c.client.Query(context.Background(), &block, variables)
	check(err)
	return block.Block
}

func (c *Client) GetTransaction(transactionId uint32) Transaction {
	variables := map[string]interface{}{
		"transactionId": U32(strconv.Itoa(int(transactionId))),
	}
	err := c.client.Query(context.Background(), &transaction, variables)
	check(err)
	return transaction.Transaction
}
