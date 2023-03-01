package client_test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/musinit/go-defi/v2/client"
	"github.com/musinit/go-defi/v2/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AccountCreate(t *testing.T) {
	pass := "secrete_passphrase"
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	newAcc, _ := ks.NewAccount(pass)
	fmt.Println(newAcc)
	_, err := ks.Export(newAcc, pass, pass)
	assert.Nil(t, err)
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: "http://localhost:8545",
		KeyFile:  fmt.Sprintf("%s", newAcc.URL.Path),
		KeyPass:  pass,
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	assert.Nil(t, err)
	bclient := client.NewBClient(auth, ethclient)

	assert.NotNil(t, bclient)
}

func Test_GetSuplyRate(t *testing.T) {
	ctx := context.Background()
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	newAcc, _ := ks.NewAccount("Creation password")
	fmt.Println(newAcc)
	_, err := ks.Export(newAcc, "Creation password", "Export password")
	assert.Nil(t, err)
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: "https://mainnet.infura.io/v3/e2d37f84e4a34fa4bc2997f45e1c2883",
		KeyFile:  fmt.Sprintf("%s", newAcc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	assert.Nil(t, err)
	bclient := client.NewBClient(auth, ethclient)
	assert.NotNil(t, bclient)

	value, err := bclient.SupplyRatePerBlock(ctx, client.CompoundUSDC)
	vv := value.String()

	assert.True(t, vv != "")
	assert.Nil(t, err)
}

func Test_AccountPlay(t *testing.T) {
	ctx := context.Background()
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	newAcc, _ := ks.NewAccount("Creation password")
	fmt.Println(newAcc)
	_, err := ks.Export(newAcc, "Creation password", "Export password")
	assert.Nil(t, err)
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: "http://localhost:8545",
		KeyFile:  fmt.Sprintf("%s", newAcc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	assert.Nil(t, err)
	bclient := client.NewBClient(auth, ethclient)
	assert.NotNil(t, bclient)

	r, err := bclient.GetBorrowRate(ctx, client.CompoundDAI)
	assert.Nil(t, err)

	rr := r.Int64()
	assert.True(t, rr > 0)
}
