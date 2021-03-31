package util

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"log"
	"strings"
)

// GetMethodId 根据api的绝对路径名称和方法名获取方法的methodId
func GetMethodId(abiFileName, methodName string) (string, error) {
	file, err := ioutil.ReadFile(abiFileName)
	if err != nil {
		return "", err
	}
	log.Println(string(file))
	a, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		return "", err
	}
	method := a.Methods[methodName]
	return common.Bytes2Hex(method.ID), nil
}
