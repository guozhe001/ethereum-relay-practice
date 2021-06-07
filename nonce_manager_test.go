package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"reflect"
	"testing"
)

func TestNewNonceManager(t *testing.T) {
	tests := []struct {
		name string
		want *NonceManager
	}{
		{
			name: "test001",
			want: NewNonceManager(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if nonceManager := NewNonceManager(); !reflect.DeepEqual(nonceManager, tt.want) {
				t.Errorf("NewNonceManager() = %v, want %v", nonceManager, tt.want)
			}
		})
	}
}

func TestNonceManager_GetNonce(t *testing.T) {
	initNonce := *big.NewInt(12)
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "test001",
			args: args{address: "0xfc18d8882697552737921dF830102579f000008D"},
			want: &initNonce,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNonceManager()
			n.SetNonce(tt.args.address, &initNonce)
			if got := n.GetNonce(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNonce() = %v, want %v", got, tt.want)
			}
			n.PlusNonce(tt.args.address)
			assert.Equal(t, n.GetNonce(tt.args.address).String(), initNonce.String())
		})
	}
}
