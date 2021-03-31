package wallet

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	wallet, err := CreateWallet("")
	assert.NoError(t, err)
	log.Println(wallet)
}
