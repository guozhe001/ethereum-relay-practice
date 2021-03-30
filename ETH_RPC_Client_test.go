package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewETHRPCClient(t *testing.T) {
	rpc := NewETHRPCClient("www.baidu.com").GetRpc()
	assert.NotNil(t, rpc)
}

func TestNewETHRPCClientFailed(t *testing.T) {
	NewETHRPCClient("haha://.\\hi.com")
}
