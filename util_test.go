package main

import (
	"github.com/guozhe001/ethereum-relay-practice/util"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestWeiToEth(t *testing.T) {
	_, err := util.WeiToEth("0x1400e93770")
	assert.NoError(t, err)
	//0.00000008591463
}

func TestGetZero(t *testing.T) {
	length := 6
	zero, err := util.GetZero(length)
	assert.NoError(t, err)
	assert.Equal(t, length, len(zero))
	log.Print(zero)
}
