package client_test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	"time"
)

var prodPass = "h2_nfjdsNJ%25u7*1034bjfjdsk2nk39&Y(#NjLNFdnj32PassphraseProd"

func Test_AccountApprove_aave_v3(t *testing.T) {
	acc := GetLastAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mainnet_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  prodPass,
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	appoveVal := big.NewInt(1000000)
	err = bclient.Approve(ctx, client.USDT_polygon, client.AaveLendingPoolV3, appoveVal)
	assert.Nil(t, err)
}

func Test_AccountMint_aave_v3(t *testing.T) {
	acc := GetLastAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mainnet_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  prodPass,
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	val := int64(100000)
	v := big.NewInt(val)
	accAddress := client.Address(acc.Address.String())
	fmt.Println(accAddress)
	err = bclient.MintAaveV3(ctx, client.AaveLendingPoolV3, accAddress, v, client.USDT_polygon.EthAddress())
	assert.Nil(t, err)
	assert.NotNil(t, bclient)
}

func Test_AccountMintWMATIC_aave_v3(t *testing.T) {
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
	err = bclient.MintAaveV3(ctx, client.AaveLendingPoolV3, accAddress, v, client.WMATIC_polygon_mumbai.EthAddress())
	assert.Nil(t, err)
	assert.NotNil(t, bclient)
}

func Test_AccountWithdraw_aave_v3(t *testing.T) {
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
	err = bclient.WithdrawAaveV3(ctx, client.AaveLendingPoolV3, accAddress, v, client.USDT_polygon.EthAddress())
	assert.Nil(t, err)
	assert.NotNil(t, bclient)
}

func Test_AccountGetReserveNormilizedIncome_aave_v3(t *testing.T) {
	acc := GetFirstAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mainnet_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	result, err := bclient.GetReserveNormilizedIncomeAaveV3(ctx, client.AaveLendingPoolV3, client.AaveUSDTv3)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func Test_AccountBalanceOf_aave_mainnet_v3(t *testing.T) {
	acc := GetFirstAccount()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mainnet_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  "Creation password",
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	bclient := client.NewBClient(auth, ethclient)

	accAddress := client.Address("0x7cd8053A6B0F56E2A6C60Fd907978C776A3Ac2AE")
	balance, err := bclient.BalanceOf(ctx, client.AaveUSDTv3, accAddress)
	balanceVal := float32(balance.Int64()) / 1000000

	assert.Nil(t, err)
	assert.True(t, balanceVal >= 0)
}

func Test_GetBalance_AaveV3(t *testing.T) {
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

func Test_AccountSendMatic_aave_v3(t *testing.T) {
	acc := GetFirstAccount()
	secretePassword := "passphrase"
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mainnet_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  secretePassword,
	}}
	ADDR := "0x62C859f161e860c2a127aFE78FE1AD05DF26b4EB"
	TO := "0x5333782886Fb2ad7f19Afc574d0239780796d334"
	_, ethclient, err := client.ConfigToOpts(&cfg)
	var (
		sender = common.HexToAddress(ADDR)
		to     = common.HexToAddress(TO)
		sk     = crypto.ToECDSAUnsafe(common.FromHex("private_key"))
		//value    = new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether))
		gasLimit = uint64(21000)
	)
	tipCap, _ := ethclient.SuggestGasTipCap(ctx)
	feeCap, _ := ethclient.SuggestGasPrice(ctx)
	nonce, err := ethclient.PendingNonceAt(ctx, sender)
	chainID, err := ethclient.ChainID(ctx)
	val := int64(10000000000000)
	v := big.NewInt(val)
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
	ctx, _ = context.WithTimeout(context.Background(), time.Second*5)
	rcpt, err := bind.WaitMined(ctx, ethclient, signedTx)
	assert.Nil(t, err)
	assert.True(t, rcpt.Status == 1)
	assert.Nil(t, err)
}

func Test_AccountSendUSDT(t *testing.T) {
	acc := GetSolviStorage()
	pass := "1698d4c5-c764-4ed1-a10b-0c7910797a97"
	s := "0x2d1B29919527da57676419EBff927338A29B6B1e"
	pk := "0x721eaf0157e24defee5255343e84a7661589bb29a419c1c4f3c854a05afd9ee6"
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mainnet_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  pass,
	}}
	toAddress := common.HexToAddress("0x62C859f161e860c2a127aFE78FE1AD05DF26b4EB")
	_, ethclient, err := client.ConfigToOpts(&cfg)
	var (
		usdtAddress = common.HexToAddress(client.USDT_polygon.String())
		sender      = common.HexToAddress(s)
		sk          = crypto.ToECDSAUnsafe(common.FromHex(pk))
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
	amount.SetString("1000000", 10) // 1 token
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
	ctx, _ = context.WithTimeout(context.Background(), time.Second*60)
	rcpt, err := bind.WaitMined(ctx, ethclient, tx)
	assert.Nil(t, err)
	assert.True(t, rcpt.Status == 1)
	assert.Nil(t, err)
}

func Test_AccountGetReserveData_aave_v3(t *testing.T) {
	acc := GetSolviStorage()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mainnet_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  prodPass,
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	assert.Nil(t, err)
	bclient := client.NewBClient(auth, ethclient)
	owner := client.Address(acc.Address.String())
	result, err := bclient.GetReserveDataAaveV3(ctx, client.AaveLendingPoolV3, owner, client.AaveUSDTv3)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func Test_AccountGetUserAccountData_aave_v3(t *testing.T) {
	acc := GetSolviStorage()
	cfg := config.Config{Blockchain: config.Blockchain{
		Endpoint: mainnet_polygon,
		KeyFile:  fmt.Sprintf("%s", acc.URL.Path),
		KeyPass:  prodPass,
	}}
	auth, ethclient, err := client.ConfigToOpts(&cfg)
	assert.Nil(t, err)
	bclient := client.NewBClient(auth, ethclient)

	owner := client.Address(acc.Address.String())
	result, err := bclient.GetUserAccountDataAaveV3(ctx, client.AaveLendingPoolV3, owner, client.AaveUSDTv3)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
