package wallet

import (
  "github.com/stretchr/testify/assert"
  "log"
  "testing"
)

// 这是新建钱包的测试，没有必要一直跑
//func TestCreateWallet(t *testing.T) {
//  wallet, err := CreateWallet("this_is_for_test")
//  assert.NoError(t, err)
//  log.Println(wallet)
//}

func TestUnlockETHWallet(t *testing.T) {
  err := UnlockETHWallet("0xcd97b8b5ef2f219bed97c9b5b8035454f3f9ed18", "this_is_for_test")
  assert.NoError(t, err)
}

func TestGetCurrentPath(t *testing.T) {
  log.Println(getCurrentPath())
}

func TestGetKeystoreDir(t *testing.T) {
  log.Println(getKeystoreDir())
}
