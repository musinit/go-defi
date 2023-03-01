package client_test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/musinit/go-defi/v2/client"
	"github.com/musinit/go-defi/v2/config"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	ctx             = context.TODO()
	mainnet_polygon = "https://polygon-mainnet.infura.io/v3/e2d37f84e4a34fa4bc2997f45e1c2883"
	mumbai_polygon  = "https://polygon-mumbai.infura.io/v3/e2d37f84e4a34fa4bc2997f45e1c2883"
)

func Test_GetBalance_Aave(t *testing.T) {
	acc := GetFirstAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mainnet_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	_, ethclient, err := client.ConfigToOpts(&cfg)
	fmt.Println(acc.Address.String())

	balance, err := ethclient.BalanceAt(ctx, acc.Address, nil)
	assert.Nil(t, err)
	balanceVal := float32(balance.Int64()) / 1000000000000000000

	assert.True(t, balanceVal >= 0)
}

func Test_AccountApprove_aave(t *testing.T) {
	acc := GetFirstAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mainnet_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	amount := big.NewInt(1000000000000000000)
	err = bclient.Approve(ctx, client.USDC_polygon, client.AaveUSDCv3, amount)
	assert.Nil(t, err)
	assert.NotNil(t, bclient)
}

func Test_AccountMint_aave_v2(t *testing.T) {
	acc := GetFirstAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mainnet_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	val := int64(1000000)
	v := big.NewInt(val)
	accAddress := client.Address(acc.Address.String())
	err = bclient.MintAaveV2(ctx, client.AaveLendingPoolV2, accAddress, v)
	assert.Nil(t, err)
	assert.NotNil(t, bclient)
}

func GetFirstAccount() accounts.Account {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.Accounts()[0]
}

func GetSolviStorage() accounts.Account {
	ks := keystore.NewKeyStore("./polygon_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.Accounts()[0]
}

func GetLastAccount() accounts.Account {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.Accounts()[7]
}
