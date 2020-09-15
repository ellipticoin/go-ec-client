package ecclient

import (
	"github.com/shurcooL/graphql"
	"math/rand"
	"strconv"

	"github.com/r3labs/sse"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// RandomNode Chooses a Random Node
func NewClient(privateKey []byte) *Client {
	node := RandomNode()
	return &Client{
		PrivateKey: privateKey,
		client:     graphql.NewClient(node, nil),
		sseClient:  sse.NewClient(node),
	}
}

// RandomNode Chooses a Random Node
func RandomNode() string {
	return Bootnodes[rand.Intn(len(Bootnodes))]
}

// OnNewBlock OnNewBlock
func (c *Client) OnNewBlock(callback func(uint32)) error {
	if c.sseClient == nil {
		c.sseClient = sse.NewClient(RandomNode())
	}
	c.sseClient.Subscribe("block", func(msg *sse.Event) {
		n, err := strconv.Atoi(string(msg.Data))
		check(err)
		callback(uint32(n))
	})
	return nil
}
