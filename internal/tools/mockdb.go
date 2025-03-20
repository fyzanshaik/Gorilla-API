package tools

import "time"

type mockDB struct{}



var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"maria": {
		AuthToken: "789GHI",
		Username:  "maria",
	},
	"linda": {
		AuthToken: "101JKL",
		Username:  "linda",
	},
	"john": {
		AuthToken: "112MNO",
		Username:  "john",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    250,
		Username: "jason",
	},
	"maria": {
		Coins:    500,
		Username: "maria",
	},
	"linda": {
		Coins:    300,
		Username: "linda",
	},
	"john": {
		Coins:    150,
		Username: "john",
	},
}


func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second*1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}