package util

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/guozhe001/ethereum-relay-practice/constant"
	"io/ioutil"
	"log"
	"strings"
)

// GetMethodId 根据api的绝对路径名称和方法名获取方法的methodId
func GetMethodId(abiString, methodName string) (string, error) {
	a, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return "", err
	}
	method := a.Methods[methodName]
	methodId := constant.HexPrefix + common.Bytes2Hex(method.ID)
	log.Println("methodId=", methodId)
	return methodId, nil
}

// GetMethodId 根据api的绝对路径名称和方法名获取方法的methodId
func GetMethodIdByFile(abiFileName, methodName string) (string, error) {
	file, err := ioutil.ReadFile(abiFileName)
	if err != nil {
		return "", err
	}
	return GetMethodId(string(file), methodName)
}
