package client

import (
	"errors"
	"io/ioutil"
	"net/http"

	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/musinit/go-defi/v2/models"
)

// Address is a compound contract address type
type Address string

func (a Address) String() string {
	return string(a)
}

// EthAddress returns a typed ethereum address
func (a Address) EthAddress() common.Address {
	return common.HexToAddress(a.String())
}

const (
	// Mainnet
	// CompoundBAT is the address of the cBAT contract
	CompoundBAT = Address("0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e")
	// CompoundDAI is the address of the cDAI contract
	CompoundDAI = Address("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643")
	// CompoundSAI is the address of the cSAI contract
	CompoundSAI = Address("0xf5dce57282a584d2746faf1593d3121fcac444dc")
	// CompoundETH is the address of the cETH contract
	CompoundETH = Address("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")
	// CompoundREP is the address of the cREP contract
	CompoundREP = Address("0x158079ee67fce2f58472a96584a73c7ab9ac95c1")
	// CompoundUSDC is the address of the cUSDC contract
	CompoundUSDC = Address("0x39aa39c021dfbae8fac545936693ac917d5e7563")
	// CompoundWBTC is the address of the cWBTC contract
	CompoundWBTC = Address("0xc11b1268c1a384e55c48c2391d8d480264a3a7f4")
	// CompoundZRX is th eaddress of the cZRX contract
	CompoundZRX = Address("0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407")
	// Comptroller is th address of the comptroller contract
	Comptroller = Address("0x178053c06006e67e09879C09Ff012fF9d263dF29")
	// Unitroller is the address of the unitroller contract
	Unitroller   = Address("0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b")
	CompoundUSDT = Address("0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9")

	// GOERLI
	CompoundUSDC_goerli = Address("0x73506770799Eb04befb5AaE4734e58C2C624F493")

	// Common
	USDC        = Address("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	USDC_goerli = Address("0x07865c6e87b9f70255377e024ace6630c1eaa37f")
)

// Polygon mumbai
const (
	MATIC_polygon_mumbai       = Address("0x0000000000000000000000000000000000001010")
	WMATIC_polygon_mumbai      = Address("0x9c3c9283d3e44854697cd22d3faa240cfb032889")
	WMATIC_AAVE_polygon_mumbai = Address("0xf237dE5664D3c2D2545684E76fef02A3A58A364c")
	AaveLendingPoolV3Mumbai    = Address("0x4ff9EB4C5F5f790CE7bB51d4876481C2eefDE00F")
	USDT_polygon_mumbai        = Address("0xEF4aEDfD3552db80E8F5133ed5c27cebeD2fE015")
	USDC_polygon_mumbai        = Address("0xa45C8154149C3C383c226269add5ffB5cdA7ccb7")
)

// Polygon mainnet
const (
	USDC_polygon  = Address("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174")
	USDT_polygon  = Address("0xc2132D05D31c914a87C6611C10748AEb04B58e8F")
	MATIC_polygon = Address("0x0000000000000000000000000000000000001010")
)

// Aave Mainnet
const (
	AaveWMATICv3      = Address("0x6d80113e533a2C0fe82EaBD35f1875DcEA89Ea97")
	AaveUSDCv3        = Address("0x625E7708f30cA75bfd92586e17077590C60eb4cD")
	AaveUSDTv3        = Address("0x6ab707Aca953eDAeFBc4fD23bA73294241490620")
	AaveDAIv3         = Address("0x82E64f49Ed5EC1bC6e43DAD4FC8Af9bb3A2312EE")
	AaveLendingPoolV3 = Address("0x794a61358D6845594F94dc1DB02A252b5b4814aD")
	AaveLendingPoolV2 = Address("0x8dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf")
	AaveUSDTv2        = Address("0x60D55F02A771d515e077c9C2403a1ef324885CeC")
	AaveUSDCv2        = Address("0x1a13F4Ca1d028320A707D99520AbFefca3998b7F")
	AaveDAIv2         = Address("0x27F8D03b3a2196956ED754baDc28D73be8830A6e")
)

// Aave mumbai
const (
	LendingPoolMumbai = Address("0x1758d4e6f68166C4B2d9d0F049F33dEB399Daa1F")
	AaveUSDTMumbai    = Address("0xEF4aEDfD3552db80E8F5133ed5c27cebeD2fE015")
	AaveUSDCMumbai    = Address("0xCdc2854e97798AfDC74BC420BD5060e022D14607")
	AaveDAIMumbai     = Address("0xDD4f3Ee61466C4158D394d57f3D4C397E91fBc51")
	AaveWMATICMumbai  = Address("0xC0e5f125D33732aDadb04134dB0d351E9bB5BCf6")
)

var (
	// CompoundTokens is map containing the name, and address of all compound tokens
	CompoundTokens = map[string]Address{
		"cBAT":         CompoundBAT,
		"cDAI":         CompoundDAI,
		"cSAI":         CompoundSAI,
		"cETH":         CompoundETH,
		"cREP":         CompoundREP,
		"cUSDC":        CompoundUSDC,
		"cWBTC":        CompoundWBTC,
		"cZRX":         CompoundZRX,
		"сUSDT":        CompoundUSDT,
		"сUSDT_goerli": CompoundUSDC_goerli,
		"aUSDCv3":      AaveUSDCv3,
		"aUSDTv3":      AaveUSDTv3,
		"aUSDTv2":      AaveUSDTv2,
	}
)

// GetTotalSupplyInterestedEarned is used to return the total supply interest earned for a particular account
func (c *Client) GetTotalSupplyInterestedEarned(resp *models.AccountResponse) (float64, error) {
	if len(resp.Accounts) == 0 {
		return 0, errors.New("no accounts found")
	}
	var (
		account = resp.Accounts[0]
		value   float64
	)
	for _, tk := range account.Tokens {
		// skip tokens with no balance
		if tk.LifetimeSupplyInterestAccrued.Value == "0" ||
			tk.LifetimeSupplyInterestAccrued.Value == "0.0" {
			continue
		}
		val, err := strconv.ParseFloat(tk.LifetimeSupplyInterestAccrued.Value, 64)
		if err != nil {
			return 0, err
		}
		value = value + val
	}
	return value, nil
}

// GetSupplyInterestEarned is used to retrieve the interest earned by supply a particular token
func (c *Client) GetSupplyInterestEarned(token Address, resp *models.AccountResponse) (float64, error) {
	tkn, err := models.GetTokenByAddress(token.String(), resp)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(tkn.LifetimeSupplyInterestAccrued.Value, 64)
}

// GetBorrowInterestedAccrued is used to retrieve the interest you owe for borrowing
func (c *Client) GetBorrowInterestedAccrued(token Address, resp *models.AccountResponse) (string, error) {
	tkn, err := models.GetTokenByAddress(token.String(), resp)
	if err != nil {
		return "", err
	}
	return tkn.LifetimeBorrowInterestAccrued.Value, nil
}

// GetLiquidatableAccounts is used to return all accounts with health below 1.0
// indicating they can be liquidated. The keys of the map are the addresses
// and the values are their health
func (c *Client) GetLiquidatableAccounts(pageSize, pageNum string) (map[string]float64, error) {
	resp, err := c.GetAccounts(pageSize, pageNum)
	if err != nil {
		return nil, err
	}
	if len(resp.Accounts) == 0 {
		return nil, errors.New("an unexpected error occurred")
	}
	out := make(map[string]float64)
	for _, acct := range resp.Accounts {
		health, err := strconv.ParseFloat(acct.Health.Value, 64)
		if err != nil {
			return nil, err
		}
		if health < 1.0 {
			out[acct.Address] = health
		}
	}
	if len(out) == 0 {
		return nil, errors.New("no liquidatable accounts found")
	}
	return out, nil
}

// sendRequest is used to send a request, and return the given body bytes
func (c *Client) sendRequest(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("content-type", "application/json")
	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}
