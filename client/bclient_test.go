package client_test

import (
	client "github.com/musinit/go-defi/v2/client"
	"github.com/musinit/go-defi/v2/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	api         = "https://api.compound.finance/api/v2"
	account     = "0xe18999d3f7e1a84e35bf9c15699b79e46eb44704"
	selfAccount = "0xcd0fBe49Ac5e009858DFd1c5F7330907C710fe96"
)

func Test_BClient(t *testing.T) {
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: "http://localhost:8545",
		KeyFile:  "",
		KeyPass:  "",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	assert.Nil(t, err)
	bclient := client.NewBClient(auth, ethclient)

	assert.NotNil(t, bclient)
}
