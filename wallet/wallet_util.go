package wallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateWallet(password string) (string, error) {
	if "" == password {
		return "", fmt.Errorf("please input password！")
	}
	if len(password) < 6 {
		return "", fmt.Errorf("password length must big than 6！")
	}
	keyDir := ".keystores"
	ks := keystore.NewKeyStore(keyDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		return "", err
	}
	return account.Address.String(), nil
}
