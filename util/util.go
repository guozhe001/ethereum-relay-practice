package util

import (
	"fmt"
	"github.com/guozhe001/ethereum-relay-practice/constant"
	"log"
	"math/big"
)

// WeiToEth 单位转换，把wei转换成eth
func WeiToEth(hexWei string) (string, error) {
	ten, _ := new(big.Int).SetString(hexWei[2:], 16)
	log.Println(ten)
	return ten.String(), nil
}

// GetZero 获取指定数量的0的字符串；主要是用在调用以太坊时组装data
func GetZero(number int) (string, error) {
	if number < 1 || number > 64 {
		return "", fmt.Errorf("只能获取1到64个0组成的字符串")
	}
	return constant.ZeroOf64[:number], nil
}
