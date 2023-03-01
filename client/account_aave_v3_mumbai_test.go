package client_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/musinit/go-defi/v2/client"
	"github.com/musinit/go-defi/v2/config"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
	"math/big"
	"testing"
)

func Test_AccountApprove_aave_v3_mumbai(t *testing.T) {
	acc := GetFirstAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mumbai_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	appoveVal := big.NewInt(1000000)
	err = bclient.Approve(ctx, client.USDC_polygon_mumbai, client.AaveUSDCMumbai, appoveVal)
	assert.Nil(t, err)
}

func Test_AccountMint_aave_v3_mumbai(t *testing.T) {
	acc := GetFirstAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mumbai_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	auth.GasLimit = 50000
	bclient := client.NewBClient(auth, ethclient)

	val := int64(1000000000000000000)
	v := big.NewInt(val)
	accAddress := client.Address(acc.Address.String())
	err = bclient.MintAaveV3(ctx, client.AaveLendingPoolV3Mumbai, accAddress, v, client.MATIC_polygon_mumbai.EthAddress())
	assert.Nil(t, err)
	assert.NotNil(t, bclient)
}

func Test_AccountBalanceOf_mumbai_v3(t *testing.T) {
	acc := GetFirstAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mumbai_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	accAddress := client.Address("0x3ac166e3D4EB04ddd3c3c3475BbB2aCC09A13705")
	balance, err := bclient.BalanceOf(ctx, client.MATIC_polygon_mumbai, accAddress)
	balanceVal := float32(balance.Int64()) / 1000000000000000000

	assert.Nil(t, err)
	assert.True(t, balanceVal >= 0)
}

func Test_GetBalance_AaveMumbai(t *testing.T) {
	acc := GetFirstAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mumbai_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	_, ethclient, err := client.ConfigToOpts(&cfg)
	fmt.Println(acc.Address.String())

	a := common.HexToAddress("0xE7dc3a5156448aE71521EE64cD8c5b20CcAf55d8")

	balance, err := ethclient.BalanceAt(ctx, a, nil)

	assert.Nil(t, err)
	balanceVal := float32(balance.Int64()) / 1000000000000000000

	assert.True(t, balanceVal >= 0)
}

func Test_AccountSendWMATIC_mumbai(t *testing.T) {
	acc := GetFirstAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mumbai_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	ADDR := "0x7cd8053A6B0F56E2A6C60Fd907978C776A3Ac2AE"
	toAddress := common.HexToAddress("0x5333782886Fb2ad7f19Afc574d0239780796d334")
	_, ethclient, err := client.ConfigToOpts(&cfg)
	var (
		usdtAddress = common.HexToAddress(client.WMATIC_polygon_mumbai.String())
		sender      = common.HexToAddress(ADDR)
		sk          = crypto.ToECDSAUnsafe(common.FromHex("d2adaab3d03aadee0031e4de04fc5ff3a0c2048932198e753e710c78d414f10c"))
		gasLimit    = uint64(250000)
	)
	tipCap, _ := ethclient.SuggestGasTipCap(ctx)
	feeCap, _ := ethclient.SuggestGasPrice(ctx)
	nonce, err := ethclient.PendingNonceAt(ctx, sender)
	chainID, err := ethclient.ChainID(ctx)
	v := big.NewInt(0)
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))
	amount := new(big.Int)
	amount.SetString("10000000000000000", 10) // 1 token
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       gasLimit,
			To:        &usdtAddress,
			Value:     v,
			Data:      data,
		})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), sk)
	assert.Nil(t, err)
	err = ethclient.SendTransaction(ctx, signedTx)
	assert.Nil(t, err)
}
