package client_test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/musinit/go-defi/v2/client"
	"github.com/musinit/go-defi/v2/config"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	json       = "{\"address\":\"7cd8053a6b0f56e2a6c60fd907978c776a3ac2ae\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"998860706b2fecabe5490a05bd3eee08be0cb99a880ad35810a2cec0210f3c5f\",\"cipherparams\":{\"iv\":\"b5a5e69bd4eba814ee2092a4bef8b4d8\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"d5a77475edb7b181508806f8cebeafe4e315bf9e005a97eda6a3cf51569c4526\"},\"mac\":\"3b1cfd670392a5b056c222773e9d89d57a4fccd29b9e6012ffdbe283ae7a0624\"},\"id\":\"80803f50-3853-4fd1-aa30-84c65c8b8282\",\"version\":3}"
	name       = "UTC--2022-12-11T08-53-22.972243000Z--7cd8053a6b0f56e2a6c60fd907978c776a3ac2ae"
	passphrase = "Creation password"
)

func Test_AccountMainnetCreate(t *testing.T) {
	ctx := context.TODO()
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
	addr := common.BytesToAddress([]byte("0xdafea492d9c6733ae3d56b7ed1adb60692c98bc5"))
	tr, isPending, err := ethclient.TransactionByHash(ctx, common.BytesToHash([]byte("0xe4ad6a5afbf24f4a1bcec7f714738d9c3d4e4cb79ab3240a098be74c73874dea")))
	fmt.Println(isPending)
	fmt.Println(tr)
	bal, err := ethclient.BalanceAt(ctx, addr, nil)
	assert.Nil(t, err)
	bb := bal.String()
	fmt.Println(bb)
	bclient := client.NewBClient(auth, ethclient)

	assert.NotNil(t, bclient)
}

func Test_ExportAndImportAccount(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	passphrase := "Creation password"
	newAcc, _ := ks.NewAccount(passphrase)
	fmt.Println(newAcc)
	json, err := ks.Export(newAcc, passphrase, passphrase)
	assert.Nil(t, err)
	err = ks.Unlock(newAcc, passphrase)
	assert.Nil(t, err)
	err = ks.Delete(newAcc, passphrase)
	assert.Nil(t, err)

	ac, err := ks.Import([]byte(json), passphrase, passphrase)
	assert.Nil(t, err)
	assert.NotNil(t, ac)
}

func Test_ImportAccount(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)

	acc := ks.Accounts()[0]

	assert.NotNil(t, acc)
}

func Test_AccountGetBalance(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc0 := ks.Accounts()[0]
	acc1 := ks.Accounts()[1]
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: "https://mainnet.infura.io/v3/e2d37f84e4a34fa4bc2997f45e1c2883",
		KeyFile:  fmt.Sprintf("%s", acc0.URL.Path),
		KeyPass:  "Creation password",
	}}
	_, ethclient, err := client.ConfigToOpts(&cfg)
	assert.Nil(t, err)

	fmt.Printf("address 0: %s\n", acc0.Address.String())
	fmt.Printf("address 1: %s\n", acc1.Address.String())

	balance, err := ethclient.BalanceAt(ctx, acc0.Address, nil)
	assert.Nil(t, err)
	balanceVal := float32(balance.Int64()) / 1000000000000000000

	assert.True(t, balanceVal >= 0)
}

func Test_AccountApprove(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc := ks.Accounts()[0]
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: "https://mainnet.infura.io/v3/e2d37f84e4a34fa4bc2997f45e1c2883",
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	amount := big.NewInt(1000000000)
	err = bclient.Approve(ctx, client.USDC, client.CompoundUSDC, amount)
	assert.Nil(t, err)
	assert.NotNil(t, bclient)
}

func Test_AccountMint(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc := ks.Accounts()[0]
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: "https://mainnet.infura.io/v3/e2d37f84e4a34fa4bc2997f45e1c2883",
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	amount := big.NewInt(1000000)
	err = bclient.Mint(ctx, client.CompoundUSDC, amount)
	assert.Nil(t, err)
	assert.NotNil(t, bclient)
}

func Test_AccountTransact(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc := ks.Accounts()[0]
	ADDR := "0x25eCfA4F0042baA5b30114D4C25E7599F97Ff8f7"
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: "https://mainnet.infura.io/v3/e2d37f84e4a34fa4bc2997f45e1c2883",
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	_, ethclient, err := client.ConfigToOpts(&cfg)
	var (
		sender = common.HexToAddress(ADDR)
		sk     = crypto.ToECDSAUnsafe(common.FromHex("87d2362b9880f795dd006dbb4fe4059034ad4915c008042102ad4498ebde1f5c"))
		//value    = new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether))
		gasLimit = uint64(21000)
	)
	tipCap, _ := ethclient.SuggestGasTipCap(context.Background())
	feeCap, _ := ethclient.SuggestGasPrice(context.Background())
	nonce, err := ethclient.PendingNonceAt(context.Background(), sender)
	chainID, err := ethclient.ChainID(ctx)
	val := int64(0.003 * 1000000000000000000)
	v := big.NewInt(val)
	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       gasLimit,
			To:        &acc.Address,
			Value:     v,
			Data:      nil,
		})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), sk)
	assert.Nil(t, err)
	err = ethclient.SendTransaction(ctx, signedTx)
	assert.Nil(t, err)
}

func Test_AccountBalanceOfUnderlying(t *testing.T) {
	ks := keystore.NewKeyStore("./test_accounts/", keystore.StandardScryptN, keystore.StandardScryptP)
	ctx := context.TODO()
	acc := ks.Accounts()[0]
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: "https://mainnet.infura.io/v3/e2d37f84e4a34fa4bc2997f45e1c2883",
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)
	assert.Nil(t, err)
	owner := client.Address("")
	tx, err := bclient.BalanceOfUnderlying(ctx, client.CompoundUSDC, owner)
	assert.NotNil(t, tx)

	tt := tx.Value()

	if tt != nil {
		kk := tt.String()
		fmt.Println(kk)
	}
	fmt.Println(tt)

}
