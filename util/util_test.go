package util

import (
  "github.com/stretchr/testify/assert"
  "log"
  "testing"
)

func TestWeiToEth(t *testing.T) {
  _, err := WeiToEth("0x1400e93770")
  assert.NoError(t, err)
  //0.00000008591463
}

func TestGetZero(t *testing.T) {
  length := 6
  zero, err := GetZero(length)
  assert.NoError(t, err)
  assert.Equal(t, length, len(zero))
  log.Print(zero)
}
