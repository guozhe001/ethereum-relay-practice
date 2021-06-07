package wallet

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	wallet, err := CreateWallet("this_is_for_test")
	assert.NoError(t, err)
	log.Println(wallet)
}
