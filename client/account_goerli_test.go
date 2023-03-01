package client_test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/musinit/go-defi/v2/client"
	"github.com/musinit/go-defi/v2/config"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	goerliAddress = "https://goerli.infura.io/v3/e2d37f84e4a34fa4bc2997f45e1c2883"
	daiAddress    = client.Address("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")
)

func Test_GetBalance_Goerli(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc := ks.Accounts()[0]
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: goerliAddress,
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

func Test_AccountApprove_goerli(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc := ks.Accounts()[0]
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: goerliAddress,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	amount := big.NewInt(1000000000000000000)
	err = bclient.Approve(ctx, client.USDC_goerli, client.CompoundUSDC_goerli, amount)
	assert.Nil(t, err)
	assert.NotNil(t, bclient)
}

func Test_AccountMint_goerli(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc := ks.Accounts()[0]
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: goerliAddress,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	val := int64(500000)
	v := big.NewInt(val)
	err = bclient.Mint(ctx, client.CompoundUSDC_goerli, v)
	assert.Nil(t, err)
	assert.NotNil(t, bclient)
}

func Test_AccountTransact_goerli(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc := ks.Accounts()[0]
	ADDR := "0x7cd8053A6B0F56E2A6C60Fd907978C776A3Ac2AE"
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: goerliAddress,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	_, ethclient, err := client.ConfigToOpts(&cfg)
	var (
		sender = common.HexToAddress(ADDR)
		sk     = crypto.ToECDSAUnsafe(common.FromHex("d2adaab3d03aadee0031e4de04fc5ff3a0c2048932198e753e710c78d414f10c"))
		//value    = new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether))
		gasLimit = uint64(21000)
	)
	tipCap, _ := ethclient.SuggestGasTipCap(context.Background())
	feeCap, _ := ethclient.SuggestGasPrice(context.Background())
	nonce, err := ethclient.PendingNonceAt(context.Background(), sender)
	chainID, err := ethclient.ChainID(ctx)
	val := int64(0.005 * 1000000000000000000)
	v := big.NewInt(val)
	to := common.BytesToAddress([]byte("0x73506770799Eb04befb5AaE4734e58C2C624F493"))
	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       gasLimit,
			To:        &to,
			Value:     v,
			Data:      nil,
		})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), sk)
	assert.Nil(t, err)
	err = ethclient.SendTransaction(ctx, signedTx)
	assert.Nil(t, err)
}

func Test_AccountBalanceOfUnderlying_goerli(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc := ks.Accounts()[0]
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: goerliAddress,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)
	assert.Nil(t, err)
	owner := client.Address(acc.Address.String())
	tx, err := bclient.BalanceOfUnderlying(ctx, client.CompoundUSDC_goerli, owner)
	assert.NotNil(t, tx)

	tt := tx.Value()

	if tt != nil {
		kk := tt.String()
		fmt.Println(kk)
	}
	fmt.Println(tt)
}

func Test_SupplyRate_goerli(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc := ks.Accounts()[0]
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: goerliAddress,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)
	assert.Nil(t, err)
	tx, err := bclient.SupplyRatePerBlock(ctx, client.CompoundUSDC_goerli)
	assert.NotNil(t, tx)

	tt := tx.Int64()
	v := tt / 1000000000000000000
	fmt.Println(v)
	fmt.Println(tt)
}

func Test_SupplyExchangeRate_goerli(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc := ks.Accounts()[0]
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: goerliAddress,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)
	assert.Nil(t, err)
	tx, err := bclient.ExchangeRateCurrent(ctx, client.CompoundUSDC_goerli)
	assert.NotNil(t, tx)

	tt := tx.Value().Int64()
	v := float32(tt) / 1000000000000000000
	fmt.Println(v)
	fmt.Println(tt)
}

func Test_DecryptKey(t *testing.T) {
	pass := "_"
	jsonAccount := "_"
	//ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	//acc := ks.Accounts()[7]

	key, err := keystore.DecryptKey([]byte(jsonAccount), pass)
	assert.Nil(t, err)
	assert.NotNil(t, key)

	privateKey := key.PrivateKey
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Printf("%s", hexutil.Encode(privateKeyBytes))

	assert.NotNil(t, privateKey)
}
